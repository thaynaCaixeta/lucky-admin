-- +migrate Up

CREATE TABLE admins (
    id UUID PRIMARY KEY,
    username TEXT NOT NULL,
    pass TEXT NOT NULL,
    created_at TIMESTAMP,
    is_active BOOL NOT NULL
);

-- +migrate Down

DROP TABLE admins;
