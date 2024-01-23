package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	input1 := http.Header{
		"Authorization": []string{"ApiKey abc"},
	}
	input2 := http.Header{
		"Authorization": []string{"ApiKey "},
	}
	input3 := http.Header{
		"auth": []string{"ApiKey abc"},
	}
	input4 := http.Header{
		"Authorization": []string{"Bearer abc"},
	}

	tests := []struct {
		input  http.Header
		apiKey string
		err    error
	}{
		{input: input1, apiKey: "abc", err: nil},
		{input: input2, apiKey: "", err: errors.New("malformed authorization header")},
		{input: input3, apiKey: "", err: errors.New("no authorization header included")},
		{input: input4, apiKey: "", err: errors.New("malformed authorization header")},
	}
	for i, test := range tests {
		key, err := GetAPIKey(test.input)
		if key != test.apiKey && err != test.err {
			t.Fatalf("Test %d failed!. Expected '%s' and '%v' as outputs, got '%s' and '%v' instead.", i, test.apiKey, test.err, key, err)
		}
	}
}
