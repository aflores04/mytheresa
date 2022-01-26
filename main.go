package main

import (
	"log"
	"net/http"
)

func Handle(w http.ResponseWriter, r *http.Request)  {
	log.Println("hola")
}

func main() {
	http.HandleFunc("/", Handle)

	_ = http.ListenAndServe(":8080", nil)
}