package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
)

func main() {
	encrypted := make([]byte, aes.BlockSize+len(plainText))

	iv := encrypted[:aes.BlockSize]
}