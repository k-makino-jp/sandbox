// package main ファイル暗号/復号化向けプログラム
// Usage of main.go
//   -dec
//         decryption mode. default (false) is encryption mode
//   -i string
//         plaintext file path (default "../plaintext.json")
//   -o string
//         encrypted file path (default "encrypted.json")
// Example
//   Encryption
//     go run main.go -i plaintext.json -o encrypted.json
//   Decryption
//     go run main.go -i encrypted.json -dec
package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"gosample/pkg/decrypt"
	"gosample/pkg/encrypt"
	"gosample/pkg/file"
	"log"
)

var (
	key = []byte("12345678901234567890123456789012") // The key should be 32 bytes (AES-256)

	dec = flag.Bool("dec", false, "decryption mode. default (false) is encryption mode")
	i   = flag.String("i", "plaintext.json", "plaintext file path")
	o   = flag.String("o", "ciphertext.json", "ciphertext file path")
)

func generateKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	return key, err
}

func execEncrypt() error {
	plaintext, err := file.Read(*i)
	if err != nil {
		return err
	}

	e := encrypt.NewEncryptor(key, plaintext)
	ciphertext, err := e.Execute()
	if err != nil {
		return err
	}

	err = file.Write(*o, ciphertext, 0666)
	if err != nil {
		return err
	}
	return nil
}

func execDecrypt() error {
	ciphertext, err := file.Read(*o)
	if err != nil {
		return err
	}

	d := decrypt.NewDecryptor(key, ciphertext)
	plaintext, err := d.Execute()
	if err != nil {
		return err
	}

	fmt.Println(string(plaintext))
	return nil
}

func main() {
	flag.Parse()
	// key, err := generateKey()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var err error
	if !*dec {
		err = execEncrypt()
	} else {
		err = execDecrypt()
	}
	if err != nil {
		log.Fatal(err)
	}
}
