package main

import (
	"fmt"
	"log"

	"mykorelay/internal/connectors/fake"
	"mykorelay/internal/db"
)

func main() {
	fmt.Println("hello bitchez!")
	log.Println("starting relay daemon...")
	
	store := db.NewMemoryDB() //store layer to keep DB type swappable

	connector := fake.NewFakeConnector()

	go connector.Start(store) //run connector loop

	select {} //keepalive
	// add simple ticker to log a message forever?
}
