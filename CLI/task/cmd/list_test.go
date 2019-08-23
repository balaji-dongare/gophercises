package cmd

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func TestListCmd(t *testing.T) {
	file, _ := os.Create("./testresult.txt")
	defer file.Close()
	defer os.Remove(file.Name())
	old := os.Stdout
	os.Stdout = file
	initdb()
	args := []string{}
	ListTask.Run(ListTask, args)

	expected := `1. do testing 2. do development 3. do deployment 4. do release `
	fp, _ := ioutil.ReadFile(file.Name())
	listoftask := string(fp)
	if expected != listoftask {
		expected = `No Tasks are in todos...Great!`
		fp, _ := ioutil.ReadFile(file.Name())
		listoftask := string(fp)
		if expected != listoftask {
			t.Error("error in list command")
		}
	}
	os.Stdout = old
}
func TestListCmdError(t *testing.T) {
	initdb()
	testdef := listTask
	defer func() {
		listTask = testdef
	}()

	listTask = func() ([]string, error) {
		return nil, errors.New("Got Error in ListTask")
	}
	ListTask.Run(&cobra.Command{}, []string{"Eoror got", "in list task"})
}
