package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type Class struct {
	ID         string  `json:"id"`
	Title      string  `json:"title"`
	Instructor string  `json:"instructor"`
	Location   string  `json:"location"`
	Price      float64 `json:"price"`
}

func GetAllClasses(w http.ResponseWriter, _ *http.Request) {
	sampleClasses := []Class{
		{
			ID:         "123",
			Title:      "Beginner Painting",
			Instructor: "Jane Doe",
			Location:   "Online",
			Price:      25,
		},
	}

	writeJSON(w, http.StatusOK, sampleClasses)
}

func CreateClass(w http.ResponseWriter, r *http.Request) {
	var payload map[string]any
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	payload["id"] = uuid.NewString()
	writeJSON(w, http.StatusCreated, map[string]any{
		"message": "Class created",
		"class":   payload,
	})
}
