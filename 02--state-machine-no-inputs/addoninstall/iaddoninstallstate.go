package addoninstall

import "github.com/nikhil-thomas/statemachine/state"

type IAddonState interface {
	state.State
	SetStatePending() error
	SetStateInstalling() error
	SetStateFailed() error
	SetStateSucceeded() error
}
