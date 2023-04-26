package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	message := "IST 402 message"
	key, _ := hex.DecodeString("6368616e676520746869732070617373")
	plaintext := []byte(message)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	nonce := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	stream := cipher.NewCTR(block, nonce)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	fmt.Printf("Message: %s\n", message)
	fmt.Printf("Ciphertext: %x\n", ciphertext[aes.BlockSize:])

	plaintext2 := make([]byte, len(plaintext))

	stream = cipher.NewCTR(block, nonce)
	stream.XORKeyStream(plaintext2, ciphertext[aes.BlockSize:])

	fmt.Printf("Decrypted message: %s\n", plaintext2)
}
