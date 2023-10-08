package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marsharl JSON response : %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)

}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("responding with 5XX error: ", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}
	RespondWithJSON(w, code, errResponse{Error: msg})

}
