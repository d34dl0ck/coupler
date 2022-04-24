package coupler_test

import (
	"testing"

	c "github.com/d34dl0ck/coupler/pkg/coupler"
	"github.com/stretchr/testify/require"
)

type totallyUnexportedStruct struct {
	unexportedField string
}

type unexportedStructExportedField struct {
	ExportedField string
}

type unexportedStructMixedFields struct {
	ExportedField   string
	unexportedField string
}

type TotallyExportedStruct struct {
	ExportedField string
}

type ExportedStructMixedFields struct {
	ExportedField   string
	unexportedField string
}

type ExportedStructUnexportedField struct {
	unexportedField string
}

func TestTotallyUnexportedStruct(t *testing.T) {
	c.Register(c.ByInstance("some_string"))

	err := c.Register(c.ByType[totallyUnexportedStruct]())

	require.ErrorIs(t, err, c.ErrRegistration)
}

func TestUnexportedStructMixedFields(t *testing.T) {
	c.Register(c.ByInstance("some_string"))

	err := c.Register(c.ByType[unexportedStructMixedFields]())

	require.ErrorIs(t, err, c.ErrRegistration)
}

func TestExportedStructUnexportedField(t *testing.T) {
	c.Register(c.ByInstance("some_string"))

	err := c.Register(c.ByType[ExportedStructUnexportedField]())

	require.ErrorIs(t, err, c.ErrRegistration)
}

func TestExportedStructMixedFields(t *testing.T) {
	c.Register(c.ByInstance("some_string"))

	err := c.Register(c.ByType[ExportedStructMixedFields]())

	require.ErrorIs(t, err, c.ErrRegistration)
}

func TestUnexportedStructExportedField(t *testing.T) {
	c.Register(c.ByInstance("some_string"))

	err := c.Register(c.ByType[unexportedStructExportedField]())
	require.NoError(t, err, "error was not expected")

	instance, err := c.Resolve[unexportedStructExportedField]()
	require.NoError(t, err, "error was not expected")
	require.Equal(t, "some_string", instance.ExportedField)
}

func TestTotallyExportedStruct(t *testing.T) {
	c.Register(c.ByInstance("some_string"))

	err := c.Register(c.ByType[TotallyExportedStruct]())
	require.NoError(t, err, "error was not expected")

	instance, err := c.Resolve[TotallyExportedStruct]()
	require.NoError(t, err, "error was not expected")
	require.Equal(t, "some_string", instance.ExportedField)
}
