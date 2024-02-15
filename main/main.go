package main

import (
	"dz/pkg/api"
	"log"
)

func main() {
	a := api.New("localhost:8080")
	a.FillEndpoints()
	log.Fatal(a.ListenAndServe())
}
