package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/pratika/VideoPlatform/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/video/{videoId}", handlers.IncrementViewCount).Methods("PUT")
	r.HandleFunc("/video/{videoId}", handlers.GetViews).Methods("GET")
	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
