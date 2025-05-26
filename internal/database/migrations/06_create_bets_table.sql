-- +migrate Up

CREATE TABLE bets (
    id UUID PRIMARY KEY,
    round_id UUID NOT NULL,
    client_id UUID NOT NULL,
    payment_status TEXT NOT NULL,
    FOREIGN KEY(round_id) REFERENCES rounds(id),
    FOREIGN KEY(client_id) REFERENCES clients(id)
);

-- +migrate Down

DROP TABLE bets;
