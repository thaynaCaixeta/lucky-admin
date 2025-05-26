-- +migrate Up

CREATE TABLE rounds (
    id UUID PRIMARY KEY,
    game_id UUID NOT NULL,
    numbers INTEGER[] NOT NULL,
    FOREIGN KEY(game_id) REFERENCES games(id)
);

-- +migrate Down

DROP TABLE rounds;
