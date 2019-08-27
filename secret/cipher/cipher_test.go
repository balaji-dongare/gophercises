package cipher

import "testing"

func TestCipher(t *testing.T) {
	testSuit := []struct {
		key       string
		plainText string
	}{
		{key: "key123", plainText: "Testcase1"},
		{key: "key123", plainText: "Testcase2"},
		{key: "key123", plainText: "Testcase3"},
		{key: "key123", plainText: "Testcase4"},
	}
	for i, test := range testSuit {
		hex, err := Encrypt(test.key, test.plainText)
		if err != nil {
			t.Error(err)
		}
		plainText, err := Decrypt(test.key, hex)
		if err != nil {
			t.Error(err)
		}
		if test.plainText != plainText {
			t.Errorf("not value match %d", i)
		}
	}
}
