package cipher

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "errors"
    "io"
)

func encodeBase64(b []byte) string {
    return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
    data, err := base64.StdEncoding.DecodeString(s)
    
    if err != nil {
        panic(err)
    }
    
    return data
}

func Encrypt(key string, text string) (string, error) {
    block, err := aes.NewCipher([]byte(key))
    
    if err != nil {
        return "", err
    }
    
    b := base64.StdEncoding.EncodeToString([]byte(text))
    cipt := make([]byte, aes.BlockSize+len(b))
    iv := cipt[:aes.BlockSize]
    
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }
    
    cfb := cipher.NewCFBEncrypter(block, iv)
    
    cfb.XORKeyStream(cipt[aes.BlockSize:], []byte(b))
    
    return encodeBase64(cipt), nil
}

func Decrypt(key string, t string) (string, error) {
    block, err := aes.NewCipher([]byte(key))
    
    if err != nil {
        return "", err
    }
    
    text := decodeBase64(t)
    
    if len(text) < aes.BlockSize {
        return "", errors.New("too short")
    }
    
    iv := text[:aes.BlockSize]
    text = text[aes.BlockSize:]
    cfb := cipher.NewCFBDecrypter(block, iv)
    
    cfb.XORKeyStream(text, text)
    
    data, err := base64.StdEncoding.DecodeString(string(text))
    
    if err != nil {
        return "", err
    }
    
    return string(data), nil
}