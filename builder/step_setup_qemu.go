package builder

import (
	"context"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
)

// StepSetupQemu configures chroot environment to run binaries via qemu
type StepSetupQemu struct {
	ImageMountPointKey string
}

// Run the step
func (s *StepSetupQemu) Run(_ context.Context, state multistep.StateBag) multistep.StepAction {
	return multistep.ActionContinue
}

// Cleanup after step execution
func (s *StepSetupQemu) Cleanup(state multistep.StateBag) {
}
