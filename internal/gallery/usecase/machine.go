package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/gallery/models"
)

type MachineUseCase struct {
	machineRepo gallery.MachineRepo
	categoryRepo gallery.CategoryRepo
	transRepo gallery.TransmissionRepo
}

func MachineUseCaseInit(mr gallery.MachineRepo, cr gallery.CategoryRepo, tr gallery.TransmissionRepo) gallery.MachineUseCase {
	return MachineUseCase{machineRepo: mr, categoryRepo: cr, transRepo: tr}
}

func (muc MachineUseCase) Create(m *models.Machine) (int,error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return muc.machineRepo.Create(ctx, m)
}

func (muc MachineUseCase) GetMachineByID(id int) ( *models.Machine, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	machine,err :=  muc.machineRepo.GetMachineByID(ctx,id)
	if err != nil {
		return nil, err
	}
	//get, set category name
	category, err := muc.categoryRepo.GetCategoryByID(ctx, machine.Category.ID)
	if err != nil {
		return nil, err
	}
	machine.Category.Name = category.Name

	trans, err := muc.transRepo.GetTransByID(ctx, machine.Transmission.ID)
	if err != nil {
		return nil, err
	}
	machine.Transmission.Name = trans.Name
//todo: get etc rel table data ;
	return machine, nil
}

func (muc MachineUseCase) GetRelevantMachines() ( []*models.Machine, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	list ,err :=  muc.machineRepo.GetRelevantMachines(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}