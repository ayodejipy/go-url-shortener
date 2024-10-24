package service

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	db "rest/api/internals/db/sqlc"
	"rest/api/internals/dto"
	"rest/api/internals/logger"
	"rest/api/internals/middleware"

	"github.com/jackc/pgx/v5/pgtype"
)

type UrlService struct {
	Store  db.Store
	Logger *logger.Logger
}

func (s *UrlService) generateShortCode(url string) string {
	// hash the url
	algo := sha256.New()
	algo.Write([]byte(url))

	// conver to base 64
	code := base64.URLEncoding.EncodeToString(algo.Sum(nil))

	// get the first 8 characters to use a url code
	return code[:8]
}

func (s *UrlService) GetUrlByShortCode(ctx context.Context, shortCode string) (db.Url, error) {
	record, err := s.Store.GetUrlByCode(ctx, shortCode)
	if err != nil {
		s.Logger.Error("[s.Store.GetUrlByCode:] %v", err)
		return db.Url{}, errors.New("invalid url shortCode")
	}

	return record, nil
}

func (s *UrlService) GetUrlsByUser(ctx context.Context) ([]db.Url, error) {
	user := ctx.Value(middleware.UserKey).(db.GetUserRow)

	record, err := s.Store.GetUrlsByUser(ctx, user.ID)
	if err != nil {
		s.Logger.Error("[s.Store.GetUrlsByUser:] %v", err)
		return []db.Url{}, errors.New("missing user")
	}

	return record, nil
}

func (s *UrlService) ShortenLongUrl(ctx context.Context, payload dto.CreateShortPayload) (db.CreateUrlRow, error) {
	user := ctx.Value(middleware.UserKey).(db.GetUserRow)

	if payload.OriginalUrl == "" {
		return db.CreateUrlRow{}, errors.New("url field is required")
	}

	shortCode := s.generateShortCode(payload.OriginalUrl)

	createUrlPayload := db.CreateUrlParams{
		OriginalUrl: payload.OriginalUrl,
		ShortCode:   shortCode,
		UserID:      user.ID,
	}

	createdUrl, err := s.Store.CreateUrl(ctx, createUrlPayload)
	if err != nil {
		s.Logger.Error("[s.Store.CreateUrl]: %v", err)
		return db.CreateUrlRow{}, errors.New("unable to shorten url")
	}

	return createdUrl, nil
}

func (s *UrlService) UpdateUrl(ctx context.Context) error {
	return nil
}

func (s *UrlService) DeleteUrl(ctx context.Context, id pgtype.UUID) error {
	err := s.Store.DeleteUrl(ctx, id)
	if err != nil {
		s.Logger.Error("[s.Store.DeleteUrl:] %v", err)
		return errors.New("invalid ID")
	}

	return nil
}
