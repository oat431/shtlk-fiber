package model

import (
	"time"
)

type ShortLink struct {
	ID          string    `db:"id" json:"id"`
	OriginalURL string    `db:"url_original" json:"url_original"`
	ShortURL    string    `db:"url_short" json:"url_short"`
	Type        LinkType  `db:"link_type" json:"link_type"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
