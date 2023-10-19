CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
    id VARCHAR(64) PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name VARCHAR(512) NOT NULL,
    email VARCHAR(512) NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    created_by VARCHAR(64) NOT NULL,
    updated_at timestamptz DEFAULT now(),
    updated_by VARCHAR(64),
    deleted_at timestamptz,
    deleted_by VARCHAR(64),
    status VARCHAR(16) NOT NULL DEFAULT 'disabled'
);