-- +migrate Up

CREATE TABLE bets (
    id UUID PRIMARY KEY,
    round_id UUID NOT NULL REFERENCES rounds(id),
    client_id UUID NOT NULL REFERENCES clients(id),
    payment_status TEXT NOT NULL
);

-- +migrate Down

DROP TABLE bets;
