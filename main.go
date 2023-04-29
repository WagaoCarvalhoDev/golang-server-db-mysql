package main

import (
	client "clientdao"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users/", client.UserHandler)
	log.Println("Executing...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
