// package decrypt 復号化向けパッケージ
package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"io/ioutil"
)

var (
	ioutilReadFile = ioutil.ReadFile
	cipherNewGCM   = cipher.NewGCM
)

// Decryptor 復号化向けインターフェース
type Decryptor interface {
	Execute() error
}

type decryptor struct {
	key               []byte
	encryptedFilePath string
}

// Execute 復号処理関数
func (d *decryptor) Execute() ([]byte, error) {
	ciphertext, err := ioutilReadFile(d.encryptedFilePath)
	if err != nil {
		return nil, err
	}
	// 暗号文ブロック生成
	block, err := aes.NewCipher(d.key)
	if err != nil {
		return nil, err
	}
	// GCMモードでラップされた128bitの暗号文ブロック取得
	gcm, err := cipherNewGCM(block)
	if err != nil {
		return nil, err
	}
	// 復号化
	nonce := ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

// NewDecyptor コンストラクタ
func NewDecryptor(key []byte, encrypted string) *decryptor {
	return &decryptor{
		key:               key,
		encryptedFilePath: encrypted,
	}
}
