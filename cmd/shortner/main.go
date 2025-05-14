package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/thoraf20/url-shortner/internal/api"
	"github.com/thoraf20/url-shortner/internal/config"
	"github.com/thoraf20/url-shortner/internal/services"
	"github.com/thoraf20/url-shortner/internal/storage"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	store, err := storage.NewRedisStore(cfg.RedisAddr)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to redis")
	}

	shortener := services.NewShortenerService(store)
	redirect := services.NewRedirectService(store)
	handler := api.NewHandler(shortener, redirect)

	r := mux.NewRouter()
	r.HandleFunc("/shorten", handler.ShortenURL).Methods("POST")
	r.HandleFunc("/{tenant_id}/{short_code}", handler.Redirect).Methods("GET")

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Info().Msgf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal().Err(err).Msg("Server Failed to start")
	}
}