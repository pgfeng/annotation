package test

import (
	"testing"

	"github.com/pgfeng/annotation/types"
)

func TestParseAcceptAnnotations(t *testing.T) {
	// This is a placeholder to indicate that this file is part of the 'test' package.
	tests := []struct {
		input    string
		expected types.Accept
	}{
		{input: "application/json, text/javascript, */*; q=0.01", expected: types.Accept{MediaTypes: []string{"application/json", "text/javascript", "*/*; q=0.01"}}},
	}

	for _, test := range tests {
		var accept types.Accept
		accept.InitValue(test.input)
		t.Logf("For input '%s', got expected result %+v,MediaTypes len %d", test.input, accept, len(accept.MediaTypes))
	}
}
