// package encrypt 暗号化向けパッケージ
package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// Encryptor 暗号化向けインターフェース
type Encryptor interface {
	Execute()
}

type encryptor struct {
	key       []byte
	plaintext []byte
}

// Execute 暗号化処理関数
// 暗号化アルゴリズム: AES-256 GCM Mode
// Reference: https://golang.org/pkg/crypto/cipher/#NewGCM
func (e encryptor) Execute() ([]byte, error) {
	// 暗号文ブロック生成(鍵長は32bytes)
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return nil, err
	}

	// GCMモードでラップされた128bitの暗号文ブロック取得
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Nonce(初期化ベクトル)設定
	nonce := make([]byte, gcm.NonceSize()) // gcm.NonceSize() equals 12
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Nonceに暗号文を結合しciphertextとして利用
	ciphertext := gcm.Seal(nonce, nonce, e.plaintext, nil)
	return ciphertext, nil
}

// NewEncryptor コンストラクタ
func NewEncryptor(key, plaintext []byte) *encryptor {
	return &encryptor{
		key:       key,
		plaintext: plaintext,
	}
}
