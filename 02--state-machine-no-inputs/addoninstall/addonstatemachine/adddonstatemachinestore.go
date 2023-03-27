package addonstatemachine

import "github.com/nikhil-thomas/statemachine/addoninstall"

type AddonStateMachineStore struct {
	config *addoninstall.AddonStateMachineConfig
}

func (asms *AddonStateMachineStore) GetCurrentStateMachineConfig() *addoninstall.AddonStateMachineConfig {
	return asms.config
}

func (asms *AddonStateMachineStore) SetCurrentStateMachineConfig(config *addoninstall.AddonStateMachineConfig) {
	asms.config = config
}
