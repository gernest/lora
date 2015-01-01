CREATE TABLE session (
  session_key    CHAR(64)  NOT NULL,
  session_data   BYTEA,
  session_expiry TIMESTAMP NOT NULL,
  CONSTRAINT session_key PRIMARY KEY (session_key)
);