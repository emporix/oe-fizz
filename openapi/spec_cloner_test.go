package openapi

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

// TestClone tests that the Clone function creates a deep copy of the OpenAPI object.
// This includes the labels and the labels of the operations.
func TestClone(t *testing.T) {
	specData, err := os.ReadFile("../testdata/spec.json")
	require.NoError(t, err)

	var original OpenAPI
	err = json.Unmarshal(specData, &original)
	require.NoError(t, err)

	original.IncludeLabels = []string{"label1"}
	if original.Paths["/test/{a}"].GET != nil {
		original.Paths["/test/{a}"].GET.Labels = []string{"label1"}
	}

	clone, err := Clone(&original)
	require.NoError(t, err)

	assert.EqualValues(t, original, clone)
}
