package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	curve := elliptic.P256()
	priv, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Private key: %x\n", priv)

	message := "IST 402"
	fmt.Printf("Message: %s\n", message)

	// encrypt
	plaintext := []byte(message)
	k := new(big.Int).SetBytes(priv)              // convert private key to big.Int
	x1, _ := curve.ScalarBaseMult(k.Bytes())      // calculate public key
	y1, _ := curve.ScalarMult(x, y, k.Bytes())    // calculate shared secret
	ciphertext := elliptic.Marshal(curve, x1, y1) // encode public key
	ciphertext = append(ciphertext, plaintext...) // append plaintext to encoded public key
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	// decrypt
	x1, y1 = elliptic.Unmarshal(curve, ciphertext[:curve.Params().BitSize/8+1])
	sharedSecret, _ := curve.ScalarMult(x1, y1, priv)
	decrypted := ciphertext[curve.Params().BitSize/8+1:]
	for i := range decrypted {
		decrypted[i] ^= sharedSecret.Bytes()[i%len(sharedSecret.Bytes())]
	}
	fmt.Printf("Decrypted: %s\n", string(decrypted))
}
