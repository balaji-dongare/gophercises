package cmd

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestList(t *testing.T) {
	file, err := os.Create("testcase.txt")
	if err != nil {
		t.Error("error in creating file")
	}
	defer file.Close()
	old := os.Stdout
	os.Stdout = file
	args := []string{}
	listCmd.Run(listCmd, args)

	expected := `1. testcase1`
	fp, _ := ioutil.ReadFile(file.Name())
	listofkeys := string(fp)
	if expected != listofkeys {
		fp, _ := ioutil.ReadFile(file.Name())
		listofkeys := string(fp)
		if expected != listofkeys {
			t.Error("error in list command")
		}
	}
	os.Stdout = old
}
