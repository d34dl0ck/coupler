package core

import "errors"

var (
	ErrDependencyNotRegistered = errors.New("dependency was not registered")
	ErrStrategyIsEmpty         = errors.New("strategy was empty")
)
