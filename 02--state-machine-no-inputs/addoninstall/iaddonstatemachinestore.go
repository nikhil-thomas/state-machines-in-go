package addoninstall

type IAddonStateMachineStore interface {
	GetCurrentStateMachineConfig() *AddonStateMachineConfig
	SetCurrentStateMachineConfig(*AddonStateMachineConfig)
}
