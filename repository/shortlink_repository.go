package repository

import (
	"context"
	"oat431/shtlk-fiber/model"

	"github.com/jmoiron/sqlx"
)

type ShortLinkRepository interface {
	GetAllShortLink(ctx context.Context) ([]model.ShortLink, error)
}

type shortLinkRepository struct {
	db *sqlx.DB
}

func NewShortLinkRepository(db *sqlx.DB) ShortLinkRepository {
	return &shortLinkRepository{db: db}
}

func (s shortLinkRepository) GetAllShortLink(ctx context.Context) ([]model.ShortLink, error) {
	query := "SELECT id, url_original, url_short, link_type, created_at FROM tb_short_link"
	rows, err := s.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shortLinks []model.ShortLink
	for rows.Next() {
		var sl model.ShortLink
		if err := rows.Scan(&sl.ID, &sl.OriginalURL, &sl.ShortURL, &sl.Type, &sl.CreatedAt); err != nil {
			return nil, err
		}
		shortLinks = append(shortLinks, sl)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return shortLinks, nil
}
