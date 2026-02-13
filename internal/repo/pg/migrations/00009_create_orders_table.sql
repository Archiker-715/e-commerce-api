CREATE TABLE orders (
    order_id TEXT,
    user_id UUID NOT NULL,
    order_price uint NOT NULL,
    products JSONB NOT NULL,
    temp BOOLEAN NOT NULL,
    inserted_by UUID,
    inserted TIMESTAMP,
    updated_by UUID,
    updated TIMESTAMP,
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users (user_id)
        ON DELETE CASCADE
);
