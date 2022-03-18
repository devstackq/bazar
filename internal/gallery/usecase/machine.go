package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/gallery"
	"github.com/devstackq/bazar/internal/models"
	"github.com/google/uuid"
)

type MachineUseCase struct {
	machineRepo gallery.MachineRepoInterface
	categoryRepo gallery.CategoryRepoInterface
	transRepo gallery.TransmissionRepoInterface
	countryRepo gallery.CountryRepoInterface
	cityRepo gallery.CityRepoInterface
	stateRepo gallery.StateRepoInterface
	brandRepo gallery.BrandRepoInterface
	modelRepo gallery.ModelRepoInterface
	fuelRepo gallery.FuelRepoInterface
	driveUnit gallery.DriveUnitRepoInterface
	bodyRepo gallery.BodyTypeRepoInterface
	colorRepo gallery.ColorRepoInterface
	// userRepo gallery.UserRepoInterface
}

func MachineUseCaseInit(
	machineRepo gallery.MachineRepoInterface, 
	categoryRepo gallery.CategoryRepoInterface, 
	transRepo gallery.TransmissionRepoInterface, 
	countryRepo gallery.CountryRepoInterface,
	cityRepo gallery.CityRepoInterface,
	stateRepo gallery.StateRepoInterface,
	brandRepo gallery.BrandRepoInterface,
	modelRepo gallery.ModelRepoInterface,
	fuelRepo gallery.FuelRepoInterface,
	driveUnit gallery.DriveUnitRepoInterface,
	bodyRepo gallery.BodyTypeRepoInterface,
	colorRepo gallery.ColorRepoInterface,
	) gallery.MachineUseCaseInterface {

	return MachineUseCase{
		machineRepo: machineRepo, categoryRepo: categoryRepo,transRepo: transRepo,
		countryRepo: countryRepo,  cityRepo: cityRepo,  stateRepo: stateRepo,
		brandRepo: brandRepo,  modelRepo: modelRepo,  fuelRepo: fuelRepo, 
		driveUnit: driveUnit, bodyRepo: bodyRepo, colorRepo: colorRepo,
	}
}

func (muc MachineUseCase) Create(m *models.Machine) (int,error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	vin := uuid.New()
	m.VIN = vin.String()
	return muc.machineRepo.Create(ctx, m)
}

func (muc MachineUseCase) GetMachineByID(id int) ( *models.Machine, error) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	machine,err :=  muc.machineRepo.GetByID(ctx,id)
	if err != nil {
		return nil, err
	}
	//get, set category name
	category, err := muc.categoryRepo.GetByID(ctx, machine.Category.ID)
	if err != nil {
		return nil, err
	}
	machine.Category.Name = category.Name

	trans, err := muc.transRepo.GetByID(ctx, machine.Transmission.ID)
	if err != nil {
		return nil, err
	}
	
	machine.Transmission.Name = trans.Name

	country, err := muc.countryRepo.GetByID(ctx, machine.Country.ID)
	if err != nil {
		return nil, err
	}
	machine.Country.Name = country.Name

	city, err := muc.cityRepo.GetByID(ctx, machine.City.ID)
	if err != nil {
		return nil, err
	}
	machine.City.Name = city.Name

	state, err := muc.stateRepo.GetByID(ctx, machine.State.ID)
	if err != nil {
		return nil, err
	}
	machine.State.Name = state.Name


	brand, err := muc.brandRepo.GetByID(ctx, machine.Brand.ID)

	if err != nil {
		return nil, err
	}
	machine.Brand.Name = brand.Name

	model, err := muc.modelRepo.GetByID(ctx, machine.Brand.Model.ID)
	if err != nil {
		return nil, err
	}
	machine.Model.Name = model.Name


	fuel, err := muc.fuelRepo.GetByID(ctx, machine.Fuel.ID)
	if err != nil {
		return nil, err
	}
	machine.Fuel.Name = fuel.Name
	
	driveUnit, err := muc.driveUnit.GetByID(ctx, machine.DriveUnit.ID)
	if err != nil {
		return nil, err
	}
	machine.DriveUnit.Name = driveUnit.Name

	bodyType, err := muc.bodyRepo.GetByID(ctx, machine.BodyType.ID)
	if err != nil {
		return nil, err
	}
	machine.BodyType.Name = bodyType.Name

	color, err := muc.colorRepo.GetByID(ctx, machine.Color.ID)
	if err != nil {
		return nil, err
	}
	machine.Color.Name = color.Name

	return machine, nil
}

func (muc MachineUseCase) GetRelevantMachines() ( []*models.Machine, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	list ,err :=  muc.machineRepo.GetList(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (muc MachineUseCase) GetListMachineByUserID(id int) ( []*models.Machine, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	list ,err :=  muc.machineRepo.GetListByUserID(ctx, id)
	if err != nil {
		return nil, err
	}
	return list, nil
}