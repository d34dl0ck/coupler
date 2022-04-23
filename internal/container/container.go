package container

import (
	"fmt"
	"sync"

	"github.com/d34dl0ck/coupler/internal/core"
	"github.com/pkg/errors"
)

const (
	defaultStartCapacity = 16
)

type Registrations map[core.ResolvingKey]core.ResolvingStrategy

type MapContainer struct {
	registrations Registrations
	strategy      ConflictSolveStrategy
	mu            *sync.RWMutex
}

func NewContainer() *MapContainer {
	return &MapContainer{
		registrations: make(Registrations, defaultStartCapacity),
		strategy:      OverwriteStrategy{},
		mu:            &sync.RWMutex{},
	}
}

func (c *MapContainer) Register(k core.ResolvingKey, s core.ResolvingStrategy) {
	result := c.checkExistingRegistrations(k, s)

	c.mu.Lock()
	defer c.mu.Unlock()
	c.registrations[k] = result
}

func (c *MapContainer) Resolve(k core.ResolvingKey) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	s, hasValue := c.registrations[k]

	if !hasValue {
		return nil, errors.Wrap(core.ErrKeyNotRegistered, fmt.Sprintf("failed to find key %s", k.Value()))
	}

	return s.Resolve(c)
}

func (c *MapContainer) checkExistingRegistrations(k core.ResolvingKey, s core.ResolvingStrategy) core.ResolvingStrategy {
	c.mu.RLock()
	defer c.mu.RUnlock()

	result := s

	_, hasValue := c.registrations[k]
	if hasValue {
		result = c.strategy.Solve(k, s, &c.registrations)
	}

	return result
}
