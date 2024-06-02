package utils

import (
	"log"
	"net/http"
	"encoding/json"
)

type JSONError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(JSONError{Code: code, Message: message}); err != nil {
		log.Printf("Could not encode JSON response: %v", err)
	}
}

func LogAndRespondWithError(w http.ResponseWriter, code int, err error) {
	log.Printf("Error: %v", err)
}