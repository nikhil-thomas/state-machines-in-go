package addoninstallation

import "fmt"

type AddonInstallationModel struct {
	ClusterID      string
	AddonVersionID string
	Addon          string
	State          AddonStateMachine
}

func main() {
	addonInstalltion := &AddonInstallationModel{
		ClusterID:      "abcd",
		AddonVersionID: "1.2.3",
		Addon:          "reference-addon",
		State:          AddonStateMachine{},
	}
	fmt.Println("current state:", addonInstalltion.State.GetCurrentState())

}
