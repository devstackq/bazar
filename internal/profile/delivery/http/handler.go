package v1

import (
	"github.com/devstackq/bazar/internal/config"
	"github.com/devstackq/bazar/internal/profile"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	profileUseCases profile.ProfileUseCasesInterface
	// galleryUseCase gallery.UseCases
	logger *logrus.Logger
	cfg    *config.Config
}

func NewHandler(profile profile.ProfileUseCasesInterface, logger *logrus.Logger, cfg *config.Config) *Handler {
	return &Handler{profileUseCases: profile, logger: logger, cfg: cfg}
}
