package main

import (
	"fmt"
	"net/http"
	"server/router"

	"log"
)

func main() {

	r := router.Router()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
