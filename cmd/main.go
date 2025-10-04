package main

import (
	"log"

	"github.com/PatrochR/disko/db"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatalln(err)
	}
}

