-- Add migration script here
CREATE TABLE IF NOT EXISTS threads (
    id SERIAL PRIMARY KEY,
    public_id CHAR(23) UNIQUE NOT NULL DEFAULT concat('thread_', nanoid()),

    message VARCHAR(200) NOT NULL,
    
    created_at bigint NOT NULL DEFAULT EXTRACT(EPOCH FROM CURRENT_TIMESTAMP),
    updated_at bigint NOT NULL DEFAULT EXTRACT(EPOCH FROM CURRENT_TIMESTAMP)
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = EXTRACT(EPOCH FROM CURRENT_TIMESTAMP);
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_users_updated_at BEFORE UPDATE
ON threads FOR EACH ROW EXECUTE PROCEDURE
update_updated_at_column();
