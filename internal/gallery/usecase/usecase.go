package usecase

import (
	"github.com/devstackq/bazar/internal/gallery"
)

func UseCasesInit(r gallery.Repositories) gallery.UseCases {
	return gallery.UseCases{
		MachineUseCaseInterface: MachineUseCaseInit(
			r.MachineRepoInterface,
		),
	}
}
