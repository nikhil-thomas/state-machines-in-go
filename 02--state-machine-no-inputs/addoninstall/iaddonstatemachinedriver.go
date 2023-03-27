package addoninstall

type IAddonStateMachineDriver interface {
	SetCurrentState(IAddonState) error
	GetCurrentState() IAddonState
	AttemptTransition(string) error
}
