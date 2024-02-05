CREATE OR REPLACE FUNCTION update_function()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = current_timestamp at time zone 'utc';
    NEW.version = OLD.version + 1;
    return NEW;
END;
$$ LANGUANGE plpgsql;