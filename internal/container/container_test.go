package container

import (
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
	k := core.NewRawResolvingKey("some_key")
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
	k := core.NewRawResolvingKey("some_key")
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
	k := core.NewRawResolvingKey("some_key")

	actual, err := c.Resolve(k)
	require.Nil(t, actual, "actual should be nil")
	require.ErrorIs(t, err, core.ErrKeyNotRegistered)
}
