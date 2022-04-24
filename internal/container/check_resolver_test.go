package container

import (
	"testing"

	"github.com/d34dl0ck/coupler/internal/core"
	"github.com/stretchr/testify/require"
)

func TestFindDependency(t *testing.T) {
	t.Parallel()

	registrations := make(Registrations, 1)
	resolver := newCheckResolver(registrations)
	k := core.NewRawDependencyKey("some")
	registrations[k] = nil

	value, err := resolver.Resolve(k)

	require.Nil(t, value, "returning value should be nil")
	require.NoError(t, err, "error was not expected")
}

func TestNotFindDependency(t *testing.T) {
	t.Parallel()

	registrations := make(Registrations, 0)
	resolver := newCheckResolver(registrations)
	k := core.NewRawDependencyKey("some")

	value, err := resolver.Resolve(k)

	require.Nil(t, value, "returning value should be nil")
	require.ErrorIs(t, err, core.ErrDependencyNotRegistered, "error mismatch")
}
