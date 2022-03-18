package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/models"
)

type DriveUnitUseCase struct {
	driveUnitRepo gallery.DriveUnitRepoInterface
}

func DriveUnitUseCaseInit(r gallery.DriveUnitRepoInterface) gallery.DriveUnitUseCaseInterface {
	return DriveUnitUseCase{driveUnitRepo: r}
}

func (cuc DriveUnitUseCase) CreateDriveUnit(c *models.DriveUnit) (int,error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.driveUnitRepo.Create(ctx, c)
}

func (cuc DriveUnitUseCase) GetDriveUnitByID(id int)(*models.DriveUnit, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.driveUnitRepo.GetByID(ctx, id)
}

func (cuc DriveUnitUseCase) GetListDriveUnit()([]*models.DriveUnit, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return cuc.driveUnitRepo.GetList(ctx)
}