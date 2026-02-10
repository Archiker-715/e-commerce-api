CREATE TABLE markets (
    market_id UUID NOT NULL,
    market_name VARCHAR(255) NOT NULL UNIQUE,
    inserted_by UUID,
    inserted TIMESTAMP,
    updated_by UUID,
    updated TIMESTAMP
);
