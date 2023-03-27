package main

type State3 struct {
	defaultHandler
	sd StateMachineDriver
}

func newState3(sd StateMachineDriver) State {
	return &State3{
		sd: sd,
	}
}

func (d *State3) name() string {
	return "state-3"
}

func (s *State3) handleInput0() error {
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

func (s *State3) handleInput1() error {
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
