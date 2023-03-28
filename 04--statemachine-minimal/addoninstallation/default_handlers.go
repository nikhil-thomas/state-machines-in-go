package addoninstallation

import (
	"github.com/nikhil-thomas/addoninstall-minimal/statemachine_types"
)

type DefaultHandlers struct {
	sName statemachine_types.StateName
}

func (d *DefaultHandlers) Name() statemachine_types.StateName {
	return d.sName
}

func (d *DefaultHandlers) GoToInstalling() (bool, error) {
	return false, nil
}

func (d *DefaultHandlers) GoToPending() (bool, error) {
	return false, nil
}

func (d *DefaultHandlers) GoToReady() (bool, error) {
	return false, nil
}

func (d *DefaultHandlers) RunPreHooks() error {
	return nil
}

func (d *DefaultHandlers) RunPostHooks() error {
	return nil
}
