package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/balaji-dongare/gophercises/CLI/task/dbrepository"
	"github.com/spf13/cobra"
)

// initdb initalize db for test environment
func initdb() {
	dir, _ := os.Getwd()
	databasepath := filepath.Join(dir, "tasks.db")
	dbrepository.InitDatabase(databasepath)
}

// TestAddCMD testcase for test task add command
func TestAddCMD(t *testing.T) {
	initdb()
	testSuit := []struct {
		args     []string
		expected string
	}{
		{args: []string{"do", "testing"}, expected: `Task:"do testing" is Added in todo list`},
		{args: []string{"do", "development"}, expected: `Task:"do development" is Added in todo list`},
		{args: []string{"do", "deployment"}, expected: `Task:"do deployment" is Added in todo list`},
		{args: []string{"do", "release"}, expected: `Task:"do release" is Added in todo list`},
		{args: []string{}, expected: `Please provide task`},
	}
	file, _ := os.Create("./testresult.txt")
	file.Truncate(0)
	defer file.Close()
	defer os.Remove(file.Name())
	old := os.Stdout
	os.Stdout = file
	for _, testcase := range testSuit {
		AddTask.Run(AddTask, testcase.args)
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
func TestAddCMDError(t *testing.T) {
	initdb()
	testdef := addTask
	defer func() {
		addTask = testdef
	}()

	addTask = func(task string) (bool, error) {
		return false, errors.New("Got Error in add task")
	}
	AddTask.Run(&cobra.Command{}, []string{"Got", "Error"})
}
