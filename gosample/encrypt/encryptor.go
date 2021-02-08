package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"log"
)

type Encryptor interface {
	Encrypt()
}

type EncryptorImpl struct {
	inputFilePath  string
	outputFilePath string
}

// cipher algorithm: AES-256 GCM Mode
func (e *EncryptorImpl) Encrypt() {
	plaintext, err := ioutil.ReadFile(e.inputFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// The key should be 32 bytes (256 bytes) (AES-256)
	key := []byte("12345678901234567890123456789012")
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Panic(err)
	}

	// nonce is an arbitrary number that can be used just once in a cryptographic communication.
	// gcm.NonceSize() equals 12
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	err = ioutil.WriteFile(e.outputFilePath, ciphertext, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func NewEncryptorImpl(input, output string) *EncryptorImpl {
	return &EncryptorImpl{
		inputFilePath:  input,
		outputFilePath: output,
	}
}
