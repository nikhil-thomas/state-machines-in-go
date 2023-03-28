package addoninstallation

type InstallingState struct {
	DefaultHandlers
}

func (i *InstallingState) GoToReady() (bool, error) {
	return true, nil
}
