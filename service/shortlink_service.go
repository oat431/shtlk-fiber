package service

import (
	"context"
	"oat431/shtlk-fiber/payload/response"
	"oat431/shtlk-fiber/repository"
)

type ShortLinkService interface {
	GetAllLinks(ctx context.Context) ([]response.ShortLinkDTO, error)
}

type shortLinkService struct {
	repo repository.ShortLinkRepository
}

func NewShortLinkService(repo repository.ShortLinkRepository) ShortLinkService {
	return &shortLinkService{repo: repo}
}

func (s shortLinkService) GetAllLinks(ctx context.Context) ([]response.ShortLinkDTO, error) {
	shortLinks, err := s.repo.GetAllShortLink(ctx)
	if err != nil {
		return nil, err
	}

	var shortLinkDTOs []response.ShortLinkDTO
	for _, sl := range shortLinks {
		shortLinkDTO := response.ShortLinkDTO{
			ShortLink:    sl.ShortURL,
			OriginalLink: sl.OriginalURL,
			LinkType:     string(sl.Type),
		}
		shortLinkDTOs = append(shortLinkDTOs, shortLinkDTO)
	}
	return shortLinkDTOs, nil
}
