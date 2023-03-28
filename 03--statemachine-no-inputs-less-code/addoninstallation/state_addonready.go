package addoninstallation

import (
	"github.com/nikhil-thomas/addoninstall/statemachine_types"
)

type ReadyState struct {
	DefaultHandlers
	driver statemachine_types.StateMachineDriver
}

func (d *ReadyState) Name() string {
	return StateNameReady
}
