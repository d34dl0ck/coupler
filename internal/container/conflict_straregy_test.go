package container

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/d34dl0ck/coupler/internal/core"
	"github.com/d34dl0ck/coupler/internal/core/testdata"
)

func TestOverwriteStrategy(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := testdata.NewMockResolvingStrategy(ctrl)
	strategy := OverwriteStrategy{}
	registrations := make(Registrations, 0)

	actual := strategy.Solve(
		core.NewRawResolvingKey("some"),
		expected,
		&registrations,
	)

	require.Equal(t, expected, actual)
}
