package main

import (
	"log"
	"net/http"

	"MyKorr/internal/api"
	"MyKorr/internal/connectors/fake"
	"MyKorr/internal/db"
)

func main() {
	log.Println("relay server accessible on http://localhost:5656")

	store := db.NewMemoryDB() //store layer to keep DB type swappable

	apiHandler := api.NewAPI(store)

	connector := fake.NewFakeConnector()
	go connector.Start(store) //run connector loop

	// routes
	http.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		apiHandler.HandleGetMessages(w, r)
	case http.MethodPost:
		apiHandler.HandlePostMessage(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(":5656", nil))

	select {} //keepalive

}
