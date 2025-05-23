-- +migrate Up

CREATE TABLE rounds (
    id UUID PRIMARY KEY,
    game_id UUID REFERENCES games(id),
    numbers INTEGER[] NOT NULL
);

-- +migrate Down

DROP TABLE rounds;
