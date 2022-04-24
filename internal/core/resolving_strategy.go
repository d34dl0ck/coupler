package core

//go:generate mockgen --source=./resolving_strategy.go --destination=./testdata/resolving_strategy.go --package=testdata
type ResolvingStrategy interface {
	Resolve(r Resolver) (interface{}, error)
	ProvideDefaultKey() DependencyKey
}
