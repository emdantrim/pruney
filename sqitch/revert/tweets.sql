-- Revert pruney:tweets from sqlite

BEGIN;

  DROP TABLE tweets;

COMMIT;
