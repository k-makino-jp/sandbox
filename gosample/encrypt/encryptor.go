// package encrypt 暗号化向けパッケージ
package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
)

// 暗号化向けインターフェース
type Encryptor interface {
	Execute()
}

type encryptor struct {
	key               []byte
	plaintextFilePath string
	encryptedFilePath string
}

// Execute 暗号化処理関数
// cipher algorithm: AES-256 GCM Mode
func (e encryptor) Execute() error {
	plaintext, err := ioutil.ReadFile(e.plaintextFilePath)
	if err != nil {
		return err
	}

	// The key should be 32 bytes (256 bytes) (AES-256)
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// nonce is an arbitrary number that can be used just once in a cryptographic communication.
	// gcm.NonceSize() equals 12
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	err = ioutil.WriteFile(e.encryptedFilePath, ciphertext, 0666)
	if err != nil {
		return err
	}
	return nil
}

func NewEncryptor(key []byte, plaintextFilePath, encryptedFilePath string) *encryptor {
	return &encryptor{
		key:               key,
		plaintextFilePath: plaintextFilePath,
		encryptedFilePath: encryptedFilePath,
	}
}
