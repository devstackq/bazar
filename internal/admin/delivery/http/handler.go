package v1

import (
	"github.com/devstackq/bazar/internal/admin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	useCases admin.UseCases
	logger   *logrus.Logger
}

func NewHandler(useCases admin.UseCases, logger *logrus.Logger) *Handler {
	return &Handler{useCases: useCases, logger: logger}
}
