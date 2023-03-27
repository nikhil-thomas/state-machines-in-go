package main

type State1 struct {
	defaultHandler
	sd StateMachineDriver
}

func newState1(sd StateMachineDriver) State {
	return &State1{
		sd: sd,
	}
}

func (d *State1) name() string {
	return "state-1"
}

func (s *State1) handleInput0() error {
	nextState := "state-0"
	sharedContext := s.sd.getSharedContext()
	sharedContext++
	if sharedContext > 10 {
		sharedContext = 0
		nextState = "state-0"
	}
	s.sd.updateSharedContext(sharedContext)
	s.sd.setCurrentState(nextState)
	return nil
}

func (s *State1) handleInput1() error {
	nextState := "state-3"
	sharedContext := s.sd.getSharedContext()
	sharedContext++
	if sharedContext > 10 {
		sharedContext = 0
		nextState = "state-0"
	}

	s.sd.updateSharedContext(sharedContext)
	s.sd.setCurrentState(nextState)
	return nil
}
