package container

import (
	"sync"

	"github.com/d34dl0ck/coupler/internal/core"
	"github.com/pkg/errors"
)

const (
	defaultStartCapacity = 16
)

type Registrations map[core.DependencyKey]core.ResolvingStrategy

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

func (c *MapContainer) Register(k core.DependencyKey, s core.ResolvingStrategy) {
	result := c.checkExistingRegistrations(k, s)

	c.mu.Lock()
	defer c.mu.Unlock()
	c.registrations[k] = result
}

func (c *MapContainer) Resolve(k core.DependencyKey) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	s, hasValue := c.registrations[k]

	if !hasValue {
		return nil, errors.Wrapf(core.ErrDependencyNotRegistered, "failed to find key %s", k)
	}

	return s.Resolve(c)
}

func (c *MapContainer) Check() error {
	checkResolver := newCheckResolver(c.registrations)
	c.mu.RLock()
	defer c.mu.RUnlock()

	for _, strategy := range c.registrations {
		_, err := strategy.Resolve(checkResolver)
		if errors.Is(err, core.ErrDependencyNotRegistered) {
			return err
		}
	}

	return nil
}

func (c *MapContainer) checkExistingRegistrations(k core.DependencyKey, s core.ResolvingStrategy) core.ResolvingStrategy {
	c.mu.RLock()
	defer c.mu.RUnlock()

	result := s

	_, hasValue := c.registrations[k]
	if hasValue {
		result = c.strategy.Solve(k, s, &c.registrations)
	}

	return result
}
