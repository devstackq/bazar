package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/admin"
	"github.com/devstackq/bazar/internal/models"
)

type RoleUseCase struct {
	roleRepo admin.RoleRepoInterface
}

func RoleUseCaseInit(r admin.RoleRepoInterface) admin.RoleUseCaseInterface {
	return RoleUseCase{roleRepo: r}
}

func (ruc RoleUseCase) CreateRole(r *models.Role) (int, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return ruc.roleRepo.Create(ctx, r)
}

func (ruc RoleUseCase) GetListRole() ([]*models.Role, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return ruc.roleRepo.GetList(ctx)
}
