CREATE TABLE roles (
    id uuid PRIMARY KEY,
    name text NOT NULL UNIQUE,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz DEFAULT NULL,
    version int DEFAULT 0
);


CREATE TRIGGER role_update_trigger
    BEFORE UPDATE
    ON roles
    FOR EACH ROW
        WHEN (
            NEW.name IS DISTINCT FROM OLD.name
        )
EXECUTE FUNCTION update_function();

