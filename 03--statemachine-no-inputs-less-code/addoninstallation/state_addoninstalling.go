package addoninstallation

import (
	"github.com/nikhil-thomas/addoninstall/statemachine_types"
)

type InstallingState struct {
	DefaultHandlers
	driver statemachine_types.StateMachineDriver
}

func (d *InstallingState) Name() string {
	return StateNameInstalling
}

func (d *InstallingState) GoToReady() error {
	return nil
}
