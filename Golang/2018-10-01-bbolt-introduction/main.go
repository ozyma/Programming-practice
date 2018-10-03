package main

import (
	"fmt"
	"log"

	bolt "go.etcd.io/bbolt"
)

func main() {
	fmt.Println("bbolt database introduction")
	db, err := bolt.Open("mytest.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database statistics: ", db.Stats())
	db.Close()
	fmt.Println("Please press enter to exit")
	fmt.Scanln()
}
