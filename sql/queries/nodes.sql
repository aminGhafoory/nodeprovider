-- name: GetBestList :many
SELECT 
    nodes.nodeID,
    nodeURL,
    chain_ID,
    chain_name,
    avg(responsetime) as avg_responsetime,
	sum(case successful when true then 1 else 0 end) as successful_count,
	sum(case successful when false then 1 else 0 end) as failure_count
FROM
    nodes JOIN requests ON nodes.nodeID = requests.nodeID
where 
	chain_ID = $1  
GROUP BY 
	nodes.nodeID
ORDER BY 
	avg_responsetime ASC;
	
	
	


-- name: CreateNode :one
INSERT INTO nodes(
    nodeID ,
    nodeURL ,
    chain_ID ,
    chain_name,
    last_fetched_at ,
    created_at ,
    updated_at 
)VALUES(
    $1,$2,$3,$4,$5,$6,$7
)
RETURNING *;
