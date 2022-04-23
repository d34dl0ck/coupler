package coupler

import (
	"errors"
	"testing"

	"github.com/d34dl0ck/coupler/internal/core"
	"github.com/stretchr/testify/require"
)

const (
	expectedString = "expected"
	expectedInt    = 32
)

var (
	ErrExpected = errors.New("some expected error")
)

type testStruct struct {
	SomeString string
	SomeInt    int
}

func (testStruct) TestMethod(t *testing.T) {
	t.Helper()
}

type testInterface interface {
	TestMethod(t *testing.T)
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

func TestRegistrationByTypeAsImplementation(t *testing.T) {
	t.Parallel()

	expected := createExpected(t)
	err := Register(
		ByType[testStruct](),
		AsImplementationOf[testInterface]())
	require.NoError(t, err, "no err was expected for struct registration")

	actual, err := Resolve[testInterface]()

	require.NoError(t, err, "no err was expected for resolving")
	require.Equal(t, expected, actual)
}

func TestErrorWhenStrategyWasNotSet(t *testing.T) {
	t.Parallel()

	err := Register(byEmptyResolve())

	require.ErrorIs(t, err, core.ErrStrategyIsEmpty)
}

func TestErrorWhenResolveOptionReturnError(t *testing.T) {
	t.Parallel()

	err := Register(byErrorResolve())
	require.ErrorIs(t, ErrExpected, err)
}

func TestErrorWhenRegistrationOptionReturnError(t *testing.T) {
	t.Parallel()

	err := Register(byEmptyResolve(), withErrorRegistration())
	require.ErrorIs(t, ErrExpected, err)
}

func byEmptyResolve() ResolveOption {
	return func(r *Registration) error {
		return nil
	}
}

func byErrorResolve() ResolveOption {
	return func(r *Registration) error {
		return ErrExpected
	}
}

func withErrorRegistration() RegistrationOption {
	return func(r *Registration) error {
		return ErrExpected
	}
}

func createExpected(t *testing.T) testInterface {
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
