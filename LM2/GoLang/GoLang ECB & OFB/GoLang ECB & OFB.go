// ECB and OFB Block Cipher Modes
package main

import (
	"fmt"
)

/* an array with 4 rows and 2 columns*/
var codebook = [4][2]int{{0b00, 0b01}, {0b01, 0b10}, {0b10, 0b11}, {0b11, 0b00}}
var codebook2 = [4][2]int{{0b00, 0b10}, {0b01, 0b11}, {0b10, 0b00}, {0b11, 0b01}}
var message = [4]int{0b00, 0b01, 0b10, 0b11}
var message2 = []int{}
var iv int = 0b10

func codebookLookup(codebookNum int, xor int) (lookupValue int) {
	var i, j int = 0, 0
	for i = 0; i < 4; i++ {
		if codebook[i][j] == xor {
			j++
			if codebookNum == 2 {
				lookupValue = codebook2[i][j]
			} else {
				lookupValue = codebook[i][j]
			}
			break
		}
	}
	return lookupValue
}

func main() {
	// ECB Mode
	fmt.Printf("ECB: \n")
	var x, i int = 0, 0
	var lookupValue int = 0
	for i = 0; i < 4; i++ {
		lookupValue = codebookLookup(2, message[x])
		fmt.Printf("The ciphered value of %b is %b\n", message[x], lookupValue)
		message2 = append(message2, lookupValue)
		x++
	}
	x = 0
	for i = 0; i < 4; i++ {
		lookupValue = codebookLookup(2, message2[x])
		fmt.Printf("The plaintext value of %b is %b\n", message2[x], lookupValue)
		x++
	}

	// OFB Mode
	fmt.Printf("\nOFB: \n")
	var y int = 0
	var keystream int = iv
	var output int
	var ciphertext [4]int
	var plaintext [4]int
	// Encryption
	for y = 0; y < 4; y++ {
		keystream = codebookLookup(2, keystream)
		output = keystream ^ message[y]
		ciphertext[y] = output
		fmt.Printf("The ciphered value of %b is %b\n", message[y], output)
	}
	// Decryption
	for y = 0; y < 4; y++ {
		keystream = codebookLookup(2, keystream)
		plaintext[y] = keystream ^ ciphertext[y]
		fmt.Printf("The plaintext value of %b is %b\n", ciphertext[y], plaintext[y])
	}

}
