package main

type State2 struct {
	defaultHandler
	sd StateMachineDriver
}

func newState2(sd StateMachineDriver) State {
	return &State2{
		sd: sd,
	}
}

func (d *State2) name() string {
	return "state-2"
}

func (s *State2) handleInput0() error {
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

func (s *State2) handleInput1() error {
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
