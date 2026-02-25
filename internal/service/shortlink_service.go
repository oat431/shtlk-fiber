package service

import (
	"context"
	"oat431/shtlk-fiber/internal/dto/response"
	"oat431/shtlk-fiber/internal/repository"
	"oat431/shtlk-fiber/pkg/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

type ShortLinkService interface {
	GetAllLinks(ctx context.Context) ([]response.ShortLinkDTO, error)
	GetLinkByCode(ctx context.Context, code string, linkType string) (*response.ShortLinkDTO, error)
	CreateRandomShortLink(ctx context.Context, originalURL string) (*response.ShortLinkDTO, error)
	CreateCustomShortLink(ctx context.Context, originalURL string, customCode string) (*response.ShortLinkDTO, error)
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

func (s shortLinkService) GetLinkByCode(ctx context.Context, code string, linkType string) (*response.ShortLinkDTO, error) {
	shortLink, err := s.repo.GetLinkByShortCode(ctx, code, linkType)
	if err != nil {
		log.Error("Error fetching short link by code: ", err)
		return nil, err
	}

	shortLinkDTO := &response.ShortLinkDTO{
		ShortLink:    shortLink.ShortURL,
		OriginalLink: shortLink.OriginalURL,
		LinkType:     string(shortLink.Type),
	}
	return shortLinkDTO, nil
}

func (s shortLinkService) CreateRandomShortLink(ctx context.Context, originalURL string) (*response.ShortLinkDTO, error) {
	shortName := utils.GenerateName()
	isUnique := false
	for !isUnique {
		existingLink, _ := s.repo.GetLinkByShortCode(ctx, shortName, "RANDOM")
		if existingLink == nil {
			isUnique = true
		} else {
			shortName = utils.GenerateName()
		}
	}

	shortLink, err := s.repo.CreateShortLink(ctx, originalURL, shortName, "RANDOM")
	if err != nil {
		return nil, err
	}

	shortLinkDTO := &response.ShortLinkDTO{
		ShortLink:    shortLink.ShortURL,
		OriginalLink: shortLink.OriginalURL,
		LinkType:     string(shortLink.Type),
	}
	return shortLinkDTO, nil
}

func (s shortLinkService) CreateCustomShortLink(ctx context.Context, originalURL string, customCode string) (*response.ShortLinkDTO, error) {
	shortName := customCode
	existLink, err := s.repo.GetLinkByShortCode(ctx, shortName, "CUSTOM")
	if existLink != nil {
		return nil, fiber.NewError(fiber.StatusConflict, "Custom short link already exists")
	}

	shortLink, err := s.repo.CreateShortLink(ctx, originalURL, shortName, "CUSTOM")
	if err != nil {
		return nil, err
	}

	shortLinkDTO := &response.ShortLinkDTO{
		ShortLink:    shortLink.ShortURL,
		OriginalLink: shortLink.OriginalURL,
		LinkType:     string(shortLink.Type),
	}
	return shortLinkDTO, nil
}
