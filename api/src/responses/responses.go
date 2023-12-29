package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON return a Json response to the requisition
func JSON(w http.ResponseWriter, statusCode int, dados any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(dados); err != nil {
		log.Fatal(err)
	}
}

// Error return a error in Json format
func Error(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Error string `json:"erro"`
	}{
		Error: err.Error(),
	})
}
