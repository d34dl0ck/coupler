package strategies

import (
	"errors"
	"reflect"
	"testing"

	"github.com/d34dl0ck/coupler/internal/core"
	"github.com/d34dl0ck/coupler/internal/core/testdata"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

const (
	ExpectedTestString = "some_expected"
)

var (
	ErrExpected = errors.New("some_expected_error")
)

type TestStruct struct {
	SomeString string
}

type TestStructUnexportedField struct {
	someString string
}

func TestResolveByFields(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := TestStruct{
		SomeString: ExpectedTestString,
	}
	strategy, err := NewFieldStrategy(reflect.TypeOf(expected))
	require.NoError(t, err, "error was not expected")
	resolverMock := testdata.NewMockResolver(ctrl)
	dependencyKey := core.NewTypeDependencyKey(reflect.TypeOf(ExpectedTestString))
	resolverMock.EXPECT().Resolve(dependencyKey).Return(ExpectedTestString, nil)

	actual, err := strategy.Resolve(resolverMock)
	require.NoError(t, err, "error was not expected")
	require.Equal(t, expected, actual, "resolved instance mismatch")
}

func TestReturnFieldDependencyResolveError(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := TestStruct{
		SomeString: ExpectedTestString,
	}
	strategy, err := NewFieldStrategy(reflect.TypeOf(expected))
	require.NoError(t, err, "error was not expected")
	resolverMock := testdata.NewMockResolver(ctrl)
	dependencyKey := core.NewTypeDependencyKey(reflect.TypeOf(ExpectedTestString))
	resolverMock.EXPECT().Resolve(dependencyKey).Return(nil, ErrExpected)

	_, err = strategy.Resolve(resolverMock)

	require.ErrorIs(t, err, ErrExpected, "error was expected")
}

func TestFieldStrategyDefaultKey(t *testing.T) {
	t.Parallel()

	input := TestStruct{
		SomeString: ExpectedTestString,
	}
	expected := core.NewTypeDependencyKey(reflect.TypeOf(input))
	strategy, err := NewFieldStrategy(reflect.TypeOf(input))
	require.NoError(t, err, "err was not expected")

	actual := strategy.ProvideDefaultKey()

	require.Equal(t, expected, actual, "key mismatch")
}

func TestErrorNilType(t *testing.T) {
	t.Parallel()

	_, err := NewFieldStrategy(nil)
	require.ErrorIs(t, err, ErrNilInput, "error mismatch")
}

func TestErrorFieldNilDependency(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := TestStruct{
		SomeString: ExpectedTestString,
	}
	strategy, err := NewFieldStrategy(reflect.TypeOf(expected))
	require.NoError(t, err, "error was not expected")
	resolverMock := testdata.NewMockResolver(ctrl)
	dependencyKey := core.NewTypeDependencyKey(reflect.TypeOf(ExpectedTestString))
	resolverMock.EXPECT().Resolve(dependencyKey).Return(nil, nil)

	_, err = strategy.Resolve(resolverMock)
	require.ErrorIs(t, err, ErrNilDependency, "error mismatch")
}

func TestUnexportedFieldDetectedOnConstruction(t *testing.T) {
	t.Parallel()
	var def TestStructUnexportedField
	brokenType := reflect.TypeOf(def)

	_, err := NewFieldStrategy(brokenType)

	require.ErrorIs(t, err, ErrUnexportedFieldDetected, "error mismatch")
}

func TestUnexportedFieldDetectedOnResolve(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var brokenDef TestStructUnexportedField
	var def TestStruct
	brokenType := reflect.TypeOf(brokenDef)
	normalType := reflect.TypeOf(def)
	s, err := NewFieldStrategy(normalType)
	fieldStrategy := s.(FieldStrategy)
	fieldStrategy.t = brokenType
	resolverMock := testdata.NewMockResolver(ctrl)

	_, err = fieldStrategy.Resolve(resolverMock)

	require.ErrorIs(t, err, ErrUnexportedFieldDetected, "error mismatch")
}
