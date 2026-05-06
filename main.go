package main

import (
	"log"
	"os"

	"sprint-final/pkg/db"
	"sprint-final/pkg/server"
)

func main() {
	dbFile := "scheduler.db"

	if f := os.Getenv("TODO_DBFILE"); f != "" {
		dbFile = f
	}

	if err := db.Init(dbFile); err != nil {
		log.Fatal(err)
	}
	defer db.DB.Close()

	server.Run()
}
