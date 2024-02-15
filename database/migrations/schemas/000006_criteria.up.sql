CREATE TABLE IF NOT EXISTS criterias (
    id UUID PRIMARY KEY,
    name TEXT UNIQUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT NULL,
    version INT DEFAULT 0
);

CREATE TRIGGER criteria_update_trigger
    BEFORE UPDATE
    ON criterias
    FOR EACH ROW
        WHEN (
            NEW.name IS DISTINCT FROM OLD.name
        )
EXECUTE FUNCTION update_function();

