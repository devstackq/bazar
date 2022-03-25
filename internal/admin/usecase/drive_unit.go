package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/admin"
	"github.com/devstackq/bazar/internal/models"
)

type DriveUnitUseCase struct {
	driveUnitRepo admin.DriveUnitRepoInterface
}

func DriveUnitUseCaseInit(r admin.DriveUnitRepoInterface) admin.DriveUnitUseCaseInterface {
	return DriveUnitUseCase{driveUnitRepo: r}
}

func (cuc DriveUnitUseCase) CreateDriveUnit(c *models.DriveUnit) (int, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.driveUnitRepo.Create(ctx, c)
}

func (cuc DriveUnitUseCase) GetDriveUnitByID(id int) (*models.DriveUnit, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.driveUnitRepo.GetByID(ctx, id)
}

func (cuc DriveUnitUseCase) GetListDriveUnit() ([]*models.DriveUnit, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.driveUnitRepo.GetList(ctx)
}
