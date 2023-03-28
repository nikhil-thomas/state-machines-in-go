package main

import (
	"fmt"
	"github.com/nikhil-thomas/addoninstall-minimal/addoninstallation"
	"github.com/nikhil-thomas/addoninstall-minimal/statemachine_types"
	"log"
)

type AddonInstallationModel struct {
	ClusterID      string
	AddonVersionID string
	Addon          string
	State          statemachine_types.StateName
}

func (a *AddonInstallationModel) TransitionToState(toStateName statemachine_types.StateName) error {
	d := addoninstallation.NewAddonInstalltionDriver(a.State)
	err := d.TransitionToState(toStateName)
	if err != nil {
		return err
	}
	a.State = toStateName
	return nil
}

func main() {
	addonInstInstance := &AddonInstallationModel{
		ClusterID:      "abcd",
		AddonVersionID: "1.2.3",
		Addon:          "reference-addon",
		State:          addoninstallation.StateNamePending,
	}

	fmt.Printf("Attempt transition %s to %s\n", addoninstallation.StateNamePending, addoninstallation.StateNameReady)
	fmt.Printf("before: current state: %s\n", addonInstInstance.State)
	err := addonInstInstance.TransitionToState(addoninstallation.StateNameReady)
	if err != nil && err != addoninstallation.ErrTransitionNotAllowed {
		log.Fatalln(err)
	}
	fmt.Printf("after: current state: %s\n", addonInstInstance.State)
	fmt.Println()
	// Output
	// Attempt transition pending-state to ready-state
	// before: current state: pending-state
	// after: current state: pending-state

	fmt.Printf("Attempt transition %s to %s\n", addoninstallation.StateNamePending, addoninstallation.StateNameInstalling)
	fmt.Printf("before: current state: %s\n", addonInstInstance.State)
	err = addonInstInstance.TransitionToState(addoninstallation.StateNameInstalling)
	if err != nil && err != addoninstallation.ErrTransitionNotAllowed {
		log.Fatalln(err)
	}
	fmt.Printf("after: current state: %s\n", addonInstInstance.State)
	fmt.Println()
	// Output
	// Attempt transition pending-state to installing-state
	// before: current state: pending-state
	// after: current state: installing-state

	fmt.Printf("Attempt transition %s to %s\n", addoninstallation.StateNameInstalling, addoninstallation.StateNameReady)
	fmt.Printf("before: current state: %s\n", addonInstInstance.State)
	err = addonInstInstance.TransitionToState(addoninstallation.StateNameReady)
	if err != nil && err != addoninstallation.ErrTransitionNotAllowed {
		log.Fatalln(err)
	}
	fmt.Printf("after: current state: %s\n", addonInstInstance.State)
	fmt.Println()
	// Output
	// Attempt transition installing-state to ready-state
	// before: current state: installing-state
	// after: current state: ready-state

}
