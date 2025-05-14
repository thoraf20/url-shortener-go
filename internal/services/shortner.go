package services

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/thoraf20/url-shortner/internal/models"
	"github.com/thoraf20/url-shortner/internal/storage"
	"time"
)

type ShortenerService struct {
	store * storage.RedisStore
}

func NewShortenerService(store *storage.RedisStore) *ShortenerService {
	return &ShortenerService{store: store}
}

func (s *ShortenerService) Shorten(ctx context.Context, req *models.ShortenRequest) (*models.ShortenResponse, error) {
	shortCode, err := generateShortCode()
	if err != nil {
		return nil, err
	}

	url := &models.URL{
		TenantID: req.TenantID,
		ShortCode: shortCode,
		LongUrl: req.LongUrl,
		CreatedAt: time.Now(),
	}

	if err := s.store.SaveURL(ctx, url); err != nil {
		return nil, err
	}

	return &models.ShortenResponse{
		ShortCode: shortCode,
		ShortURL: "https://" + req.TenantID + "./short/" + shortCode,
	}, nil
}

func generateShortCode() (string, error) {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(b)[:6], nil
}