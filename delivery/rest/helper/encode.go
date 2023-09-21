package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct {
	Message string `json:"message,omitempty"`
	Error   any    `json:"error,omitempty"`
}

type Success struct {
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func ErrorEncode(w http.ResponseWriter, err error) {
	resp := Error{
		Message: err.Error(),
		Error:   err,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Fatalf("failed encoded")
	}

}

func SuccessEncode(w http.ResponseWriter, data any, message string) {
	resp := Success{
		Data:    data,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Fatalf("failed encoded")
	}
}
