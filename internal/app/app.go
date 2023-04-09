package app

import (
	"log"
	"net/http"
	"rest/config"
	"rest/internal/delivery"
	"rest/internal/usecase"

	"github.com/go-chi/chi"
)

func Run(cfg *config.Config) {

	usecase := usecase.New()

	handler := delivery.New(usecase)

	//  HTTP Server
	Mux := chi.NewRouter()
	delivery.NewRouter(Mux, handler)

	log.Printf("Starting up on http://localhost:%s", cfg.HTTP.Port)

	log.Fatal(http.ListenAndServe(":"+cfg.HTTP.Port, Mux))

}
