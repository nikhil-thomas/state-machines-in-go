package addoninstallation

type PendingState struct {
	DefaultHandlers
}

func (d *PendingState) GoToInstalling() (bool, error) {
	return true, nil
}
