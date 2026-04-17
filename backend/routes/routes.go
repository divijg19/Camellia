package routes

import (
	"net/http"

	"github.com/divijg19/camellia/backend/controllers"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/classes", classesHandler)
	mux.HandleFunc("/api/bookings", bookingsHandler)
	mux.HandleFunc("/api/users", usersHandler)
}

func classesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		controllers.GetAllClasses(w, r)
	case http.MethodPost:
		controllers.CreateClass(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func bookingsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		controllers.BookClass(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		controllers.CreateUser(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
