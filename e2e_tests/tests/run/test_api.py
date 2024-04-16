import time

import pytest

from determined.common import api
from determined.common.api import bindings, errors
from tests import api_utils
from tests import config as conf
from tests import experiment as exp


def wait_for_run_state(
    test_session: api.Session,
    run_id: int,
    expected_state: bindings.trialv1State,
    timeout: int = 30,
) -> None:
    deadline = time.time() + timeout
    while time.time() < deadline:
        resp = bindings.get_SearchRuns(
            test_session,
            limit=1,
            filter="""{"filterGroup":{"children":[{"columnName":"id","kind":"field",
        "location":"LOCATION_TYPE_RUN","operator":"=","type":"COLUMN_TYPE_NUMBER","value":"""
            + str(run_id)
            + """}],"conjunction":"and","kind":"group"},"showArchived":false}""",
        )
        if expected_state == resp.runs[0].state:
            return
        time.sleep(0.1)
    pytest.fail(f"task failed to complete after {timeout} seconds")


@pytest.mark.e2e_cpu
def test_run_pause_and_resume() -> None:
    sess = api_utils.user_session()
    exp_id = exp.create_experiment(
        sess, conf.fixtures_path("no_op/single.yaml"), conf.fixtures_path("no_op")
    )

    searchResp = bindings.get_SearchRuns(
        sess,
        limit=1,
        filter="""{"filterGroup":{"children":[{"columnName":"experimentId","kind":"field",
        "location":"LOCATION_TYPE_RUN","operator":"=","type":"COLUMN_TYPE_NUMBER","value":"""
        + str(exp_id)
        + """}],"conjunction":"and","kind":"group"},"showArchived":false}""",
    )

    assert searchResp.runs[0].state == bindings.trialv1State.ACTIVE
    run_id = searchResp.runs[0].id
    pauseResp = bindings.post_PauseRuns(
        sess, body=bindings.v1PauseRunsRequest(runIds=[run_id], projectId=1, skipMultitrial=False)
    )

    # validate response
    assert len(pauseResp.results) == 1
    assert pauseResp.results[0].id == run_id
    assert pauseResp.results[0].error == ""

    # ensure that run is paused
    wait_for_run_state(sess, run_id, bindings.trialv1State.PAUSED)

    resumeResp = bindings.post_ResumeRuns(
        sess, body=bindings.v1ResumeRunsRequest(runIds=[run_id], projectId=1, skipMultitrial=False)
    )

    assert len(resumeResp.results) == 1
    assert resumeResp.results[0].id == run_id
    assert resumeResp.results[0].error == ""

    # ensure that run is unpaused
    wait_for_run_state(sess, run_id, bindings.trialv1State.ACTIVE)


@pytest.mark.e2e_cpu
def test_run_pause_and_resume_filter_no_skip() -> None:
    sess = api_utils.user_session()
    exp_id = exp.create_experiment(
        sess,
        conf.fixtures_path("mnist_pytorch/adaptive_short.yaml"),
        conf.fixtures_path("mnist_pytorch"),
    )

    runFilter = (
        """{"filterGroup":{"children":[{"columnName":"experimentId","kind":"field",
            "location":"LOCATION_TYPE_RUN","operator":"=","type":"COLUMN_TYPE_NUMBER","value":"""
        + str(exp_id)
        + """}, {"columnName":"hp.n_filters2","kind":"field",
            "location":"LOCATION_TYPE_RUN_HYPERPARAMETERS","operator":">=","type":"COLUMN_TYPE_NUMBER",
            "value":40}],"conjunction":"and","kind":"group"},"showArchived":false}"""
    )

    pauseResp = bindings.post_PauseRuns(
        sess,
        body=bindings.v1PauseRunsRequest(
            runIds=[],
            filter=runFilter,
            projectId=1,
            skipMultitrial=False,
        ),
    )

    # validate response
    for res in pauseResp.results:
        assert res.error == ""
        wait_for_run_state(sess, res.id, bindings.trialv1State.PAUSED)

    resumeResp = bindings.post_ResumeRuns(
        sess,
        body=bindings.v1ResumeRunsRequest(
            runIds=[], projectId=1, filter=runFilter, skipMultitrial=False
        ),
    )

    for res in resumeResp.results:
        assert res.error == ""
        wait_for_run_state(sess, res.id, bindings.trialv1State.ACTIVE)


@pytest.mark.e2e_cpu
def test_run_pause_and_resume_filter_skip_empty() -> None:
    sess = api_utils.user_session()
    exp_id = exp.create_experiment(
        sess,
        conf.fixtures_path("mnist_pytorch/adaptive_short.yaml"),
        conf.fixtures_path("mnist_pytorch"),
    )

    runFilter = (
        """{"filterGroup":{"children":[{"columnName":"experimentId","kind":"field",
            "location":"LOCATION_TYPE_RUN","operator":"=","type":"COLUMN_TYPE_NUMBER","value":"""
        + str(exp_id)
        + """}, {"columnName":"hp.n_filters2","kind":"field",
            "location":"LOCATION_TYPE_RUN_HYPERPARAMETERS","operator":">=","type":"COLUMN_TYPE_NUMBER",
            "value":40}],"conjunction":"and","kind":"group"},"showArchived":false}"""
    )
    pauseResp = bindings.post_PauseRuns(
        sess,
        body=bindings.v1PauseRunsRequest(
            runIds=[],
            filter=runFilter,
            projectId=1,
            skipMultitrial=True,
        ),
    )

    # validate response
    for r in pauseResp.results:
        assert r.error == "Skipping run '" + str(r.id) + "' (part of multi-trial)."

    resumeResp = bindings.post_ResumeRuns(
        sess,
        body=bindings.v1ResumeRunsRequest(
            runIds=[], projectId=1, filter=runFilter, skipMultitrial=False
        ),
    )

    for res in resumeResp.results:
        assert res.error == ""
        wait_for_run_state(sess, res.id, bindings.trialv1State.ACTIVE)


@pytest.mark.e2e_cpu
def test_run_pause_and_resume_filter_no_skip() -> None:
    sess = api_utils.user_session()

    with pytest.raises(
        errors.APIException, match="if filter is provided run id list must be empty"
    ):
        bindings.post_PauseRuns(
            sess,
            body=bindings.v1PauseRunsRequest(
                runIds=[123],
                filter="filter",
                projectId=1,
                skipMultitrial=True,
            ),
        )
