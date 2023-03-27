package addonstatemachine

import (
	"github.com/nikhil-thomas/statemachine/addoninstall"
	"github.com/nikhil-thomas/statemachine/addoninstall/addonstates"
)

type AddonStateMachine struct {
	driver addoninstall.IAddonStateMachineDriver
	store  addoninstall.IAddonStateMachineStore
}

func NewAddonStateMachine() *AddonStateMachine {
	statePending := &addonstates.PendingHandler{}
	stateInstalling := &addonstates.InstallingHandler{}
	stateFailed := &addonstates.FailedHandler{}
	stateSucceeded := &addonstates.SucceededHandler{}

	driver := &AddonStateMachineDriver{
		states: map[string]addoninstall.IAddonState{
			statePending.Name():    statePending,
			stateInstalling.Name(): stateInstalling,
			stateFailed.Name():     stateFailed,
			stateSucceeded.Name():  stateSucceeded,
		},
	}
	driver.SetCurrentState(statePending)

	store := &AddonStateMachineStore{
		config: &addoninstall.AddonStateMachineConfig{},
	}

	asm := AddonStateMachine{
		driver: driver,
		store:  store,
	}
	return &asm
}

func (a *AddonStateMachine) TransitionTo(stateName string) error {
	return a.driver.AttemptTransition(stateName)
}
