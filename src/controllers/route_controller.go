package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"skyroute-api/src/models"
)

func createRoute(w http.ResponseWriter, r *http.Request) {
	var route models.Route
	json.NewDecoder(r.Body).Decode(&route)
	// Save route to database
	w.WriteHeader(http.StatusCreated)
}

func getRoutes(w http.ResponseWriter, r *http.Request) {
	// Retrieve routes from database
	w.WriteHeader(http.StatusOK)
}

func getRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// Retrieve route from database
	w.WriteHeader(http.StatusOK)
}

func updateRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var route models.Route
	json.NewDecoder(r.Body).Decode(&route)
	// Update route in database
	w.WriteHeader(http.StatusOK)
}

func deleteRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	// Delete route from database
	w.WriteHeader(http.StatusNoContent)
}