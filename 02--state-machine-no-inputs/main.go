package main

import (
	"fmt"
	"github.com/nikhil-thomas/statemachine/addoninstall/addonstatemachine"
)

func main() {
	addonInstalStateMachine := addonstatemachine.NewAddonStateMachine()

	addonInstalStateMachine.TransitionTo("state-installing")

	addonInstalStateMachine.TransitionTo("state-succeeded")

	fmt.Printf("\n\nexample 2\n")

	addonInstalStateMachine2 := addonstatemachine.NewAddonStateMachine()
	addonInstalStateMachine2.TransitionTo("state-succeeded")
	addonInstalStateMachine2.TransitionTo("state-installing")
	addonInstalStateMachine2.TransitionTo("state-pending")
	addonInstalStateMachine2.TransitionTo("state-failed")
	addonInstalStateMachine2.TransitionTo("state-succeeded")

}
