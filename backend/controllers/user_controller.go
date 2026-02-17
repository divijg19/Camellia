package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var payload map[string]any
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	payload["id"] = uuid.NewString()
	writeJSON(w, http.StatusCreated, map[string]any{
		"message": "User registered",
		"user":    payload,
	})
}
