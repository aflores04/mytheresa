package main

import (
	"log"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request)  {
	log.Println("I'm alive!")
}

func main() {
	http.HandleFunc("/", Handle)

	_ = http.ListenAndServe(":8080", nil)
}