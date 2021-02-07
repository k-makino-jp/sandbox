package decryptor

import (
	"crypto/cipher"
	"errors"
	"reflect"
	"testing"
)

func TestDecryptorImpl_Decrypt(t *testing.T) {
	type args struct {
		key []byte
	}
	tests := []struct {
		name         string
		e            *DecryptorImpl
		args         args
		wantErr      bool
		testSetup    func()
		testTeardown func()
	}{
		{
			name: "DecryptorImpl.Decrypt Decypted ReturnsEqualsNil",
			e: &DecryptorImpl{
				encryptedFilePath: "encrypted.json",
				decryptedFilePath: "decrypted.json",
			},
			args: args{
				key: []byte("12345678901234567890123456789012"),
			},
			wantErr:      false,
			testSetup:    func() {},
			testTeardown: func() {},
		},
		{
			name: "DecryptorImpl.Decrypt EncryptedFileIsNotFound ReturnsEqualsError",
			e: &DecryptorImpl{
				encryptedFilePath: "hoge.json",
				decryptedFilePath: "decrypted.json",
			},
			args: args{
				key: []byte("12345678901234567890123456789012"),
			},
			wantErr:      true,
			testSetup:    func() {},
			testTeardown: func() {},
		},
		{
			name: "DecryptorImpl.Decrypt InvalidKeySize31byte ReturnsEqualsError",
			e: &DecryptorImpl{
				encryptedFilePath: "encrypted.json",
				decryptedFilePath: "decrypted.json",
			},
			args: args{
				key: []byte("1234567890123456789012345678901"),
			},
			wantErr:      true,
			testSetup:    func() {},
			testTeardown: func() {},
		},
		{
			name: "DecryptorImpl.Decrypt NewGCMError ReturnsEqualsError",
			e: &DecryptorImpl{
				encryptedFilePath: "encrypted.json",
				decryptedFilePath: "decrypted.json",
			},
			args: args{
				key: []byte("12345678901234567890123456789012"),
			},
			wantErr: true,
			testSetup: func() {
				cipherNewGCM = func(cipher cipher.Block) (cipher.AEAD, error) {
					return nil, errors.New("NewGCM Error Occurred")
				}
			},
			testTeardown: func() {
				cipherNewGCM = cipher.NewGCM
			},
		},
		{
			name: "DecryptorImpl.Decrypt CipherMessageAuthenticationFailed ReturnsEqualsError",
			e: &DecryptorImpl{
				encryptedFilePath: "encrypted.json",
				decryptedFilePath: "decrypted.json",
			},
			args: args{
				key: []byte("12345678901234567890123456789013"),
			},
			wantErr:      true,
			testSetup:    func() {},
			testTeardown: func() {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.testSetup()
			if err := tt.e.Decrypt(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("DecryptorImpl.Decrypt() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.testTeardown()
		})
	}
}

func TestNewDecryptorImpl(t *testing.T) {
	type args struct {
		input  string
		output string
	}
	tests := []struct {
		name string
		args args
		want *DecryptorImpl
	}{
		{
			name: "NewDecryptorImpl CreateInstance ReturnsEqualsInstance",
			args: args{
				input:  "input.json",
				output: "output.json",
			},
			want: &DecryptorImpl{
				encryptedFilePath: "input.json",
				decryptedFilePath: "output.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDecryptorImpl(tt.args.input, tt.args.output); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDecryptorImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}
