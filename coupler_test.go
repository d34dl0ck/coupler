package coupler

import (
	"errors"
	"testing"

	"github.com/d34dl0ck/coupler/internal/container"
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

type notRegisteredType struct{}

func (testStruct) TestMethod(t *testing.T) {
	t.Helper()
}

type testInterface interface {
	TestMethod(t *testing.T)
}

func TestRegistrationByInstance(t *testing.T) {
	err := Register(ByInstance(expectedString))

	require.NoError(t, err, "no err was expected for string registration")

	actual, err := Resolve[string]()

	require.NoError(t, err, "no err was expected for resolving")
	require.Equal(t, expectedString, actual)
}

func TestRegistrationByFunc(t *testing.T) {
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
	expected := createExpected(t)
	err := Register(ByType[testStruct]())
	require.NoError(t, err, "no err was expected for struct registration")

	actual, err := Resolve[testStruct]()

	require.NoError(t, err, "no err was expected for resolving")
	require.Equal(t, expected, actual)
}

func TestRegistrationByTypeWithName(t *testing.T) {
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
	err := Register(byEmptyResolve())

	require.ErrorIs(t, err, core.ErrStrategyIsEmpty)
}

func TestErrorWhenResolveOptionReturnError(t *testing.T) {
	err := Register(byErrorResolve())
	require.ErrorIs(t, ErrExpected, err)
}

func TestErrorWhenRegistrationOptionReturnError(t *testing.T) {
	err := Register(byEmptyResolve(), withErrorRegistration())
	require.ErrorIs(t, ErrExpected, err)
}

func TestErrorTypeInterfaceRegistration(t *testing.T) {
	err := Register(ByType[testInterface]())
	require.ErrorIs(t, err, ErrRegistration, "error mismatch")
}

func TestErrorNilInstanceRegistration(t *testing.T) {
	err := Register(ByInstance(nil))
	require.ErrorIs(t, err, ErrRegistration, "error mismatch")
}

func TestErrorNilFuncRegistration(t *testing.T) {
	err := Register(ByFunc(nil))
	require.ErrorIs(t, err, ErrRegistration, "error mismatch")
}

func TestCheckSuccess(t *testing.T) {
	c = container.NewContainer()

	err := Register(ByInstance(expectedString))
	require.NoError(t, err, "no err was expected for string registration")

	err = Register(ByInstance(expectedInt))
	require.NoError(t, err, "no err was expected for int registration")

	err = Register(ByType[testStruct]())
	require.NoError(t, err, "no err was expected for custom type registration")

	err = Check()

	require.NoError(t, err, "no error was expected")
}

func TestCheckFail(t *testing.T) {
	c = container.NewContainer()

	err := Register(ByInstance(expectedInt))
	require.NoError(t, err, "no err was expected for int registration")

	err = Register(ByType[testStruct]())
	require.NoError(t, err, "no err was expected for custom type registration")

	err = Check()

	require.ErrorIs(t, err, ErrDependenciesInconsistent, "error mismatch")
}

func TestNoDependencyError(t *testing.T) {
	_, err := Resolve[notRegisteredType]()

	require.ErrorIs(t, err, core.ErrDependencyNotRegistered)
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
