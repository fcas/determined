DROP VIEW trials_augmented_view; -- TODO delete when we rebase.

ALTER TABLE experiments RENAME TO experiments_v2;
ALTER TABLE experiments_v2 RENAME COLUMN id TO run_collection_id;

CREATE TABLE run_collections (
  id integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
  name text NOT NULL,
  state experiment_state NOT NULL,
  notes text NOT NULL,
  project_id integer NOT NULL REFERENCES projects(id),
  owner_id integer NOT NULL REFERENCES users(id),
  progress double precision,
  archived bool NOT NULL DEFAULT false,
  start_time timestamptz NOT NULL,
  end_time timestamptz,
  checkpoint_size bigint NOT NULL DEFAULT 0,
  checkpoint_count integer NOT NULL DEFAULT 0,
  external_run_collection_id TEXT UNIQUE NULL
);

INSERT INTO run_collections (
  id,
  name,
  state,
  notes,
  project_id,
  owner_id,
  progress,
  archived,
  start_time,
  end_time,
  checkpoint_size,
  checkpoint_count,
  external_run_collection_id
)
SELECT run_collection_id,
  'experiment_id:' || run_collection_id || CASE WHEN external_experiment_id IS NULL THEN
      '' ELSE ', external_experiment_id:' || external_experiment_id END,
  state,
  notes,
  project_id,
  owner_id,
  progress,
  archived,
  start_time,
  end_time,
  checkpoint_size,
  checkpoint_count,
  external_experiment_id
FROM experiments_v2;

ALTER TABLE experiments_v2
  DROP COLUMN state,
  DROP COLUMN notes,
  DROP COLUMN project_id,
  DROP COLUMN owner_id,
  DROP COLUMN progress,
  DROP COLUMN archived,
  DROP COLUMN start_time,
  DROP COLUMN end_time,
  DROP COLUMN checkpoint_size,
  DROP COLUMN checkpoint_count,
  DROP COLUMN external_experiment_id;

ALTER TABLE experiments_v2 ALTER COLUMN run_collection_id DROP IDENTITY;

ALTER TABLE experiments_v2 ADD CONSTRAINT fk_experiment_run_collection
  FOREIGN KEY (run_collection_id) REFERENCES run_collections(id) ON DELETE CASCADE;

CREATE VIEW experiments AS
SELECT
  -- experiments_v2.
  e.run_collection_id AS id,
  e.config AS config,
  e.model_definition AS model_definition,
  e.git_remote AS git_remote,
  e.git_commit AS git_commit,
  e.git_committer AS git_committer,
  e.git_commit_date AS git_commit_date,
  e.model_packages AS model_packages,
  e.parent_id AS parent_id,
  e.original_config AS original_config,
  e.job_id AS job_id,
  e.best_trial_id AS best_trial_id,
  e.unmanaged AS unmanaged,

  -- run_collections.
  rc.state AS state,
  rc.notes AS notes,
  rc.project_id AS project_id,
  rc.owner_id AS owner_id,
  rc.progress AS progress,
  rc.archived AS archived,

  rc.start_time AS start_time,
  rc.end_time AS end_time,
  rc.checkpoint_size AS checkpoint_size,
  rc.checkpoint_count AS checkpoint_count,

  rc.external_run_collection_id AS external_experiment_id
FROM experiments_v2 e
JOIN run_collections rc ON e.run_collection_id = rc.id;


CREATE OR REPLACE FUNCTION autoupdate_exp_best_trial_metrics() RETURNS trigger AS $$
BEGIN
    WITH bt AS (
        SELECT id, best_validation_id
        FROM trials
        WHERE experiment_id = NEW.experiment_id
        ORDER BY searcher_metric_value_signed LIMIT 1)
    UPDATE experiments_v2 SET best_trial_id = bt.id FROM bt
    WHERE experiments_v2.run_collection_id = NEW.experiment_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
