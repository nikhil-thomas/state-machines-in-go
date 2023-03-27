package main

type State0 struct {
	defaultHandler
	sd StateMachineDriver
}

func newState0(sd StateMachineDriver) State {
	return &State0{
		sd: sd,
	}
}

func (d *State0) name() string {
	return "state-0"
}

func (s *State0) handleInput0() error {
	nextState := "state-1"
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

func (s *State0) handleInput1() error {
	nextState := "state-2"
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
