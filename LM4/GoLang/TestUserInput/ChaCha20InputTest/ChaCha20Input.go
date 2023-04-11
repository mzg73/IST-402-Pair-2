package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/chacha20"
)

func main() {
	// Sets the key and nonce
	key := make([]byte, chacha20.KeySize)
	copy(key[:], "my-secret-key-123456789012345")
	nonce := make([]byte, chacha20.NonceSize)
	copy(nonce[:], "my-unique-nonce")

	// Create the ChaCha20 cipher object
	chacha, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		panic(err)
	}

	// Prompts the user to input the encrypted message
	fmt.Println("Enter the message to be encrypted:")
	reader := bufio.NewReader(os.Stdin)
	plaintext, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	plaintext = strings.TrimSuffix(plaintext, "\n")

	// Encrypt the message inputted from the user using the ChaCha20 cipher
	ciphertext := make([]byte, len(plaintext))
	chacha.XORKeyStream(ciphertext, []byte(plaintext))

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

	// Print the decrypted text to the console
	fmt.Println("Decrypted text:")
	fmt.Println(string(decrypted))
}
