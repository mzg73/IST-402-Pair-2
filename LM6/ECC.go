package main

import (
	"bufio"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Chooses the elliptical curve
	curve := elliptic.P224()

	// Generates the private keys for the sender and recipient
	senderPrivateKey, _ := ecdsa.GenerateKey(curve, rand.Reader)
	recipientPrivateKey, _ := ecdsa.GenerateKey(curve, rand.Reader)

	// Obtains the recipient's public key
	recipientPublicKey := recipientPrivateKey.PublicKey

	// Performs ECDH key exchange to establish a shared secret. The shared secret is actually a new x-coordinate that was
	// calculated by scalar multiplying the recipient's public key coordinates with the sender's private key's secret integer (the "D" value)
	sharedX, _ := curve.ScalarMult(recipientPublicKey.X, recipientPublicKey.Y, senderPrivateKey.D.Bytes())

	// Prompts the user to input the encrypted message
	fmt.Println("\nEnter the message to be encrypted:")
	reader := bufio.NewReader(os.Stdin)
	plaintext, _ := reader.ReadString('\n')
	plaintext = strings.TrimSuffix(plaintext, "\n")

	// Applies XOR encryption to the plaintext message using the shared secret
	fmt.Println("\nPlain text:", string(plaintext))
	ciphertext := make([]byte, len(plaintext))
	encryptionKey := sharedX.Bytes()
	for i := range plaintext {
		ciphertext[i] = plaintext[i] ^ encryptionKey[i%len(encryptionKey)]
	}

	// Print the encrypted text as a hex string. Makes the encrypted output easier to read.
	fmt.Println("Encrypted text:", hex.EncodeToString(ciphertext))

	// Applies XOR decryption to the ciphertext using the shared secret
	decryptionKey := sharedX.Bytes()
	decrypted := make([]byte, len(ciphertext))
	for i := range ciphertext {
		decrypted[i] = ciphertext[i] ^ decryptionKey[i%len(decryptionKey)]
	}

	// Print the decrypted text
	fmt.Println("Decrypted text:", string(decrypted))
}
