package cmd

import "testing"

//TestRootCmd testcase for root task command
func TestRootCmd(t *testing.T) {
	err := execute()
	if err != nil {
		t.Error("Command not fould")
	}
}
