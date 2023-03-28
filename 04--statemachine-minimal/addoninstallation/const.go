package addoninstallation

import (
	"fmt"
	"github.com/nikhil-thomas/addoninstall-minimal/statemachine_types"
)

const (
	StateNamePending    statemachine_types.StateName = "pending-state"
	StateNameInstalling                              = "installing-state"
	StateNameReady                                   = "ready-state"
)

var (
	ErrTransitionNotAllowed = fmt.Errorf("Transition not allowed")
)
