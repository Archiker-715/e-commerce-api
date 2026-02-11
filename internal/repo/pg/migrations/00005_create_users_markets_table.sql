CREATE TABLE users_markets (
    user_id UUID NOT NULL,
    market_id BIGINT NOT NULL,
    market_owner_user_id UUID,
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
        REFERENCES markets (market_id)
        ON DELETE CASCADE,
    CONSTRAINT fk_market_owner
        FOREIGN KEY (market_owner_user_id)
        REFERENCES users (user_id)
        ON DELETE CASCADE
);