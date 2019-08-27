package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/spf13/cobra"
)

func TestDoCmd(t *testing.T) {
	dotestSuit := []struct {
		args     []string
		expected string
	}{
		{args: []string{"1", "2", "3", "4"}, expected: `1. do testing2. do development3. do deployment4. do release`},
		{args: []string{"1", "2", "jhon", "3"}, expected: `Failed to parse Arg: jhon`},
		{args: []string{}, expected: `Please provide task`},
	}
	file, _ := os.Create("./testresult.txt")
	defer file.Close()
	defer os.Remove(file.Name())
	old := os.Stdout
	os.Stdout = file
	initdb()
	for _, testcase := range dotestSuit {
		DoTask.Run(DoTask, testcase.args)
		file.Seek(0, 0)
		fp, _ := ioutil.ReadFile(file.Name())

		match, err := regexp.Match(testcase.expected, fp)
		if err != nil {
			t.Error("Error in expected result regex")
		}
		if match {
			t.Log("Result is as Expected")
		} else {
			t.Error("Result is not as Expected")
		}
	}
	os.Stdout = old
}

//TestDoCmdError test Error Condition
func TestDoCmdError(t *testing.T) {
	initdb()
	testdef := doTask
	defer func() {
		doTask = testdef
	}()

	doTask = func(ids []int) ([]int, []string, error) {
		return nil, nil, errors.New("Got Error in DOTASK")
	}
	DoTask.Run(&cobra.Command{}, []string{"Eoror got", "in do task"})
}
