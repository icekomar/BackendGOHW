package main

import (
	"dz/api"
	"log"
)

func main() {
	api := api.New("localhost:8080")
	api.FillEndpoints()
	log.Fatal(api.ListenAndServe())
}
