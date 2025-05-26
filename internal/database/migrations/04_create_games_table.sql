-- +migrate Up

CREATE TABLE games (
    id UUID PRIMARY KEY,
    num_rounds INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    closes_at TIMESTAMP NOT NULL,
    completion_status TEXT NOT NULL,
    created_by UUID NOT NULL,
    FOREIGN KEY(created_by) REFERENCES admins(id)
);

-- +migrate Down

DROP TABLE games;
