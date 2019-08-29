package cipher

import (
	"crypto/cipher"
	"errors"
	"io"

	"testing"
)

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
func TestDecryptError(t *testing.T) {

	testdef := hexDecode
	defer func() {
		hexDecode = testdef
	}()

	hexDecode = func(s string) ([]byte, error) {
		return nil, errors.New("Got Error in Decrypt")
	}
	Decrypt("key123", "hex")
}

func TestEncryptError(t *testing.T) {

	testdef := ioReadFull
	defer func() {
		ioReadFull = testdef
	}()

	ioReadFull = func(r io.Reader, buf []byte) (n int, err error) {
		return 1, errors.New("Got Error in Encrypt")
	}
	Encrypt("key123", "hex")
}
func TestDecryptError1(t *testing.T) {

	testdef := newCher
	defer func() {
		newCher = testdef
	}()

	newCher = func(key string) (cipher.Block, error) {
		return nil, errors.New("Got Error in Decrypt")
	}
	Decrypt("key123", "hex")
}
