package container

type ConflictSolveStrategy interface {
	Solve(k ResolvingKey, s ResolvingStrategy, r *Registrations) ResolvingStrategy
}

type overwriteStrategy struct{}

func (overwriteStrategy) Solve(k ResolvingKey, s ResolvingStrategy, r *Registrations) ResolvingStrategy {
	return s
}
