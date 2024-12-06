package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter password to encrypt: ")
	password, _ := reader.ReadString('\n')
	encrypted := encrypt(strings.TrimSpace(password), "secret_key_123456")
	fmt.Println("Encrypted password:", encrypted)
}

func encrypt(text, key string) string {
	block, _ := aes.NewCipher([]byte(key))
	nonce := make([]byte, 12)
	rand.Read(nonce)
	aesgcm, _ := cipher.NewGCM(block)
	return hex.EncodeToString(aesgcm.Seal(nil, nonce, []byte(text), nil))
}
