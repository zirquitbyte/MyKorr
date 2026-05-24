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
	
var store db.MessageStore

sqlite, err := db.NewSQLiteStore("data/messages.db")
if err != nil {
	log.Fatal(err)
}

store = sqlite

defer store.Close()

if err := store.Init(); err != nil {
	log.Fatal(err)
}

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
