package openapi

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"reflect"
	"testing"
)

// TestClone tests that the Clone function creates a deep copy of the OpenAPI object.
func TestClone(t *testing.T) {
	specData, err := os.ReadFile("../testdata/spec.json")
	require.NoError(t, err)

	var original OpenAPI
	err = json.Unmarshal(specData, &original)
	require.NoError(t, err)

	clone, err := Clone(&original)
	require.NoError(t, err)

	originalJSON, err := json.Marshal(original)
	require.NoError(t, err)

	cloneJSON, err := json.Marshal(clone)
	require.NoError(t, err)

	assert.JSONEq(t, string(originalJSON), string(cloneJSON))
}

// TestYAMLMarshalingRefs tests that spec types
// that contains embedded references  are properly
// marshaled to YAML.
func TestYAMLMarshalingRefs(t *testing.T) {
	// tests for references.
	tests := []interface{}{
		&ParameterOrRef{Reference: &Reference{Ref: ""}},
		&SchemaOrRef{Reference: &Reference{Ref: ""}},
		&ResponseOrRef{Reference: &Reference{Ref: ""}},
		&HeaderOrRef{Reference: &Reference{Ref: ""}},
		&MediaTypeOrRef{Reference: &Reference{Ref: ""}},
		&ExampleOrRef{Reference: &Reference{Ref: ""}},
	}
	for _, i := range tests {
		values := reflect.ValueOf(i).MethodByName("MarshalYAML").Call(nil)

		if len(values) != 2 {
			t.Errorf("expected MarshalYAML to return 2 args, got %d", len(values))
		}
		ret := values[0]

		if !ret.CanInterface() {
			t.Error("cannot get interface for returned type")
		}
		if _, ok := ret.Interface().(*Reference); !ok {
			t.Error("returned type is not a reference")
		}
	}
}

// TestYAMLMarshalingTypes tests that spec types
// that contains embedded types  are properly
// marshaled to YAML.
func TestYAMLMarshalingTypes(t *testing.T) {
	// tests for types.
	tests := map[interface{}]interface{}{
		&ParameterOrRef{Parameter: new(Parameter)}: &Parameter{},
		&SchemaOrRef{Schema: new(Schema)}:          &Schema{},
		&ResponseOrRef{Response: new(Response)}:    &Response{},
		&HeaderOrRef{Header: new(Header)}:          &Header{},
		&MediaTypeOrRef{MediaType: new(MediaType)}: &MediaType{},
		&ExampleOrRef{Example: new(Example)}:       &Example{},
	}
	for i, e := range tests {
		values := reflect.ValueOf(i).MethodByName("MarshalYAML").Call(nil)

		if len(values) != 2 {
			t.Errorf("expected MarshalYAML to return 2 args, got %d", len(values))
		}
		ret := values[0]

		if ret.Type().Kind() != reflect.Interface {
			t.Error("cannot get underlying type of non interface")
		}
		uret := ret.Elem()

		if uret.Type() != reflect.TypeOf(e) {
			t.Errorf("expected type to be %s, got %s", reflect.TypeOf(e).String(), uret.Type().String())
		}
	}
}
