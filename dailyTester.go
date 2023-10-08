package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/aminghafoory/nodeProviderProxy/internal/database"
)

func TestNodesResponseTime(node database.Node, db *database.Queries, wg *sync.WaitGroup) {
	defer wg.Done()
	startTime := time.Now()
	_, err := getchainid(node.Nodeurl)
	endTime := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	if err != nil {
		db.CreateReq(context.Background(), database.CreateReqParams{
			Nodeid:        node.Nodeid,
			Responsetime:  sql.NullInt64{},
			Successful:    false,
			LastFetchedAt: endTime,
		})

	}
	responseTime := time.Since(startTime).Milliseconds()

	db.CreateReq(context.Background(), database.CreateReqParams{
		Nodeid: node.Nodeid,
		Responsetime: sql.NullInt64{
			Int64: responseTime,
			Valid: true,
		},
		Successful:    true,
		LastFetchedAt: endTime,
	})

	db.MarkNodesASFetched(context.Background(), node.Nodeid)

}

func startTesting(
	db *database.Queries,
	concurrency int32,
	timeBetweenRequests time.Duration) {

	log.Printf("Scraping on %v goroutine every %s duration", concurrency, timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		log.Printf("now")
		nodes, err := db.GetNextNodestoFetch(context.Background(), concurrency)
		if err != nil {
			log.Panicln("error fetching nodes", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, node := range nodes {
			wg.Add(1)

			go TestNodesResponseTime(node, db, wg)
		}

		wg.Wait()
	}

}

func getchainid(url string) (string, error) {
	//{"jsonrpc":"2.0","method":"net_version","params":[],"id":67}
	type RPCData struct {
		JSONRPC string   `json:"jsonrpc"`
		Method  string   `json:"method"`
		Params  []string `json:"params"`
		ID      int      `json:"id"`
	}
	//{"jsonrpc":"2.0","id":67,"result":"11155111"}
	type RPCResponse struct {
		JSONRPC string `json:"jsonrpc"`
		ID      int    `json:"id"`
		Result  string `json:"result"`
	}

	newData := RPCData{
		JSONRPC: "2.0",
		Method:  "net_version",
		Params:  []string{},
		ID:      67,
	}

	jsonData, err := json.Marshal(newData)
	if err != nil {
		return "", err
	}

	http.DefaultClient.Timeout = 2 * time.Second
	resp, err := http.Post(fmt.Sprintf("http://%s:8545", url), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	RPCResp := RPCResponse{}
	jsonDecoder := json.NewDecoder(resp.Body)

	err = jsonDecoder.Decode(&RPCResp)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return RPCResp.Result, nil

}
