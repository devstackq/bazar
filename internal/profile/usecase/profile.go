package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/models"
)

func (puc ProfileUseCase) GetBioByUserID(ctx context.Context, id int) (*models.Profile, error) {
	var result models.Profile
	bio, err := puc.profileRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	// machines, err := puc.profileRepo.GetListMachineByUserID(ctx, id)
	// if err != nil {
	// 	return nil, err
	// }
	result.Bio = bio

	return &result, nil
}
