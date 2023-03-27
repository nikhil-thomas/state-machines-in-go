package main

import "fmt"

type StateMachineSimple struct {
	driver StateMachineDriver
}

func newStateMachine() *StateMachineSimple {
	driver := newStateMachineDriver()
	state0 := newState0(driver)
	state1 := newState1(driver)
	state2 := newState2(driver)
	state3 := newState2(driver)

	driver.addState(state0)
	driver.addState(state1)
	driver.addState(state2)
	driver.addState(state3)

	driver.setCurrentState(state0.name())

	s := &StateMachineSimple{
		driver: driver,
	}
	return s
}

func (s *StateMachineSimple) handleInput(val int) error {
	currentState := s.driver.getCurrentState()
	fmt.Printf("\ninput %d, currentState: %s\n", val, currentState.name())
	switch val {
	case 0:
		currentState.handleInput0()
	case 1:
		currentState.handleInput1()
	}
	return nil
}
