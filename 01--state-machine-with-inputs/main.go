package main

func main() {
	sm := newStateMachine()
	sm.handleInput(0)
	sm.handleInput(1)
	sm.handleInput(0)
	sm.handleInput(1)
	sm.handleInput(0)
}
