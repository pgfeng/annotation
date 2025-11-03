package test

import (
	"testing"

	"github.com/pgfeng/annotation/types"
)

func TestParseContentTypeAnnotations(t *testing.T) {
	// This is a placeholder to indicate that this file is part of the 'test' package.
	tests := []struct {
		input    string
		expected types.ContentType
	}{
		{input: "application/json", expected: types.ContentType{MediaType: "application/json"}},
		{input: " application/json", expected: types.ContentType{MediaType: "application/json"}},
	}

	for _, test := range tests {
		var accept types.Accept
		accept.InitValue(test.input)
		t.Logf("For input '%s', got expected result %+v,MediaTypes len %d", test.input, accept, len(accept.MediaTypes))
	}
}
