package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aminghafoory/nodeProviderProxy/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	db *database.Queries
}

func main() {

	portString, dbURL, dataReaderEnabeld, dataReaderTimeBetween := GetEnvVariables()
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can NOT connect to the database")
	}
	apiCfg := apiConfig{
		db: database.New(conn),
	}

	if dataReaderEnabeld == "true" {
		dataReader(apiCfg.db)
	}
	go startTesting(apiCfg.db, 50, time.Minute*time.Duration(dataReaderTimeBetween))

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Get("/best/{chain_id}", apiCfg.Best)
	fmt.Println("started on: ", portString)
	http.ListenAndServe(":"+portString, r)
}

func GetEnvVariables() (string, string, string, int64) {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is NOTfound in the .env file")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL not found in .env file: ")
	}

	dataReaderEnabeld := os.Getenv("ENABLED")
	if dataReaderEnabeld == "" {
		log.Fatal("ENABLED not found in .env file: ")
	}

	dataReaderTimeBetween := os.Getenv("TIMEBETWEEN")
	if dataReaderEnabeld == "" {
		log.Fatal("TIMEBETWEEN not found in .env file: ")
	}
	dataReaderTimeBetweenInt, err := strconv.ParseInt(dataReaderTimeBetween, 10, 64)
	if err != nil {
		log.Fatal("can not Parse TIMEBETWEEN")
	}

	return portString, dbURL, dataReaderEnabeld, dataReaderTimeBetweenInt

}
