package statemachine_types

import "context"

type StateName string

type State interface {
	Name() StateName
	RunPreHooks() error
	RunPostHooks() error
}

type PreHookFn func(context.Context) error

type PostHookFn func(context.Context) error
