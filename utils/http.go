package utils

import (
	"fmt"
	"net/http"
)

// SendResponse writes the HTTP response to the response writer
func SendResponse(w http.ResponseWriter, body string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,Content-Type,Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PING")
	w.WriteHeader(statusCode)
	fmt.Fprintf(w, body)
}

// SendErrorResponse writes the HTTP error response to the response writer
func SendErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	SendResponse(w, fmt.Sprintf(`{"error": "%s"}`, err), statusCode)
}
