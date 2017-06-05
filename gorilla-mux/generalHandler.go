package main

import (
	"fmt"
	"net/http"
)

// HomeHandler Simple Testing home function
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is HOME")
}

// HealthCheck end point
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
