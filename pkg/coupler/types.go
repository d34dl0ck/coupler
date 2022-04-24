package coupler

import "github.com/d34dl0ck/coupler/internal/core"

type Strategy core.ResolvingStrategy

type Key core.DependencyKey

type RegistrationOption func(r *Registration) error

type ResolveOption func(r *Registration) error

type Registration struct {
	Key      Key
	Strategy Strategy
}
