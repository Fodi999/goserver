package handlers

import (
	"encoding/json"
	"net/http"
)

// APIError представляет структуру ошибки API
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// WriteError отправляет JSON-ошибку клиенту
func WriteError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(APIError{Code: code, Message: message})
}

