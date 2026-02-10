CREATE TABLE product_rules (
    product_id BIGINT NOT NULL,
    group_id UUID NOT NULL,
    rule VARCHAR(7) NOT NULL,
    CONSTRAINT fk_group
        FOREIGN KEY (group_id)
        REFERENCES acl_groups (group_id)
        ON DELETE CASCADE
);
