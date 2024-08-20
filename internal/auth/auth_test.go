package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testCases := []struct {
		headerType string
		tokenType  string
		input      string
		expected   string
	}{
		{
			headerType: "Authorization",
			tokenType:  "ApiKey",
			input:      "sampleAPIKey",
			expected:   "sampleAPIKey",
		},
		{
			headerType: "Authentication",
			tokenType:  "ApiKey",
			input:      "sampleAPIKey",
			expected:   "sampleAPIKey",
		},
		{
			headerType: "Authorization",
			tokenType:  "ApiKeyy",
			input:      "sampleAPIKey",
			expected:   "sampleAPIKey",
		},
		{
			headerType: "Authentication",
			tokenType:  "ApiKeyy",
			input:      "sampleAPIKey",
			expected:   "sampleAPIKey",
		},
	}

	for _, v := range testCases {
		header := http.Header{}
		header.Set(v.headerType, fmt.Sprintf("%v %v", v.tokenType, v.input))

		apiKey, err := GetAPIKey(header)
		if err != nil {
			if v.headerType == "Authorization" && v.tokenType == "ApiKey" {
				t.Errorf("Unexpected error: %v", err)
			} else {
				continue
			}
		}
		if apiKey != v.expected {
			t.Errorf("Expected apiKey %v, got %v", v.expected, apiKey)
		}
	}
}
