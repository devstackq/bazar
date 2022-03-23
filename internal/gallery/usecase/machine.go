package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/models"
	"github.com/google/uuid"
)

type MachineUseCase struct {
	machineRepo gallery.MachineRepoInterface
}

func MachineUseCaseInit(
	machineRepo gallery.MachineRepoInterface,
) gallery.MachineUseCaseInterface {
	return MachineUseCase{
		machineRepo: machineRepo,
	}
}

func (muc MachineUseCase) Create(m *models.Machine) (int, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	vin := uuid.New()
	m.VIN = vin.String()
	return muc.machineRepo.Create(ctx, m)
}

func (muc MachineUseCase) GetMachineByID(id int) (*models.Machine, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	machine, err := muc.machineRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return machine, nil
}

func (muc MachineUseCase) GetRelevantMachines() ([]*models.Machine, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	list, err := muc.machineRepo.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (muc MachineUseCase) GetListMachineByUserID(id int) ([]*models.Machine, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	list, err := muc.machineRepo.GetListByUserID(ctx, id)
	if err != nil {
		return nil, err
	}
	return list, nil
}
