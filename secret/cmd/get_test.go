package cmd

import (
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

func TestGet(t *testing.T) {
	file, err := os.Create("testcase.txt")
	if err != nil {
		t.Error("error in creating file")
	}
	defer file.Close()
	old := os.Stdout
	os.Stdout = file
	testSuit := []struct {
		encodingKey string
		key         string
		expected    string
	}{
		{encodingKey: "key123", key: "testcase1", expected: "This is testcase1"},
		{encodingKey: "key123", key: "tescase2", expected: "Key not found"},
	}
	for _, test := range testSuit {
		encodingKey = test.encodingKey
		args := []string{
			test.key,
		}
		getCmd.Run(getCmd, args)
		file.Seek(0, 0)
		b, err := ioutil.ReadAll(file)
		if err != nil {
			t.Error(err)
		}
		match, err := regexp.Match(test.expected, b)
		if err != nil {
			t.Error(err)
		}
		if !match {
			t.Error("Not match")
		}

	}
	os.Stdout = old
}
