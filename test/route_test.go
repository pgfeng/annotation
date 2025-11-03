package test

import (
	"testing"

	"github.com/pgfeng/annotation/types"
)

func TestParseRoute(t *testing.T) {
	tests := []struct {
		input    string
		expected types.Route
	}{
		{input: "/api/v1/users GET", expected: types.Route{Path: "/api/v1/users", Method: "GET"}},
		{input: "/api/v1/users post", expected: types.Route{Path: "/api/v1/users", Method: "POST"}},
		{input: "/api/v1/users OPTIONS", expected: types.Route{Path: "/api/v1/users", Method: "OPTIONS"}},
		{input: "/api/{d+}/users OPTIONS", expected: types.Route{Path: "/api/v1/users", Method: "OPTIONS"}},
		{input: "/api/v1/users", expected: types.Route{Path: "/api/v1/users", Method: "GET"}},
		{input: "/api/v1/users   DELETE ", expected: types.Route{Path: "/api/v1/users", Method: "DELETE"}},
		{input: "", expected: types.Route{Path: "/", Method: "GET"}},
	}

	for _, test := range tests {
		var route types.Route
		route.InitValue(test.input)
		t.Logf("For input '%s', got expected result %+v", test.input, route)
	}
}
