package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

var key = []byte("asdfasdfasdfasdfasdfasdfasdfasdf")

func Encrypt(date string) string {
	plaintext := []byte(date)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	nonce := []byte("gopostmedium")
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}
func Decrypt(date string) string {
	ciphertext, _ := hex.DecodeString(date)
	nonce := []byte("gopostmedium")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return fmt.Sprintf("%s", plaintext)
}
func main() {
	str := "my string"
	strEncrypted := Encrypt(str)
	fmt.Println("encrypt: ", strEncrypted)
	fmt.Printf("decrypt: %s", Decrypt(strEncrypted))
}
