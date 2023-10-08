-- +goose Up

CREATE TABLE requests(
    request_id serial NOT NULL,
    nodeID UUID NOT NULL,
    responsetime BIGINT,
    successful bool NOT NULL,
    last_fetched_at TIMESTAMPTZ,
    CONSTRAINT fk_node
    FOREIGN KEY(nodeID)
	  REFERENCES nodes(nodeID)
);

-- +goose Down
DROP TABLE IF EXISTS requests;