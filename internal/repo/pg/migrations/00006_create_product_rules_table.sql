CREATE TABLE product_rules (
    product_id BIGINT NOT NULL,
    market_id BIGINT NOT NULL,
    rule VARCHAR(3),
    inserted_by UUID,
    inserted TIMESTAMP,
    updated_by UUID,
    updated TIMESTAMP,
    CONSTRAINT fk_market
        FOREIGN KEY (market_id)
        REFERENCES markets (market_id)
        ON DELETE CASCADE
);
