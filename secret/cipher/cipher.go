package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

var ioReadFull = io.ReadFull

// Encrypt it encrypt's the key
func Encrypt(key, plaintext string) (string, error) {
	block, _ := newCipher(key)

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	initnvector := ciphertext[:aes.BlockSize]
	_, err := ioReadFull(rand.Reader, initnvector)
	if err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, initnvector)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return fmt.Sprintf("%x", ciphertext), err
}

var hexDecode = hex.DecodeString
var newCher = newCipher

// Decrypt it decrypt's the key value
func Decrypt(key, cipherHex string) (string, error) {
	var ciphertext []byte
	block, err := newCher(key)
	if err != nil {
		return "", err
	}
	ciphertext, err = hexDecode(cipherHex)
	if err != nil {
		return "", err
	}
	if len(ciphertext) >= aes.BlockSize {
		initnvector := ciphertext[:aes.BlockSize]
		ciphertext = ciphertext[aes.BlockSize:]

		stream := cipher.NewCFBDecrypter(block, initnvector)
		stream.XORKeyStream(ciphertext, ciphertext)
	}

	return string(ciphertext), err
}
func newCipher(key string) (cipher.Block, error) {
	hexcode := md5.New()
	fmt.Fprint(hexcode, key)
	cipherKey := hexcode.Sum(nil)
	return aes.NewCipher(cipherKey)
}
