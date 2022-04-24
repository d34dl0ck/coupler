package container

import (
	"errors"
	"testing"

	"github.com/d34dl0ck/coupler/internal/core"
	"github.com/d34dl0ck/coupler/internal/core/testdata"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestResolveWithStrategy(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := "some_expected_obj"

	c := NewContainer()
	k := core.NewRawDependencyKey("some_key")
	strategyMock := testdata.NewMockResolvingStrategy(ctrl)
	strategyMock.EXPECT().Resolve(c).Return(expected, nil)
	c.Register(k, strategyMock)

	actual, err := c.Resolve(k)
	require.NoError(t, err, "err was not expected")
	require.Equal(t, expected, actual)
}

func TestConflictSolvedWithOverwrite(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	expected := "some_expected_obj"

	c := NewContainer()
	k := core.NewRawDependencyKey("some_key")
	strategyMock := testdata.NewMockResolvingStrategy(ctrl)
	strategyMock.EXPECT().Resolve(c).Return(expected, nil)
	c.Register(k, strategyMock)
	c.Register(k, strategyMock)

	actual, err := c.Resolve(k)
	require.NoError(t, err, "err was not expected")
	require.Equal(t, expected, actual)
}

func TestErrMissingKey(t *testing.T) {
	t.Parallel()

	c := NewContainer()
	k := core.NewRawDependencyKey("some_key")

	actual, err := c.Resolve(k)
	require.Nil(t, actual, "actual should be nil")
	require.ErrorIs(t, err, core.ErrDependencyNotRegistered)
}

func TestCheck(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	normalKey := core.NewRawDependencyKey("normal")
	anotherNormalKey := core.NewRawDependencyKey("another")
	errorKey := core.NewRawDependencyKey("error")
	strategyMock := testdata.NewMockResolvingStrategy(ctrl)
	gomock.InOrder(
		strategyMock.EXPECT().Resolve(gomock.Any()).Return(nil, nil),
		strategyMock.EXPECT().Resolve(gomock.Any()).Return(nil, nil),
		strategyMock.EXPECT().Resolve(gomock.Any()).Return(nil, errors.New("some resolving error, but not about missing dependency")),
	)
	c := NewContainer()
	c.Register(normalKey, strategyMock)
	c.Register(anotherNormalKey, strategyMock)
	c.Register(errorKey, strategyMock)

	err := c.Check()

	require.NoError(t, err, "error was not expected")
}

func TestDependencyErrorCheck(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	normalKey := core.NewRawDependencyKey("normal")
	anotherNormalKey := core.NewRawDependencyKey("another")
	errorKey := core.NewRawDependencyKey("error")
	strategyMock := testdata.NewMockResolvingStrategy(ctrl)
	gomock.InOrder(
		strategyMock.EXPECT().Resolve(gomock.Any()).Return(nil, nil),
		strategyMock.EXPECT().Resolve(gomock.Any()).Return(nil, nil),
		strategyMock.EXPECT().Resolve(gomock.Any()).Return(nil, core.ErrDependencyNotRegistered),
	)
	c := NewContainer()
	c.Register(normalKey, strategyMock)
	c.Register(anotherNormalKey, strategyMock)
	c.Register(errorKey, strategyMock)

	err := c.Check()

	require.ErrorIs(t, err, core.ErrDependencyNotRegistered, "error mismatch")
}
