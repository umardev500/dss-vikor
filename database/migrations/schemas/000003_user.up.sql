CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY,
    email text NOT NULL UNIQUE,
    password text NOT NULL,
    status status_enum DEFAULT 'inactive'::status_enum,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz DEFAULT NULL,
    deleted_at timestamptz DEFAULT NULL,
    version int DEFAULT 0
);