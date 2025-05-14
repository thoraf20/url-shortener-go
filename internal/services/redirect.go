package services

import (
	"context"
	"github.com/thoraf20/url-shortner/internal/storage"
)

type RedirectService struct {
	store *storage.RedisStore
}

func NewRedirectService(store *storage.RedisStore) *RedirectService {
	return &RedirectService{store: store}
}

func (s *RedirectService) GetLongURL(ctx context.Context, tenantID, shortCode string) (string, error) {
	url, err := s.store.GetURL(ctx, tenantID, shortCode)
	if err != nil {
		return "", err
	}
	if url == nil {
		return "", nil
	}
	return url.LongUrl, nil
}