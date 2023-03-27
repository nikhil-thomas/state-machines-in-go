package main

type StateMachineDriver interface {
	setCurrentState(string)
	getCurrentState() State
	addState(State)
	updateSharedContext(int)
	getSharedContext() int
}
