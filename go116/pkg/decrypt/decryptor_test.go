package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"reflect"
	"testing"
)

func Test_decryptor_Decrypt(t *testing.T) {
	// variables
	testKey := []byte("12345678901234567890123456789012") // The key should be 32 bytes (256 bytes) (AES-256)
	errNewGCM := errors.New("NewGCM Error Occurred")
	// functions
	encyptor := func(key, plaintext []byte) []byte {
		block, _ := aes.NewCipher(key)
		gcm, _ := cipher.NewGCM(block)
		nonce := make([]byte, gcm.NonceSize())
		io.ReadFull(rand.Reader, nonce)
		ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
		return ciphertext
	}
	// tests
	tests := []struct {
		name          string
		d             decryptor
		wantPlaintext []byte
		wantErr       error
		testSetup     func()
		testTeardown  func()
	}{
		{
			name:          "decryptor.Execute 正常に復号したとき plaintext:復号した平文およびError:Nilが返ってくること",
			d:             decryptor{key: testKey, ciphertext: encyptor(testKey, []byte("this is plaintext"))},
			wantPlaintext: []byte("this is plaintext"),
			wantErr:       nil,
			testSetup:     func() {},
			testTeardown:  func() {},
		},
		{
			name:          "decryptor.Execute 鍵長が31byteのとき Errorとしてcipher.NewCipherErrorが返ってくること",
			d:             decryptor{key: []byte("1234567890123456789012345678901"), ciphertext: encyptor(testKey, []byte("this is plaintext"))},
			wantPlaintext: nil,
			wantErr:       errors.New("crypto/aes: invalid key size 31"),
			testSetup:     func() {},
			testTeardown:  func() {},
		},
		{
			name:          "decryptor.Execute NewGCMでErrorが発生したとき ErrorとしてNewGCMErrorが返ってくること",
			d:             decryptor{key: testKey, ciphertext: encyptor(testKey, []byte("this is plaintext"))},
			wantPlaintext: nil,
			wantErr:       errNewGCM,
			testSetup:     func() { cipherNewGCM = func(cipher cipher.Block) (cipher.AEAD, error) { return nil, errNewGCM } },
			testTeardown:  func() { cipherNewGCM = cipher.NewGCM },
		},
		{
			name:          "decryptor.Execute 共通鍵が暗号時と異なるとき Error:CipherMとしてMessageAuthenticationFailedが返ってくること",
			d:             decryptor{key: []byte("12345678901234567890123456789013"), ciphertext: encyptor(testKey, []byte("this is plaintext"))},
			wantPlaintext: nil,
			wantErr:       errors.New("cipher: message authentication failed"),
			testSetup:     func() {},
			testTeardown:  func() {},
		},
	}
	isSameError := func(err, want error) bool {
		var errString, wantString string
		if err != nil {
			errString = err.Error()
		}
		if want != nil {
			wantString = want.Error()
		}
		if errString == wantString {
			return true
		}
		return false
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.testSetup()
			got, err := tt.d.Execute()
			if !isSameError(err, tt.wantErr) {
				t.Errorf("decryptor.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantPlaintext) {
				t.Errorf("decryptor.Execute() = %v, want %v", got, tt.wantPlaintext)
			}
			tt.testTeardown()
		})
	}
}

func TestNewDecryptor(t *testing.T) {
	type args struct {
		key        []byte
		ciphertext []byte
	}
	tests := []struct {
		name string
		args args
		want *decryptor
	}{
		{
			name: "NewDecryptor インスタンスを生成 decryptorインスタンスが返ってくること",
			args: args{
				key:        []byte("12345678901234567890123456789012"),
				ciphertext: []byte("ciphertext"),
			},
			want: &decryptor{
				key:        []byte("12345678901234567890123456789012"),
				ciphertext: []byte("ciphertext"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDecryptor(tt.args.key, tt.args.ciphertext); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDecryptor() = %v, want %v", got, tt.want)
			}
		})
	}
}
