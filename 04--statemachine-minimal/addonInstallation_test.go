package main

import (
	"fmt"
	"github.com/nikhil-thomas/addoninstall-minimal/addoninstallation"
	"github.com/nikhil-thomas/addoninstall-minimal/statemachine_types"
	"testing"
)

func TestTransitionToState(t *testing.T) {
	tt := []struct {
		FromStateName statemachine_types.StateName
		ToStateName   statemachine_types.StateName
		err           error
	}{
		{
			FromStateName: addoninstallation.StateNamePending,
			ToStateName:   addoninstallation.StateNamePending,
			err:           addoninstallation.ErrTransitionNotAllowed,
		},
		{
			FromStateName: addoninstallation.StateNamePending,
			ToStateName:   addoninstallation.StateNameInstalling,
			err:           nil,
		},
		{
			FromStateName: addoninstallation.StateNamePending,
			ToStateName:   addoninstallation.StateNameReady,
			err:           addoninstallation.ErrTransitionNotAllowed,
		},
		{
			FromStateName: addoninstallation.StateNameInstalling,
			ToStateName:   addoninstallation.StateNamePending,
			err:           addoninstallation.ErrTransitionNotAllowed,
		},
		{
			FromStateName: addoninstallation.StateNameInstalling,
			ToStateName:   addoninstallation.StateNameInstalling,
			err:           addoninstallation.ErrTransitionNotAllowed,
		},
		{
			FromStateName: addoninstallation.StateNameInstalling,
			ToStateName:   addoninstallation.StateNameReady,
			err:           nil,
		},
		{
			FromStateName: addoninstallation.StateNameReady,
			ToStateName:   addoninstallation.StateNameReady,
			err:           addoninstallation.ErrTransitionNotAllowed,
		},
		{
			FromStateName: addoninstallation.StateNameReady,
			ToStateName:   addoninstallation.StateNameInstalling,
			err:           addoninstallation.ErrTransitionNotAllowed,
		},
		{
			FromStateName: addoninstallation.StateNameReady,
			ToStateName:   addoninstallation.StateNameReady,
			err:           addoninstallation.ErrTransitionNotAllowed,
		},
	}
	for _, tc := range tt {
		testName := fmt.Sprintf("from %q to %q", tc.FromStateName, tc.ToStateName)
		t.Run(testName, func(t *testing.T) {
			addonInst := AddonInstallationModel{
				State: tc.FromStateName,
			}
			err := addonInst.TransitionToState(tc.ToStateName)
			if err != tc.err {
				t.Errorf("expected %q got %q", tc.err, err)
			}
		})
	}
}
