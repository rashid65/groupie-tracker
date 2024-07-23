package main

import (
	"log"
	"net/http"
	"groupie/Server"
)

func main() {
	groupie.Start()

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
