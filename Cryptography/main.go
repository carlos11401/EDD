package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

var cipherStr string

func encrypt() {
	key := []byte("asdfasdfasdfasdfasdfasdfasdfasdf")
	plaintext := []byte("This is the plaintext to be encrypted")
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
	fmt.Printf("Ciphertext: %x\n", ciphertext)
	cipherStr = fmt.Sprintf("%x", ciphertext)
}
func decrypt() {
	key := []byte("keygopostmediumkeygopostmediumke")
	//test := "13ca135cef69048ae33a21f8f4d52360c3e2f640a73ba46d9633e0b092dec4931689cc0fa225cbc66eeb7d1e27472a494a0183d6b5"
	//test = cipherStr
	ciphertext, _ := hex.DecodeString("2fcd164e91cfe4a958b67110d8f5efb89ddbdb00")
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
	fmt.Printf("Plaintext: %s\n", string(plaintext))
	fmt.Println(fmt.Sprintf("%s", plaintext))
}
func main() {
	encrypt()
	decrypt()
}
