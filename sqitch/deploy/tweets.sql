-- Deploy pruney:tweets to sqlite

BEGIN;

  CREATE TABLE tweets (
    username TEXT NOT NULL,
    id INT,
    tweet TEXT,
    created DATETIME,
    retweets INT,
    favs INT
  );

COMMIT;
