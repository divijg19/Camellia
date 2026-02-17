package webui

import (
	"net/http"

	"github.com/a-h/templ"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("/", templ.Handler(homePage()))
	mux.HandleFunc("/web/classes", classesFragmentHandler)
	mux.HandleFunc("/web/book", bookingFragmentHandler)
}
