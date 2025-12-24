package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Healthy"))
	}).Methods("GET")

	log.Println("Server running on port :", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
