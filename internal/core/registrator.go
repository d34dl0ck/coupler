package core

//go:generate mockgen --source=./registrator.go --destination=./testdata/registrator.go --package=testdata
type Registrator interface {
	Register(k ResolvingKey, s ResolvingStrategy)
}
