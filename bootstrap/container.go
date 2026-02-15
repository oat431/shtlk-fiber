package bootstrap

import (
	"oat431/shtlk-fiber/controller"
	"oat431/shtlk-fiber/repository"
	"oat431/shtlk-fiber/service"

	"github.com/gofiber/fiber/v3/log"
	"github.com/jmoiron/sqlx"
)

type AppContainer struct {
	ShortLinkController *controller.ShortLinkController
}

func NewAppContainer(db *sqlx.DB) *AppContainer {
	shortLinkRepo := repository.NewShortLinkRepository(db)

	shortLinkService := service.NewShortLinkService(shortLinkRepo)

	shortLinkController := controller.NewShortLinkController(shortLinkService)

	log.Info("Starting Bootstrap Container...")
	return &AppContainer{
		ShortLinkController: shortLinkController,
	}
}
