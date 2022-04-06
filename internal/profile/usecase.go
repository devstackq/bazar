package profile

import (
	"context"

	"github.com/devstackq/bazar/internal/models"
)

type ProfileUseCasesInterface interface {
	GetBioByUserID(context.Context, int) (*models.Profile, error)
}
