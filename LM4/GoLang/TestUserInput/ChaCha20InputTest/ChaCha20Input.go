package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/chacha20"
	"os"
	"strings"
)

func main() {
	// Set the key (size of 32 bits) and nonce (size of 12 bits)
	key := make([]byte, chacha20.KeySize)
	copy(key[:], "I-Like-Cryptography")
	nonce := make([]byte, chacha20.NonceSize)
	copy(nonce[:], "I-Love-PSU")

	// Create the ChaCha20 cipher object. Sends back an error if either the key or nonce sizes are incorrect/unsupported.
	chacha, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		panic(err)
	}

	// Prompts the user to input the encrypted message
	fmt.Println("\nEnter the message to be encrypted:")
	reader := bufio.NewReader(os.Stdin)
	plaintext, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	plaintext = strings.TrimSuffix(plaintext, "\n")
	fmt.Println("\nPlain text:", string(plaintext))

	// Encrypt the message inputted from the user using the ChaCha20 cipher
	ciphertext := make([]byte, len(plaintext))
	chacha.XORKeyStream(ciphertext, []byte(plaintext))

	// Print the encrypted text as a hex string. Makes the encrypted output easier to read.
	fmt.Println("Encrypted text:", hex.EncodeToString(ciphertext))

	// Decrypt the ciphertext by using (or recreating) the same ChaCha20 cipher. Sends back an error if either the key or nonce sizes are incorrect/unsupported.
	decrypted := make([]byte, len(ciphertext))
	chacha, err = chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		panic(err)
	}
	chacha.XORKeyStream(decrypted, ciphertext)

	// Print the decrypted text to the console
	fmt.Println("Decrypted text:", string(decrypted))
}
