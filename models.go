package main

import "github.com/aminghafoory/nodeProviderProxy/internal/database"

type Node struct {
	nodeid          string
	Nodeurl         string `json:"node_url"`
	ChainID         string `json:"chain_id"`
	ChainName       string `json:"chain_name"`
	AvgResponsetime int    `json:"avg_response_time"`
	SuccessfulCount int    `json:"success_count"`
	FailureCount    int    `json:"failure_count"`
}

func DBNodeToNode(dbnodes []database.GetBestListRow) []Node {
	nodes := []Node{}
	for _, dbnode := range dbnodes {
		node := Node{
			nodeid:          string(dbnode.Nodeid.NodeID()),
			Nodeurl:         dbnode.Nodeurl,
			ChainID:         dbnode.ChainID,
			ChainName:       dbnode.ChainName.String,
			AvgResponsetime: int(dbnode.AvgResponsetime),
			SuccessfulCount: int(dbnode.SuccessfulCount),
			FailureCount:    int(dbnode.FailureCount),
		}
		nodes = append(nodes, node)
	}

	return nodes
}
