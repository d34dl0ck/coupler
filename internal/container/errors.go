package container

import "errors"

var (
	ErrKeyNotRegistered        = errors.New("key was not found in registrations")
	ErrResolveFailed           = errors.New("failed to resolve")
	ErrDependencyNotRegistered = errors.New("dependency was not registered")
	ErrStrategyIsEmpty         = errors.New("strategy was empty")
)
