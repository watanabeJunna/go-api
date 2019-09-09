package main

import (
	"crypto/aes"
	"crypto/cipher"
	// "encoding/base64"
	"log"
	"fmt"
	"io"
	"crypto/rand"
)

func main() {
	// plainText := []byte("secret text 9999")

	// key := "qyXvvfcK7swJcQCmaKlsv/eVLm5+netR"

	// encrypted := make([]byte, aes.BlockSize+len(plainText))
	// iv := encrypted[:aes.BlockSize]
	// iv, err := base64.StdEncoding.DecodeString("HRVXPTZPq8X8CS3Vt7K9ow==")

	// block, err := aes.NewCipher([]byte(key))

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// mode := cipher.NewCBCEncrypter(block, iv)
	// mode.CryptBlocks(encrypted[aes.BlockSize:], plainText)
	// fmt.Printf("encrypted: %x\n", encrypted)

	plainText := []byte("secret text 9999")
	
	// 暗号化データ。先頭に初期化ベクトル (IV) を入れるため、1ブロック分余計に確保する
	encrypted := make([]byte, aes.BlockSize+len(plainText))

	// IV は暗号文の先頭に入れておくことが多い
	iv := encrypted[:aes.BlockSize]
	// IV としてランダムなビット列を生成する
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatal(err)
	}

	// ブロック暗号として AES を使う場合
	key := []byte("secret-key-12345")
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	// CBC モードで暗号化する
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encrypted[aes.BlockSize:], plainText)
	fmt.Printf("encrypted: %x\n", encrypted)
}