package main

import "fmt"

type defaultHandler struct {
}

func (d *defaultHandler) handleInput0() error {
	fmt.Println("Not transitions")
	return nil
}

func (d *defaultHandler) handleInput1() error {
	fmt.Println("Not tranisitons")
	return nil
}

func (d *defaultHandler) name() string {
	return ""
}
