package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	host      string
	isHealthy bool = true
	serviceName string
)

type Response struct {
	Host    string
	Service string
	Message string
}

func init() {
	host, _ = os.Hostname()
	serviceName = os.Getenv("SERVICE_NAME")
}

func NewResponse(message string) Response {
	return Response{
		Host:    host,
		Service: serviceName,
		Message: message,
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	b, _ := json.Marshal(NewResponse("Hello, World"))
	fmt.Fprintf(w, string(b))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if isHealthy {
		b, _ := json.Marshal(NewResponse("Status is healthy"))
		fmt.Fprintf(w, string(b))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		b, _ := json.Marshal(NewResponse("Status is unhealthy"))
		fmt.Fprintf(w, string(b))
	}
}

func unhealthyHandler(w http.ResponseWriter, r *http.Request) {
	isHealthy = false
	w.WriteHeader(http.StatusInternalServerError)
	b, _ := json.Marshal(NewResponse("Status is unhealthy"))
	fmt.Fprintf(w, string(b))
}

func healthyHandler(w http.ResponseWriter, r *http.Request) {
	isHealthy = true
	b, _ := json.Marshal(NewResponse("Status is healthy"))
	fmt.Fprintf(w, string(b))
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", healthCheckHandler)
	http.HandleFunc("/unhealthy", unhealthyHandler)
	http.HandleFunc("/healthy", healthyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
