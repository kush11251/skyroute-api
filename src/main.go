package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc="/routes", createRoute).Methods("POST")
	router.HandleFunc="/routes", getRoutes).Methods("GET")
	router.HandleFunc="/routes/{id}", getRoute).Methods("GET")
	router.HandleFunc="/routes/{id}", updateRoute).Methods("PUT")
	router.HandleFunc="/routes/{id}", deleteRoute).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}