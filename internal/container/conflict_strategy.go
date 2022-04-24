package container

import "github.com/d34dl0ck/coupler/internal/core"

type ConflictSolveStrategy interface {
	Solve(k core.DependencyKey, s core.ResolvingStrategy, r *Registrations) core.ResolvingStrategy
}

type OverwriteStrategy struct{}

func (OverwriteStrategy) Solve(k core.DependencyKey, s core.ResolvingStrategy, r *Registrations) core.ResolvingStrategy {
	return s
}
