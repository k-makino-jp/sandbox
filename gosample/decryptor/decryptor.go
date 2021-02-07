package decryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
)

type Decryptor interface {
	Decrypt(key []byte) error
}

type DecryptorImpl struct {
	encryptedFilePath string
	decryptedFilePath string
}

var (
	cipherNewGCM = cipher.NewGCM
)

func (e *DecryptorImpl) Decrypt(key []byte) error {
	ciphertext, err := ioutil.ReadFile(e.encryptedFilePath)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipherNewGCM(block)
	if err != nil {
		return err
	}

	nonce := ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	fmt.Println(string(plaintext))
	// err = ioutil.WriteFile(e.decryptedFilePath, plaintext, 0755)
	// if err != nil {
	// 	return err
	// }
	return nil

}

func NewDecryptorImpl(input, output string) *DecryptorImpl {
	return &DecryptorImpl{
		encryptedFilePath: input,
		decryptedFilePath: output,
	}
}
