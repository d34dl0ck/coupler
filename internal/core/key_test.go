package core

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNotEmptyValue(t *testing.T) {
	t.Parallel()

	expected := "expected value"
	key := stringKey{
		value: expected,
	}

	actual := key.Value()

	require.Equal(t, expected, actual)
	require.False(t, key.IsEmpty())
}

func TestEmptyValue(t *testing.T) {
	t.Parallel()

	key := stringKey{}

	require.True(t, key.IsEmpty())
}

func TestNewRawKey(t *testing.T) {
	t.Parallel()

	expected := "expected value"
	key := NewRawDependencyKey(expected)

	actual := key.Value()

	require.Equal(t, expected, actual)
	require.False(t, key.IsEmpty())
}

func TestNewTypeKey(t *testing.T) {
	t.Parallel()

	var empty testStruct
	typeToResolve := reflect.TypeOf(empty)
	expected := typeToResolve.Name()
	key := NewTypeDependencyKey(typeToResolve)

	actual := key.Value()

	require.Equal(t, expected, actual)
	require.False(t, key.IsEmpty())
}

type testStruct struct{}
