package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"github.com/thoraf20/url-shortner/internal/models"
	"github.com/thoraf20/url-shortner/internal/services"
)

type Handler struct {
	shortener *services.ShortenerService
	redirect *services.RedirectService
}

func NewHandler(shortener *services.ShortenerService, redirect *services.RedirectService) *Handler {
	return &Handler{ shortener: shortener, redirect: redirect}
}

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req models.ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	resp, err := h.shortener.Shorten(r.Context(), &req)
	if err != nil {
		log.Error().Err(err).Msg("Failed to shorten URL")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tenantID := vars["tenant_id"]
	shortCode := vars["short_code"]

	longURL, err := h.redirect.GetLongURL(r.Context(), tenantID, shortCode)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get long URL")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if longURL == "" {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}