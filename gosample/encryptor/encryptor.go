package encryptor

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
	outputFilePath string
}

func (e *EncryptorImpl) Encrypt() {
	plaintext, err := ioutil.ReadFile("plaintext.json")
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

func NewEncryptorImpl(outputFilePath string) *EncryptorImpl {
	return &EncryptorImpl{
		outputFilePath: outputFilePath,
	}
}
