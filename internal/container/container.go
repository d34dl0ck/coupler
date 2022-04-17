package container

import (
	"fmt"

	"github.com/pkg/errors"
)

const (
	defaultStartCapacity = 16
)

type Registrations map[ResolvingKey]ResolvingStrategy

type Container struct {
	registrations Registrations
	strategy      ConflictSolveStrategy
}

func (c *Container) Register(k ResolvingKey, s ResolvingStrategy) {
	result := s

	_, hasValue := c.registrations[k]
	if hasValue {
		result = c.strategy.Solve(k, s, &c.registrations)
	}

	c.registrations[k] = result
}

func (c *Container) Resolve(k ResolvingKey) (interface{}, error) {
	s, hasValue := c.registrations[k]

	if !hasValue {
		return nil, errors.Wrap(ErrKeyNotRegistered, fmt.Sprintf("failed to find key %s", k.Value()))
	}

	return s.Resolve(c.registrations)
}

func NewContainer() *Container {
	return &Container{
		registrations: make(Registrations, defaultStartCapacity),
		strategy:      overwriteStrategy{},
	}
}
