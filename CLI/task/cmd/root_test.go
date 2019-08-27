package cmd

import (
	"testing"
)

//TestRootCmd testcase for root task command
func TestExecute(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"task list", true},
		{"task do", true},
		{"task add", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Execute(); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
