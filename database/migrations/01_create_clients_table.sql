-- +migrate Up

CREATE TABLE clients (
    id UUID PRIMARY KEY,
    fullName TEXT NOT NULL,
    birth_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL
);

-- +migrate Down

DROP TABLE clients;
