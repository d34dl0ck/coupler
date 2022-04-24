package container

import (
	"github.com/d34dl0ck/coupler/internal/core"
	"github.com/pkg/errors"
)

type checkResolver struct {
	registrations Registrations
}

func newCheckResolver(r Registrations) core.Resolver {
	return &checkResolver{
		registrations: r,
	}
}

func (r *checkResolver) Resolve(k core.DependencyKey) (interface{}, error) {
	_, hasValue := r.registrations[k]

	if !hasValue {
		return nil, errors.Wrapf(core.ErrDependencyNotRegistered, "failed to find key %s", k)
	}

	return nil, nil
}
