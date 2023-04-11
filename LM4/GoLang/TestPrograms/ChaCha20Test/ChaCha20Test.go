package main

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/chacha20"
)

func main() {
	// Set the key and nonce
	key := make([]byte, chacha20.KeySize)
	copy(key[:], []byte("my-secret-key-123456789012345"))
	nonce := make([]byte, chacha20.NonceSize)
	copy(nonce[:], []byte("my-unique-nonce"))

	// Create the ChaCha20 cipher object
	chacha, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		panic(err)
	}

	// Set the plaintext message
	plaintext := []byte("my secret message")

	// Encrypt the plaintext using the ChaCha20 cipher
	ciphertext := make([]byte, len(plaintext))
	chacha.XORKeyStream(ciphertext, plaintext)

	// Print the encrypted text as a hex string
	fmt.Println("Encrypted text:")
	fmt.Println(hex.EncodeToString(ciphertext))

	// Decrypt the ciphertext using the same ChaCha20 cipher
	decrypted := make([]byte, len(ciphertext))
	chacha, err = chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		panic(err)
	}
	chacha.XORKeyStream(decrypted, ciphertext)

	// Print the decrypted text
	fmt.Println("Decrypted text:")
	fmt.Println(string(decrypted))
}
