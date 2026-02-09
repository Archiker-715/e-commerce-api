CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    category TEXT NOT NULL,
    price BIGINT NOT NULL,
    count BIGINT NOT NULL,
    active BOOLEAN NOT NULL,
    options JSONB,
    article TEXT UNIQUE,
    inserted_by UUID,
    inserted TIMESTAMP,
    updated_by UUID,
    updated TIMESTAMP
);