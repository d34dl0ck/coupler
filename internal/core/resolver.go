package core

//go:generate mockgen --source=./resolver.go --destination=./testdata/resolver.go --package=testdata
type Resolver interface {
	Resolve(k DependencyKey) (interface{}, error)
}
