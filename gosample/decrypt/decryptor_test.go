package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func Test_decryptor_Decrypt(t *testing.T) {
	// variables
	testKey := []byte("12345678901234567890123456789012") // The key should be 32 bytes (256 bytes) (AES-256)
	testPlaintext := []byte("this is plain text")
	testEncryptedFilePath := "encrypted.json"
	// errors
	errIoutilReadFile := errors.New("ioutil.Readfile error occurred")
	errNewGCM := errors.New("NewGCM Error Occurred")
	// functions
	fileCreator := func(filepath string, perm os.FileMode) {
		block, _ := aes.NewCipher(testKey)
		gcm, _ := cipher.NewGCM(block)
		// nonce is an arbitrary number that can be used just once in a cryptographic communication.
		// gcm.NonceSize() equals 12
		nonce := make([]byte, gcm.NonceSize())
		io.ReadFull(rand.Reader, nonce)
		ciphertext := gcm.Seal(nonce, nonce, testPlaintext, nil)
		// write file
		err := ioutil.WriteFile(filepath, ciphertext, perm)
		if err != nil {
			fmt.Println("[TEST WARNING] failed to create data file.", filepath)
		}
	}
	fileDeletor := func(filepath string) {
		if err := os.Remove(filepath); err != nil {
			fmt.Println("[TEST WARNING] failed to delete data file.", filepath)
		}
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
			d:             decryptor{encryptedFilePath: "encrypted.json", key: testKey},
			wantPlaintext: testPlaintext,
			wantErr:       nil,
			testSetup:     func() { fileCreator(testEncryptedFilePath, 0666) },
			testTeardown:  func() { fileDeletor(testEncryptedFilePath) },
		},
		{
			name:          "decryptor.Execute 暗号ファイルが存在しないとき Errorとしてioutil.ReadFileErrorが返ってくること",
			d:             decryptor{encryptedFilePath: "encrypted.json", key: testKey},
			wantPlaintext: nil,
			wantErr:       errIoutilReadFile,
			testSetup:     func() { ioutilReadFile = func(string) ([]byte, error) { return nil, errIoutilReadFile } },
			testTeardown:  func() { ioutilReadFile = ioutil.ReadFile },
		},
		{
			name:          "decryptor.Execute 鍵長が31byteのとき Errorとしてcipher.NewCipherErrorが返ってくること",
			d:             decryptor{encryptedFilePath: "encrypted.json", key: []byte("1234567890123456789012345678901")},
			wantPlaintext: nil,
			wantErr:       errors.New("crypto/aes: invalid key size 31"),
			testSetup:     func() { fileCreator(testEncryptedFilePath, 0666) },
			testTeardown:  func() { fileDeletor(testEncryptedFilePath) },
		},
		{
			name:          "decryptor.Execute NewGCMでErrorが発生したとき ErrorとしてNewGCMErrorが返ってくること",
			d:             decryptor{encryptedFilePath: "encrypted.json", key: []byte("12345678901234567890123456789013")},
			wantPlaintext: nil,
			wantErr:       errNewGCM,
			testSetup: func() {
				fileCreator(testEncryptedFilePath, 0666)
				cipherNewGCM = func(cipher cipher.Block) (cipher.AEAD, error) { return nil, errNewGCM }
			},
			testTeardown: func() {
				fileDeletor(testEncryptedFilePath)
				cipherNewGCM = cipher.NewGCM
			},
		},
		{
			name:          "decryptor.Execute 共通鍵が暗号時と異なるとき Error:CipherMとしてMessageAuthenticationFailedが返ってくること",
			d:             decryptor{encryptedFilePath: "encrypted.json", key: []byte("12345678901234567890123456789013")},
			wantPlaintext: nil,
			wantErr:       errors.New("cipher: message authentication failed"),
			testSetup:     func() { fileCreator(testEncryptedFilePath, 0666) },
			testTeardown:  func() { fileDeletor(testEncryptedFilePath) },
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
		key       []byte
		encrypted string
	}
	tests := []struct {
		name string
		args args
		want *decryptor
	}{
		{
			name: "NewDecryptor インスタンスを生成 decryptorインスタンスが返ってくること",
			args: args{
				key:       []byte("12345678901234567890123456789012"),
				encrypted: "encrypted.json",
			},
			want: &decryptor{
				key:               []byte("12345678901234567890123456789012"),
				encryptedFilePath: "encrypted.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDecryptor(tt.args.key, tt.args.encrypted); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDecryptor() = %v, want %v", got, tt.want)
			}
		})
	}
}
