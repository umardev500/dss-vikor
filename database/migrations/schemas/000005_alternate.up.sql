CREATE TABLE alternates (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    role_id UUID NOT NULL,
    str VARCHAR(255) UNIQUE,
    experience INT NOT NULL,
    dob DATE NOT NULL,
    address VARCHAR(255) NOT NULL,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz DEFAULT NULL,
    FOREIGN KEY (role_id) REFERENCES roles(id)
);

CREATE INDEX idx_alternate_name ON alternates(name);

CREATE INDEX idx_alternate_role_id ON alternates(role_id);

CREATE INDEX idx_alternate_experience ON alternates(experience);

CREATE TRIGGER alternate_update_trigger
    BEFORE UPDATE
    ON alternates
    FOR EACH ROW
        WHEN (
            NEW.name IS DISTINCT FROM OLD.name OR
            NEW.role_id IS DISTINCT FROM OLD.role_id OR
            NEW.str IS DISTINCT FROM OLD.str OR
            NEW.experience IS DISTINCT FROM OLD.experience OR
            NEW.dob IS DISTINCT FROM OLD.dob OR
            NEW.address IS DISTINCT FROM OLD.address
        )
EXECUTE FUNCTION update_function();