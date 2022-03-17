package v1

import (
	"github.com/devstackq/bazar/internal/gallery"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	useCases gallery.UseCases
	logger *logrus.Logger
}
func NewHandler(useCases gallery.UseCases, logger *logrus.Logger ) *Handler {
	return &Handler{useCases: useCases, logger: logger}
}


