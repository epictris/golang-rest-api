package main

import (
	"log"

	"tris.sh/go/api"
	"tris.sh/go/db"
)

func main() {
	database := db.Init()
	defer database.Close()
	if err := api.Init(database); err != nil {
		log.Fatal(err)
	}
}
