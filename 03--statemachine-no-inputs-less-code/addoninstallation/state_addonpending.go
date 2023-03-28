package addoninstallation

import (
	"github.com/nikhil-thomas/addoninstall/statemachine_types"
)

type PendingState struct {
	DefaultHandlers
	driver statemachine_types.StateMachineDriver
}

func (d *PendingState) Name() string {
	return StateNamePending
}

func (d *PendingState) GoToInstalling() error {
	return nil
}
