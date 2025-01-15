package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Namith667/GoQuick/internal/db"
	"github.com/gorilla/mux"
)

func main() {

	database := db.Connect()

	db.RunMigrations(database)

	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(" server Healthy"))
	})

	fmt.Println("starting Server :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
