package container

type ResolvingStrategy interface {
	Resolve(r Registrations) (interface{}, error)
	ProvideDefaultKey() ResolvingKey
}
