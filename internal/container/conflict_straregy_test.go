package container_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/d34dl0ck/coupler/internal/container"
	"github.com/d34dl0ck/coupler/internal/container/testdata"
)

func TestOverwriteStrategy(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := testdata.NewMockResolvingStrategy(ctrl)
	strategy := container.OverwriteStrategy{}
	registrations := make(container.Registrations, 0)

	actual := strategy.Solve(
		container.NewRawResolvingKey("some"),
		expected,
		&registrations,
	)

	require.Equal(t, expected, actual)
}
