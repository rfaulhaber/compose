package composefile

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"strings"
	"testing"
)

func TestFunctionType_String(t *testing.T) {
	testCases := []struct{
		input FunctionType
		expected string
	}{
		{
			Go,
			"go",
		},
		{
			JS,
			"js",
		},
		{
			Java,
			"java",
		},
		{
			CS,
			"cs",
		},
		{
			Py,
			"py",
		},
	}

	for _, tc := range testCases {
		output := tc.input.String()
		assert.Equal(t, tc.expected, output)
	}
}

func TestFunctionType_MarshalYAML(t *testing.T) {
	testCases := []struct{
		input FunctionType
		expected string
	}{
		{
			Go,
			"go",
		},
		{
			JS,
			"js",
		},
		{
			Java,
			"java",
		},
		{
			CS,
			"cs",
		},
		{
			Py,
			"py",
		},
	}

	for _, tc := range testCases {
		output, err := yaml.Marshal(tc.input)
		assert.NoError(t, err)
		assert.Equal(t, tc.expected, strings.TrimSpace(string(output)))
	}
}

func TestFunctionType_UnmarshalYAML(t *testing.T) {
	testCases := []struct{
		input string
		expected FunctionType
	}{
		{
			"go",
			Go,
		},
		{
			"js",
			JS,
		},
		{
			"java",
			Java,
		},
		{
			"cs",
			CS,
		},
		{
			"py",
			Py,
		},
	}

	for _, tc := range testCases {
		output, marshalErr := yaml.Marshal(tc.input)

		var out FunctionType

		unmarshalErr := yaml.Unmarshal(output, out)

		assert.NoError(t, marshalErr)
		assert.NoError(t, unmarshalErr)
		assert.Equal(t, tc.expected, out)
	}
}
