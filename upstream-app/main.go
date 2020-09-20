package main

import (
	"fmt"
	"log"
	"net/http"
)

var isHealthy bool

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if isHealthy {
		fmt.Fprintf(w, "Status is healthy\n")
	} else {
		fmt.Fprintf(w, "Status is unhealthy\n")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func unhealthyHandler(w http.ResponseWriter, r *http.Request) {
	isHealthy = false
	fmt.Fprintf(w, "Status is unhealthy\n")
	w.WriteHeader(http.StatusInternalServerError)
}

func healthyHandler(w http.ResponseWriter, r *http.Request) {
	isHealthy = true
	fmt.Fprintf(w, "Status is healthy\n")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", healthCheckHandler)
	http.HandleFunc("/unhealthy", unhealthyHandler)
	http.HandleFunc("/healthy", healthyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}