package container

//go:generate mockgen --source=./resolving_strategy.go --destination=./testdata/resolving_strategy.go --package=testdata
type ResolvingStrategy interface {
	Resolve(r Registrations) (interface{}, error)
	ProvideDefaultKey() ResolvingKey
}
