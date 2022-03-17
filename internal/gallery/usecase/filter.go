package usecase

import "github.com/devstackq/bazar/internal/gallery"

type MachineFilterUseCase struct {
	filterRepo gallery.FilterRepoInterface
}