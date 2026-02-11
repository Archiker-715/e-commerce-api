CREATE TABLE markets (
    market_id SERIAL PRIMARY KEY,
    market_name VARCHAR(255) NOT NULL UNIQUE,
    inserted_by UUID,
    inserted TIMESTAMP,
    updated_by UUID,
    updated TIMESTAMP
);
