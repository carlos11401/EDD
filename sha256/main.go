package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	str := "my string"
	fmt.Println(getSha256(str))
}
func getSha256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
