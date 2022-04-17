package container

type ConflictSolveStrategy interface {
	Solve(k ResolvingKey, s ResolvingStrategy, r *Registrations) ResolvingStrategy
}

type OverwriteStrategy struct{}

func (OverwriteStrategy) Solve(k ResolvingKey, s ResolvingStrategy, r *Registrations) ResolvingStrategy {
	return s
}
