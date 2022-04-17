package container

import (
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

const (
	defaultStartCapacity = 16
)

type Registrations map[ResolvingKey]ResolvingStrategy

type Container struct {
	registrations Registrations
	strategy      ConflictSolveStrategy
	mu            *sync.RWMutex
}

func NewContainer() *Container {
	return &Container{
		registrations: make(Registrations, defaultStartCapacity),
		strategy:      OverwriteStrategy{},
		mu:            &sync.RWMutex{},
	}
}

func (c *Container) Register(k ResolvingKey, s ResolvingStrategy) {
	result := c.checkExistingRegistrations(k, s)

	c.mu.Lock()
	defer c.mu.Unlock()
	c.registrations[k] = result
}

func (c *Container) Resolve(k ResolvingKey) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	s, hasValue := c.registrations[k]

	if !hasValue {
		return nil, errors.Wrap(ErrKeyNotRegistered, fmt.Sprintf("failed to find key %s", k.Value()))
	}

	return s.Resolve(c.registrations)
}

func (c *Container) checkExistingRegistrations(k ResolvingKey, s ResolvingStrategy) ResolvingStrategy {
	c.mu.RLock()
	defer c.mu.RUnlock()

	result := s

	_, hasValue := c.registrations[k]
	if hasValue {
		result = c.strategy.Solve(k, s, &c.registrations)
	}

	return result
}
