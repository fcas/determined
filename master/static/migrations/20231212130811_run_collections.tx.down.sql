DROP VIEW experiments;

ALTER TABLE experiments_v2 DROP CONSTRAINT fk_experiment_run_collection;

ALTER TABLE experiments_v2
ALTER COLUMN run_collection_id
ADD GENERATED BY DEFAULT AS IDENTITY;

ALTER TABLE experiments_v2
  ADD COLUMN state experiment_state NOT NULL,
  ADD COLUMN notes text NOT NULL,
  ADD COLUMN project_id integer NOT NULL REFERENCES projects(id),
  ADD COLUMN owner_id integer NOT NULL REFERENCES users(id),
  ADD COLUMN progress double precision,
  ADD COLUMN archived bool NOT NULL DEFAULT false,
  ADD COLUMN start_time timestamptz NOT NULL,
  ADD COLUMN end_time timestamptz,
  ADD COLUMN checkpoint_size bigint NOT NULL DEFAULT 0,
  ADD COLUMN checkpoint_count integer NOT NULL DEFAULT 0,
  ADD COLUMN external_experiment_id TEXT UNIQUE NULL;

UPDATE experiments_v2 SET
  state = rc.state,
  notes = rc.notes,
  project_id = rc.project_id,
  owner_id = rc.owner_id,
  progress = rc.progress,
  archived = rc.archived,
  start_time = rc.start_time,
  end_time = rc.end_time,
  checkpoint_size = rc.checkpoint_size,
  checkpoint_count = rc.checkpoint_count,
  external_experiment_id = rc.external_run_collection_id
FROM run_collections rc
WHERE rc.id = experiments_v2.run_collection_id;

ALTER TABLE experiments_v2 RENAME COLUMN run_collection_id TO id;

ALTER TABLE experiments_v2 RENAME TO experiments;

CREATE OR REPLACE FUNCTION autoupdate_exp_best_trial_metrics() RETURNS trigger AS $$
BEGIN
    WITH bt AS (
        SELECT id, best_validation_id
        FROM trials
        WHERE experiment_id = NEW.experiment_id
        ORDER BY searcher_metric_value_signed LIMIT 1)
    UPDATE experiments SET best_trial_id = bt.id FROM bt
    WHERE experiments.id = NEW.experiment_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
