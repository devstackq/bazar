package usecase

import (
	"context"

	"github.com/devstackq/bazar/internal/gallery"
)

type FileManagerUseCase struct {
	fmRepo gallery.FileManagerRepoInterface
}

func FileManagerUseCaseInit(r gallery.FileManagerRepoInterface) gallery.FileManagerUseCaseInterface {
	return FileManagerUseCase{fmRepo: r}
}

// func (uuc FileManagerUseCase) GetListSrc(id int) ([]string, error) {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
// 	return uuc.fmRepo.GetListSrc(ctx, id)
// }

func (uuc FileManagerUseCase) CreateSrc(listSrc []string, id int) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	return uuc.fmRepo.CreateSrc(ctx, listSrc, id)
}
