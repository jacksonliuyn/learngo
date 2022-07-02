package dao

import "testing"

func TestExampleClient(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"test1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ExampleClient()
		})
	}
}
