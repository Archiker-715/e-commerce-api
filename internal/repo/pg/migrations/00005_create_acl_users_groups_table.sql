CREATE TABLE users_groups (
    user_id UUID NOT NULL,
    group_id UUID NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY (user_id)
        REFERENCES users (user_id)
        ON DELETE CASCADE,
    CONSTRAINT fk_group
        FOREIGN KEY (group_id)
        REFERENCES acl_groups (group_id)
        ON DELETE CASCADE
);