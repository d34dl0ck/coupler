package core

import "errors"

var (
	ErrKeyNotRegistered        = errors.New("key was not found in registrations")
	ErrDependencyNotRegistered = errors.New("dependency was not registered")
	ErrStrategyIsEmpty         = errors.New("strategy was empty")
)
