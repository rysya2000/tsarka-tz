package delivery

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(Mux *chi.Mux, r *Handler) {

	// Options
	Mux.Use(middleware.Logger)

	// healthcheck
	Mux.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("health"))
	})

	Mux.Post("/rest/substr/find", r.SubStrHandler)
	Mux.Post("/rest/email/check", r.EmailHandler)
	Mux.Get("/rest/self/find/{str}", r.SelfHandler)

}
