package decryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"io/ioutil"
	"log"
)

type Decryptor interface {
	Decrypt()
}

type DecryptorImpl struct {
	encryptedFilePath string
	decryptedFilePath string
}

func (e *DecryptorImpl) Decrypt() {
	ciphertext, err := ioutil.ReadFile(e.encryptedFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// The key should be 32 bytes (AES-256)
	key := []byte("12345678901234567890123456789012")
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}

	nonce := ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(e.decryptedFilePath, plaintext, 0755)
	if err != nil {
		log.Fatal(err)
	}

}

func NewDecryptorImpl(inputFilePath, outputFilePath string) *DecryptorImpl {
	return &DecryptorImpl{
		encryptedFilePath: inputFilePath,
		decryptedFilePath: outputFilePath,
	}
}
