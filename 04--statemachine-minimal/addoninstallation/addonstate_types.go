package addoninstallation

import "github.com/nikhil-thomas/addoninstall-minimal/statemachine_types"

type AddonState interface {
	statemachine_types.State
	GoToInstalling() (bool, error)
	GoToPending() (bool, error)
	GoToReady() (bool, error)
	//goToUpgrading() error
	//goToDeleted() error
	//goToDeleting() error
	//goToFailed() error
	//gotToUndefined() error
}

//type AddonStateMachineDriver interface {
//	GetCurrentState() AddonState
//	SetCurrentState(statemachine_types.StateName) error
//	AddState(AddonState)
//	UpdateStateMachineConfig(AddonStateMachineControl)
//	GetStateMachineConfig() AddonStateMachineControl
//}
