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
	"flag"
	"fmt"
	"gosample/hoge/decrypt"
	"gosample/hoge/encrypt"
	"log"
)

func main() {
	dec := flag.Bool("dec", false, "decryption mode. default (false) is encryption mode")
	i := flag.String("i", "../plaintext.json", "plaintext file path")
	o := flag.String("o", "encrypted.json", "encrypted file path")
	flag.Parse()

	key := []byte("12345678901234567890123456789012") // The key should be 32 bytes (AES-256)
	if !*dec {
		e := encrypt.NewEncryptor(key, *i, *o)
		err := e.Execute()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		d := decrypt.NewDecryptor(key, *o)
		plaintext, err := d.Execute()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(plaintext))
	}
}
