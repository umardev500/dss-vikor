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

-- create index of status
CREATE INDEX idx_user_status ON users(status);

-- Create trigger
CREATE TRIGGER user_update_trigger
    BEFORE UPDATE
    ON users
    FOR EACH ROW
        WHEN (
            NEW.email IS DISTINCT FROM OLD.email OR
            NEW.password IS DISTINCT FROM OLD.password OR
            NEW.status IS DISTINCT FROM OLD.status
        )
EXECUTE FUNCTION update_function();