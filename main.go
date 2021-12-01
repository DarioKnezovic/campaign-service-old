package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", testHome)

	log.Println("Server is running!")
	http.ListenAndServe(":4000", router)
}