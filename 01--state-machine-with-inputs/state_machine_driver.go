package main

import "fmt"

type StateMachineDirverSimple struct {
	currentState  State
	states        map[string]State
	sharedContext int
}

func newStateMachineDriver() StateMachineDriver {
	return &StateMachineDirverSimple{
		states: map[string]State{},
	}
}

func (smd *StateMachineDirverSimple) addState(s State) {
	smd.states[s.name()] = s
}

func (smd *StateMachineDirverSimple) getCurrentState() State {
	return smd.currentState
}

func (smd *StateMachineDirverSimple) setCurrentState(sKey string) {
	state := smd.states[sKey]
	smd.currentState = state
	fmt.Println("new current State:", smd.currentState.name())
}

func (smd *StateMachineDirverSimple) updateSharedContext(val int) {
	smd.sharedContext = val
}

func (smd *StateMachineDirverSimple) getSharedContext() int {
	return smd.sharedContext
}
