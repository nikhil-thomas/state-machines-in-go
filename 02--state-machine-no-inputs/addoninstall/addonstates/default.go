package addonstates

import (
	"fmt"
)

type defaultHandler struct{}

func (d *defaultHandler) Name() string {
	return ""
}

func (d *defaultHandler) SetStatePending() error {
	return fmt.Errorf("transition to %s not allwoed", "state-pending")
}

func (d *defaultHandler) SetStateInstalling() error {
	return fmt.Errorf("transition to %s not allwoed", "state-installing")
}

func (d *defaultHandler) SetStateFailed() error {
	return fmt.Errorf("transition to %s not allwoed", "state-failed")
}

func (d *defaultHandler) SetStateSucceeded() error {
	return fmt.Errorf("transition to %s not allwoed", "state-succeeded")
}
