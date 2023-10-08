package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (apiCfg *apiConfig) NodeProvider(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, 200, "")
}

func (apiCfg *apiConfig) Best(w http.ResponseWriter, r *http.Request) {
	ChainID := chi.URLParam(r, "chain_id")
	nodes, err := apiCfg.db.GetBestList(context.Background(), ChainID)
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("ERROR: %+v", err))
		return
	}

	RespondWithJSON(w, 200, DBNodeToNode(nodes))
}
