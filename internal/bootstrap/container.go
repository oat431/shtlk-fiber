package bootstrap

import (
	"oat431/shtlk-fiber/internal/controller"
	"oat431/shtlk-fiber/internal/repository"
	"oat431/shtlk-fiber/internal/service"

	"github.com/gofiber/fiber/v3/log"
	"github.com/jmoiron/sqlx"
)

type AppContainer struct {
	ShortLinkController *controller.ShortLinkController
	RedirectController  *controller.RedirectController
}

func NewAppContainer(db *sqlx.DB) *AppContainer {
	shortLinkRepo := repository.NewShortLinkRepository(db)

	shortLinkService := service.NewShortLinkService(shortLinkRepo)

	shortLinkController := controller.NewShortLinkController(shortLinkService)
	redirectController := controller.NewRedirectController(shortLinkService)

	log.Info("Starting Bootstrap Container...")
	return &AppContainer{
		ShortLinkController: shortLinkController,
		RedirectController:  redirectController,
	}
}
