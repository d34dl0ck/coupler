package strategies

import "errors"

var (
	ErrNilDependency           = errors.New("dependency was resolved as nil")
	ErrNilInput                = errors.New("input is nil")
	ErrUnexportedFieldDetected = errors.New("unexported fields cannot be used for type resolving")
)
