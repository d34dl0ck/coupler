package strategies_test

import (
	"reflect"
	"testing"

	"github.com/d34dl0ck/coupler/internal/core"
	"github.com/d34dl0ck/coupler/internal/core/testdata"
	"github.com/d34dl0ck/coupler/internal/strategies"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestResolveByInstance(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := strategies.TestStruct{
		SomeString: strategies.ExpectedTestString,
	}
	strategy, err := strategies.NewInstanceStrategy(expected)
	require.NoError(t, err, "error was not expected")
	resolverMock := testdata.NewMockResolver(ctrl)

	actual, err := strategy.Resolve(resolverMock)
	require.NoError(t, err, "error was not expected")
	require.Equal(t, expected, actual, "resolved instance mismatch")
}

func TestInstanceStrategyDefaultKey(t *testing.T) {
	t.Parallel()

	input := strategies.TestStruct{
		SomeString: strategies.ExpectedTestString,
	}
	expected := core.NewTypeDependencyKey(reflect.TypeOf(input))
	strategy, err := strategies.NewInstanceStrategy(input)
	require.NoError(t, err, "err was not expected")

	actual := strategy.ProvideDefaultKey()

	require.Equal(t, expected, actual, "key mismatch")
}

func TestErrorNilInstance(t *testing.T) {
	t.Parallel()

	_, err := strategies.NewInstanceStrategy(nil)
	require.ErrorIs(t, err, strategies.ErrNilInput, "error mismatch")
}
