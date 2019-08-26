package vault

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/balaji-dongare/gophercises/secret/cipher"
)

// Vault struct for store keys
type Vault struct {
	encodingKey string
	filepath    string
	mutex       sync.Mutex
	keyValues   map[string]string
}

// GetVault  get vault struct
func GetVault(encodingKey, filepath string) *Vault {
	return &Vault{
		encodingKey: encodingKey,
		filepath:    filepath,
	}
}

func (v *Vault) loadFile() error {
	f, err := os.Open(v.filepath)
	if err != nil {
		v.keyValues = make(map[string]string)
		return err
	}
	defer f.Close()
	return v.readData(f)
}
func (v *Vault) readData(f io.Reader) error {
	dec := json.NewDecoder(f)
	return dec.Decode(&v.keyValues)
}

func (v *Vault) save(key, value string) error {
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println(err)
	} else {
		defer f.Close()
		hex, err := cipher.Encrypt(v.encodingKey, value)
		if err != nil {
			fmt.Println(err)
		} else {
			v.keyValues[key] = hex
			err = v.writer(f)
		}
	}
	return err
}
func (v *Vault) writer(f io.Writer) error {
	enc := json.NewEncoder(f)
	return enc.Encode(v.keyValues)
}

// Set function it set's the key value in secret
func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	v.loadFile()
	return v.save(key, value)
}

// Get function it get's the value from secret
func (v *Vault) Get(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.loadFile()
	if err != nil {
		return "", errors.New("File not found")
	}
	hex, ok := v.keyValues[key]
	if ok != true {
		return "", errors.New("Key not found")
	}
	return cipher.Decrypt(v.encodingKey, hex)
}

//List all the Api keys
func (v *Vault) List() ([]string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()
	err := v.loadFile()
	if err != nil {
		fmt.Println("File not found")
	}
	hexcode := v.keyValues
	keys := make([]string, len(hexcode))
	i := 0
	for keyitem := range hexcode {
		keys[i] = keyitem // explicit array element assignment instead of append function
		i++
	}
	return keys, err
}
