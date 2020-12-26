CREATE TABLE IF NOT EXISTS sessions (
  id         SERIAL PRIMARY KEY,
  uuid       VARCHAR(64) NOT NULL UNIQUE,
  user_id    INTEGER REFERENCES users(id),
  created_at TIMESTAMP NOT NULL
);
