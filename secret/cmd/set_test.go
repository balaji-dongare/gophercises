package cmd

import (
	"io/ioutil"
	"os"
	"regexp"
	"testing"
)

func TestSet(t *testing.T) {
	file, err := os.Create("./testcase.txt")
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
	defer os.Remove(file.Name())
	old := os.Stdout
	os.Stdout = file
	testSuit := []struct {
		encodingKey string
		key         string
		plainText   string
		expected    string
	}{
		{encodingKey: "key123", key: "testcase1", plainText: "This is testcase1", expected: "Key set"},
	}
	for _, test := range testSuit {
		encodingKey = test.encodingKey
		args := []string{
			test.key,
			test.plainText,
		}
		setCmd.Run(setCmd, args)
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
