ALTER TABLE users
ADD column type VARCHAR(16) NOT NULL DEFAULT 'customer';,

CREATE INDEX IF NOT EXISTS users_type_idx ON users ( type );
