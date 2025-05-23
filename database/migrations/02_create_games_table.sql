-- +migrate Up

CREATE TABLE games (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    closes_at TIMESTAMP NOT NULL,
    completion_status TEXT NOT NULL
);

-- +migrate Down

DROP TABLE games;
