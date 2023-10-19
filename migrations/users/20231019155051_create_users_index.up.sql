CREATE INDEX IF NOT EXISTS users_status_idx ON users ( status );
CREATE UNIQUE INDEX IF NOT EXISTS users_email_unique_idx ON users ( email );
