-- name: CreateReq :one
INSERT INTO requests(
    nodeID,
    responsetime,
    successful,
    last_fetched_at
)VALUES(
    $1,$2,$3,$4
)
RETURNING *;



-- name: GetNextNodestoFetch :many
SELECT * FROM nodes
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;

-- name: MarkNodesASFetched :one
UPDATE nodes
SET
    last_fetched_at = NOW(),
    updated_at = NOW()
WHERE 
    nodeID = $1
RETURNING *;