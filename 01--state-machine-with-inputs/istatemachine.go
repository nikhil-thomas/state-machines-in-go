package main

type StateMachine interface {
	handleInput(int) error
}
