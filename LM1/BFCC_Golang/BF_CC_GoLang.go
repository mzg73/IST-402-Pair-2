package main

import (
	"fmt" // Need this package to print results
)

func main() {

	// Define the ciphertext to be decrypted
	ciphertext := "Aopz pz h alza!"

	// Try all possible shift values from 0 to 25
	for shift := 0; shift < 26; shift++ {

		// Initialize an empty string for the plaintext (we'll use this to store the "decrypted text" in!)
		plaintext := ""

		// Shift each letter in the ciphertext by the current shift value
		for _, char := range ciphertext { // "char" is assigned to the unicode value of the current letter of a string that was being iterated within the loop

			if char >= 'a' && char <= 'z' { // Unicode range for "a - z" is "97 to 122"

				// Shift lowercase letters
				shifted := ((char - 'a' + rune(shift)) % 26) + 'a'
				plaintext += string(shifted)

			} else if char >= 'A' && char <= 'Z' { // Unicode range for "A - Z" is "65 to 90"

				// Shift uppercase letters

				// To shift a char like "Z" by 0, it is like this:
				// [char - unicode val of 'A'] "90" -"65" = "25", [result of "char - unicode val of 'A'" + rune(shift)]
				// "25" + "0" = "25", [result of "char - unicode val of 'A' + rune(shift)" mod 26] "25" % "26" = "25",
				// [result of "char - 'A' + rune(shift) mod 26" + unicode val of 'A'] "25" + "65" = "90"
				shifted := ((char - 'A' + rune(shift)) % 26) + 'A'
				plaintext += string(shifted)

			} else {
				// Keep non-letter characters the way as they are
				plaintext += string(char)
			}
		}

		// Print the plaintext for the current shift value
		// Print format: "Shift [decimal of 'shift']: [string val of 'plaintext'] [start a new line]"
		fmt.Printf("Shift %d: %s\n", shift, plaintext)
	}
}
