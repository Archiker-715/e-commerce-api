CREATE TABLE users (
    user_id UUID PRIMARY KEY,
    login VARCHAR(30) NOT NULL UNIQUE,
    password BYTEA,
    inserted_by UUID,
    inserted TIMESTAMP,
    updated_by UUID,
    updated TIMESTAMP
);