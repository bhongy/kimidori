CREATE TABLE IF NOT EXISTS sessions (
  id         VARCHAR(64) PRIMARY KEY,
  user_id    VARCHAR(64) REFERENCES users(id),
  created_at TIMESTAMP NOT NULL
);

---- create above / drop below ----

DROP TABLE IF EXISTS sessions;
