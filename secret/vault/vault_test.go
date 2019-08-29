package vault

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func getSecretFilePath() string {
	dir, _ := os.Getwd()
	secretpath := filepath.Join(dir, ".secrets")
	return secretpath
}
func TestSet(t *testing.T) {
	testSuit := []struct {
		encodingKey string
		filepath    string
		key         string
		plainText   string
	}{
		{encodingKey: "key123", filepath: getSecretFilePath(), key: "testcase", plainText: "This is testcase"},
	}
	for _, test := range testSuit {
		v := GetVault(test.encodingKey, test.filepath)
		err := v.Set(test.key, test.plainText)
		if err != nil {
			t.Error("error in Set")
		}
	}
}

func TestGet(t *testing.T) {
	testSuit := []struct {
		encodingKey string
		filepath    string
		key         string
		plainText   string
	}{
		{encodingKey: "key123", filepath: getSecretFilePath(), key: "testcase", plainText: "This is testcase"},
		{encodingKey: "key123", filepath: getSecretFilePath() + "extra", key: "testcase", plainText: ""},
		{encodingKey: "key123", filepath: getSecretFilePath(), key: "testcase1", plainText: ""},
	}
	for _, test := range testSuit {
		v := GetVault(test.encodingKey, test.filepath)
		plainText, _ := v.Get(test.key)
		if plainText != test.plainText {
			t.Error("error in Get")
		}
	}
}
func TestList(t *testing.T) {
	expected := "testcase"
	v := GetVault("key123", getSecretFilePath())
	value, err := v.List()
	if err != nil {
		fmt.Print("\nNo key found")
	}
	for _, key := range value {
		if expected != key {
			t.Error("empty secret file")
		}
	}
}
func getSecretFilePathWrong() string {
	dir, _ := os.Getwd()
	secretpath := filepath.Join(dir, "./secrets")
	return secretpath
}

func TestListError(t *testing.T) {
	expected := "testcase"
	v := GetVault("key123", getSecretFilePathWrong())
	value, err := v.List()
	if err != nil {
		fmt.Print("\nNo key found")
	}
	for _, key := range value {
		if expected != key {
			t.Error("empty secret file")
		}
	}
}

// func TestSaveError(t *testing.T) {
// 	deftef := cipherEnc
// 	defer func() {
// 		cipherEnc = deftef
// 	}()
// 	deftef = func(key, plaintext string) (string, error) {
// 		return "test1", errors.New("Error")
// 	}
// 	v := GetVault("key123", getSecretFilePath())
// 	err := v.Set("testcase2", "This is testcase2")
// 	if err != nil {
// 		t.Error(err)
// 	}

// }
