package utils

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

type MyError struct {
	Location string `json:"location"`
	Message  string `json:"msg"`
	Parametr string `json:"param"`
	Value    string `json:"value"`
}

type ErrResp struct {
	Errors []MyError `json:"errors"`
}

func JSON(w http.ResponseWriter, body interface{}, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(body)

	if err != nil {
		http.Error(w, "err JSON", http.StatusInternalServerError)
		return
	}
}

func ErrorJSON(w http.ResponseWriter, text string, code int) {
	message := &Message{
		Message: text,
	}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(message)

	if err != nil {
		http.Error(w, "err JSON", http.StatusInternalServerError)
		return
	}
}
