package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type bookingRequest struct {
	UserID  string `json:"userId"`
	ClassID string `json:"classId"`
}

func BookClass(w http.ResponseWriter, r *http.Request) {
	var req bookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	booking := map[string]any{
		"id":        uuid.NewString(),
		"userId":    req.UserID,
		"classId":   req.ClassID,
		"status":    "confirmed",
		"createdAt": time.Now().UTC().Format(time.RFC3339),
	}

	writeJSON(w, http.StatusCreated, map[string]any{
		"message": "Booking successful",
		"booking": booking,
	})
}
