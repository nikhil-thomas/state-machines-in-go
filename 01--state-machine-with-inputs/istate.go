package main

type State interface {
	name() string
	handleInput0() error
	handleInput1() error
}
