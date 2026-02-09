CREATE TABLE users (
    user_id UUID NOT NULL,
    login VARCHAR(30) NOT NULL,
    password BYTEA
);