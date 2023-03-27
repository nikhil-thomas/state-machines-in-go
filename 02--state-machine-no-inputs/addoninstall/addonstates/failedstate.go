package addonstates

type FailedHandler struct {
	defaultHandler
}

func (p *FailedHandler) Name() string {
	return "state-failed"
}

// no state transitions out of failed
