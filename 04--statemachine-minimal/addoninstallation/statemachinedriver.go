package addoninstallation

import (
	"github.com/nikhil-thomas/addoninstall-minimal/statemachine_types"
)

type Driver struct {
	currentState AddonState
}

func NewAddonInstalltionDriver(currentStateName statemachine_types.StateName) *Driver {
	state := stateFromName(currentStateName)
	driver := &Driver{
		currentState: state,
	}
	return driver
}

func stateFromName(name statemachine_types.StateName) AddonState {
	var state AddonState
	switch name {
	case StateNamePending:
		state = &PendingState{}
	case StateNameInstalling:
		state = &InstallingState{}
	case StateNameReady:
		state = &ReadyState{}
	}
	return state
}

func (d *Driver) isTransitionAllowed(nextState statemachine_types.StateName) (bool, error) {
	switch nextState {
	case StateNamePending:
		return d.currentState.GoToPending()
	case StateNameInstalling:
		return d.currentState.GoToInstalling()
	case StateNameReady:
		return d.currentState.GoToReady()
	}
	return false, nil
}

func (d *Driver) TransitionToState(toStateName statemachine_types.StateName) error {
	ok, err := d.isTransitionAllowed(toStateName)
	if err != nil {
		return err
	}
	if !ok {
		return ErrTransitionNotAllowed
	}

	nextState := stateFromName(toStateName)
	d.currentState.RunPostHooks()
	nextState.RunPreHooks()
	d.currentState = nextState
	return nil
}
