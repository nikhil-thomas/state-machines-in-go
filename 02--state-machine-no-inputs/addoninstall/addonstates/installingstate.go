package addonstates

type InstallingHandler struct {
	defaultHandler
}

func (p *InstallingHandler) Name() string {
	return "state-installing"
}

func (i *InstallingHandler) SetStateSucceeded() error {
	return nil
}

func (i *InstallingHandler) SetStateFailed() error {
	return nil
}
