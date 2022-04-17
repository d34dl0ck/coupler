package coupler

import (
	"testing"

	"github.com/d34dl0ck/coupler/internal/container"
	"github.com/stretchr/testify/require"
)

const (
	expectedString = "expected"
	expectedInt    = 32
)

type testStruct struct {
	SomeString string
	SomeInt    int
}

func TestRegistrationByInstance(t *testing.T) {
	t.Parallel()

	err := Register(ByInstance(expectedString))

	require.NoError(t, err, "no err was expected for string registration")

	actual, err := Resolve[string]()

	require.NoError(t, err, "no err was expected for resolving")
	require.Equal(t, expectedString, actual)
}

func TestRegistrationByFunc(t *testing.T) {
	t.Parallel()

	expected := createExpected(t)
	err := Register(ByFunc(func(s string, i int) testStruct {
		return testStruct{
			SomeString: s,
			SomeInt:    i,
		}
	}))
	require.NoError(t, err, "no err was expected for struct registration")

	actual, err := Resolve[testStruct]()

	require.NoError(t, err, "no err was expected for resolving")
	require.Equal(t, expected, actual)
}

func TestRegistrationByType(t *testing.T) {
	t.Parallel()

	expected := createExpected(t)
	err := Register(ByType[testStruct]())
	require.NoError(t, err, "no err was expected for struct registration")

	actual, err := Resolve[testStruct]()

	require.NoError(t, err, "no err was expected for resolving")
	require.Equal(t, expected, actual)
}

func TestRegistrationByTypeWithName(t *testing.T) {
	t.Parallel()
	dependencyName := "some_mega_dependency"

	expected := createExpected(t)
	err := Register(
		ByType[testStruct](),
		WithName(dependencyName))
	require.NoError(t, err, "no err was expected for struct registration")

	actual, err := ResolveNamed[testStruct](dependencyName)

	require.NoError(t, err, "no err was expected for resolving")
	require.Equal(t, expected, actual)
}

func TestErrorWhenStrategyWasNotSet(t *testing.T) {
	t.Parallel()

	err := Register(byEmptyRegistration())

	require.ErrorIs(t, err, container.ErrStrategyIsEmpty)
}

func byEmptyRegistration() ResolveOption {
	return func(r *Registration) {}
}

func createExpected(t *testing.T) testStruct {
	t.Helper()

	err := Register(ByInstance(expectedString))
	require.NoError(t, err, "no err was expected for string registration")

	err = Register(ByInstance(expectedInt))
	require.NoError(t, err, "no err was expected for int registration")

	return testStruct{
		SomeString: expectedString,
		SomeInt:    expectedInt,
	}
}
