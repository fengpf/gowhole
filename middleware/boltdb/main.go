package main

import (
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Printf(db.Path())
}
