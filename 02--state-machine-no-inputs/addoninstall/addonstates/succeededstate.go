package addonstates

type SucceededHandler struct {
	defaultHandler
}

func (p *SucceededHandler) Name() string {
	return "state-succeeded"
}

// no state transitions out of succeeded
