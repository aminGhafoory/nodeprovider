-- +goose Up

CREATE TABLE nodes (
    nodeID UUID PRIMARY KEY,
    nodeURL VARCHAR(64) NOT NULL,
    chain_ID VARCHAR(64) NOT NULL,
    chain_name VARCHAR(64),
    last_fetched_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS nodes;