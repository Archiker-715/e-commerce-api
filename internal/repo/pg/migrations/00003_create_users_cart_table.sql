CREATE TABLE users_cart (
    user_id UUID NOT NULL,
    product_id BIGINT NOT NULL,
    name VARCHAR(500) NOT NULL,
    count BIGINT NOT NULL,
    unit_price BIGINT NOT NULL,
    total_price BIGINT NOT NULL,
    CONSTRAINT fk_product
        FOREIGN KEY (product_id)
        REFERENCES product (product_id)
        ON DELETE CASCADE
);