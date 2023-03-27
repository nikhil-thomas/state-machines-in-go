package addonstatemachine

import (
	"fmt"
	"github.com/nikhil-thomas/statemachine/addoninstall"
)

type AddonStateMachineDriver struct {
	states       map[string]addoninstall.IAddonState
	currentState addoninstall.IAddonState
}

func (a *AddonStateMachineDriver) SetCurrentState(st addoninstall.IAddonState) error {
	a.currentState = st
	return nil
}

func (a *AddonStateMachineDriver) GetCurrentState() addoninstall.IAddonState {
	return a.currentState
}

func (a *AddonStateMachineDriver) AttemptTransition(toStateName string) error {
	currentState := a.GetCurrentState()
	fmt.Println("current state:", currentState.Name(), "| attempting transition to:", toStateName)
	var err error
	switch toStateName {
	case "state-pending":
		err = currentState.SetStatePending()
	case "state-installing":
		err = currentState.SetStateInstalling()
	case "state-failed":
		err = currentState.SetStateFailed()
	case "state-succeeded":
		err = currentState.SetStateSucceeded()
	default:
		err = fmt.Errorf("unknow toStateName %s", toStateName)
	}
	if err == nil {
		nextState := a.states[toStateName]
		a.SetCurrentState(nextState)
		fmt.Println("transition successful: current state:", a.currentState.Name())
	} else {
		fmt.Println("transition \"un\"-successful: current state:", a.currentState.Name())
	}
	return err
}
