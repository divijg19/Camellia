package webui

import (
	"net/http"

	"github.com/a-h/templ"
)

func classesFragmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	templ.Handler(classesFragment()).ServeHTTP(w, r)
}

func bookingFragmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}

	userID := r.FormValue("userId")
	if userID == "" {
		userID = "guest-user"
	}
	classID := r.FormValue("classId")
	if classID == "" {
		classID = "unknown-class"
	}

	templ.Handler(bookingResult(userID, classID)).ServeHTTP(w, r)
}
