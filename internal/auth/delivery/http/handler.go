package v1

import (
	"github.com/devstackq/bazar/internal/auth"
	"github.com/devstackq/bazar/internal/config"
	"github.com/sirupsen/logrus"
)

//DI - for example mock
type Handler struct {
	useCases auth.UseCase
	logger *logrus.Logger
	cfg *config.Config
}

//for example unit test; mock service
func NewHandler(useCase auth.UseCase,  logger *logrus.Logger, cfg *config.Config) *Handler {
	return &Handler{useCases: useCase, logger: logger, cfg: cfg}
}
