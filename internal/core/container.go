package core

type Container interface {
	Resolver
	Register(k DependencyKey, s ResolvingStrategy)
	Check() error
}
