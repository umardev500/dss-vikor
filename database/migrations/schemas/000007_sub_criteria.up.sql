CREATE TABLE IF NOT EXISTS sub_criterias (
    id UUID PRIMARY KEY,
    criteria_id UUID NOT NULL,
    name VARCHAR(255) UNIQUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT NULL,
    version BIGINT DEFAULT 0,
    FOREIGN KEY (criteria_id) REFERENCES criterias(id)
);

CREATE TRIGGER sub_criteria_update_trigger
    BEFORE UPDATE
    ON sub_criterias
    FOR EACH ROW
        WHEN (
            NEW.name IS DISTINCT FROM OLD.name OR
            NEW.criteria_id IS DISTINCT FROM OLD.criteria_id
        )
EXECUTE FUNCTION update_function()