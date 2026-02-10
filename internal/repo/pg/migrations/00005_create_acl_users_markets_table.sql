CREATE TABLE users_markets (
    user_id UUID NOT NULL,
    market_id UUID NOT NULL,
    inserted_by UUID,
    inserted TIMESTAMP,
    updated_by UUID,
    updated TIMESTAMP,
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users (user_id)
        ON DELETE CASCADE,
    CONSTRAINT fk_market
        FOREIGN KEY (market_id)
        REFERENCES acl_markets (market_id)
        ON DELETE CASCADE
);