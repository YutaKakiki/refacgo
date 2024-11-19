package utils

import (
	"strings"
	"testing"
)

func TestDecideToApply(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"AcceptInput", "y\n", true},
		{"RejectInput", "n\n", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モック入力
			input := strings.NewReader(tt.input)
			result := DecideToApply(input)

			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
