package addonstates

type PendingHandler struct {
	defaultHandler
}

func (p *PendingHandler) Name() string {
	return "state-pending"
}

func (p *PendingHandler) SetStateInstalling() error {
	return nil
}

func (p *PendingHandler) SetStateFailed() error {
	return nil
}
