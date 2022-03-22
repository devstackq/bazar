package usecase

import (
	"github.com/devstackq/bazar/internal/profile"
	// "github.com/devstackq/bazar/internal/profile/repository/psql"
)


type ProfileUseCase struct {
	profileRepo       profile.ProfileRepositoryInterface
}

func ProfileUseCaseInit(profileRepo profile.ProfileRepositoryInterface) profile.ProfileUseCasesInterface {
	return &ProfileUseCase{
		profileRepo:       profileRepo,
	}
}
