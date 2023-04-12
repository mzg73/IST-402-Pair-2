package main

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/chacha20"
)

func main() {
	// Set the key (size of 32 bits) and nonce (size of 12 bits)
	key := make([]byte, chacha20.KeySize)
	copy(key[:], []byte("I-Like-Cryptography"))
	nonce := make([]byte, chacha20.NonceSize)
	copy(nonce[:], []byte("I-Love-PSU"))

	// Create the ChaCha20 cipher object. Sends back an error if either the key or nonce sizes are incorrect/unsupported.
	chacha, err := chacha20.NewUnauthenticatedCipher(key, nonce) // "Note that ChaCha20, like all stream ciphers, is not authenticated and allows attackers to silently tamper with the plaintext."
	if err != nil {
		panic(err)
	}

	// Set the plaintext message
	plaintext := []byte("Super secret message")
	fmt.Println("\nPlain text:", string(plaintext))

	// Encrypt the plaintext using the ChaCha20 cipher
	ciphertext := make([]byte, len(plaintext))
	chacha.XORKeyStream(ciphertext, plaintext)

	// Print the encrypted text as a hex string. Makes the encrypted output easier to read.
	fmt.Println("Encrypted text:", hex.EncodeToString(ciphertext)) // "Note that ChaCha20, like all stream ciphers, is not authenticated and allows attackers to silently tamper with the plaintext."

	// Decrypt the ciphertext by using (or recreating) the same ChaCha20 cipher. Sends back an error if either the key or nonce sizes are incorrect/unsupported.
	decrypted := make([]byte, len(ciphertext))
	chacha, err = chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		panic(err)
	}
	chacha.XORKeyStream(decrypted, ciphertext)

	// Print the decrypted text
	fmt.Println("Decrypted text:", string(decrypted))
}
