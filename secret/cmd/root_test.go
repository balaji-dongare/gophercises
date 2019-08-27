package cmd

import "testing"

func TestRoot(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"secret set", true},
		{"secret get", true},
		{"secret list", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
