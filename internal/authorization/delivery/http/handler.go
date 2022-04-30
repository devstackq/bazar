package v1

import (
	auth "github.com/devstackq/bazar/internal/authorization"
	"github.com/devstackq/bazar/internal/config"
	"github.com/sirupsen/logrus"
)

// DI - for example mock
type Handler struct {
	authUseCases auth.AuthUseCaseInterface
	jwtUseCases  auth.JwtTokenUseCaseInterface
	logger       *logrus.Logger
	cfg          *config.Config
}

// for example unit test; mock service
func NewHandler(auth auth.AuthUseCaseInterface, jwt auth.JwtTokenRepositoryInterface, logger *logrus.Logger, cfg *config.Config) *Handler {
	return &Handler{authUseCases: auth, jwtUseCases: jwt, logger: logger, cfg: cfg}
}
