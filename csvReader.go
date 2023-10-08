package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aminghafoory/nodeProviderProxy/internal/database"
	"github.com/google/uuid"
)

func dataReader(db *database.Queries) {

	txtFile, err := os.Open("results-test-st.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer txtFile.Close()

	scanner := bufio.NewScanner(txtFile)
	fmt.Println("Data Importer started")
	for scanner.Scan() {
		url := scanner.Text()
		splittedURL := strings.Split(url, ",")
		if len(splittedURL) != 3 {
			continue
		}
		url = splittedURL[0]
		chainID := splittedURL[1]
		ChainName := splittedURL[2]

		db.CreateNode(context.Background(), database.CreateNodeParams{
			Nodeid:  uuid.New(),
			Nodeurl: url,
			ChainID: chainID,
			ChainName: sql.NullString{
				String: ChainName,
				Valid:  true,
			},
			LastFetchedAt: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}
}
