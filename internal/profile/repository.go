package profile

import (
	"context"

	"github.com/devstackq/bazar/internal/models"
)

type ProfileRepositoryInterface interface {
	// GetBioByUserID(context.Context, int) (*models.Profile, error)
	GetUserByID(context.Context, int) (*models.User, error)
	// GetListMachineByUserID(context.Context, int) ([]*models.Machine, error)
}
