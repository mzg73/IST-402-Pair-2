// Structures used are from LM5 Assignment
package main

import (
	_ "encoding/base64"
	"fmt"
)

// Reflector represents the fixed reflector
type Reflector struct {
	wiring [26]int
}

// Plugboard structure
type Plugboard struct {
	wiring [26]int
}

// EnigmaMachine structure
type EnigmaMachine struct {
	plugboard  *Plugboard
	reflector  *Reflector
	rotorSet   *RotorSet
	inputRotor *InputRotor
}

func (s RotorSet) rotate(r *Rotor, steps int) {
	for i := 0; i < steps; i++ {
		r.position = (r.position + 1) % 26
		r.wiring[0], r.wiring[25] = r.wiring[25], r.wiring[0]
		r.wiring[0], r.wiring[r.position] = r.wiring[r.position], r.wiring[0]
		r.wiring[25], r.wiring[r.position] = r.wiring[r.position], r.wiring[25]
	}
}

// encoded method for EnigmaMachine
func (e *EnigmaMachine) encoded(c byte) byte {
	// Pass the character through the plugboard
	c = byte(e.plugboard.wiring[c-'A']) + 'A'

	// Rotate the rotors (One motor isn't dependent on another to rotate; they all rotate whenever each letter is analyzed)
	e.rotorSet.rotate(e.rotorSet.Rightrotor, 6)
	e.rotorSet.rotate(e.rotorSet.Middlerotor, 3)
	e.rotorSet.rotate(e.rotorSet.Leftrotor, 2)

	// Pass the character through the input rotor
	c = byte(e.inputRotor.wiring[c-'A']) + 'A'

	// Pass the character through the rotor set (right to left)
	c = byte(e.rotorSet.Rightrotor.wiring[(int(c)-'A'+e.rotorSet.Rightrotor.position)%26]) + 'A'
	c = byte(e.rotorSet.Middlerotor.wiring[(int(c)-'A'+e.rotorSet.Middlerotor.position)%26]) + 'A'
	c = byte(e.rotorSet.Leftrotor.wiring[(int(c)-'A'+e.rotorSet.Leftrotor.position)%26]) + 'A'

	// Pass the character through the reflector
	c = byte(e.reflector.wiring[c-'A']) + 'A'

	// Pass the character back through the rotor set (left to right)
	c = byte(e.rotorSet.Leftrotor.wiring[(int(c)-'A'-e.rotorSet.Leftrotor.position+26)%26]) + 'A'
	c = byte(e.rotorSet.Middlerotor.wiring[(int(c)-'A'-e.rotorSet.Middlerotor.position+26)%26]) + 'A'
	c = byte(e.rotorSet.Rightrotor.wiring[(int(c)-'A'-e.rotorSet.Rightrotor.position+26)%26]) + 'A'

	// Pass the character back through the input rotor
	c = byte(e.inputRotor.wiring[c-'A']) + 'A'

	// Pass the character back through the plug board
	c = byte(e.plugboard.wiring[c-'A']) + 'A'

	return c
}

// Rotor structure
type Rotor struct {
	wiring   [26]int
	position int
}

// RotorSet structure
type RotorSet struct {
	Rightrotor  *Rotor
	Leftrotor   *Rotor
	Middlerotor *Rotor
}

// InputRotor Input Rotor structure
type InputRotor struct {
	wiring [26]int
}

// XOR each letter of the string
func xorString(s string, key string) string {
	// Convert the string and key to byte slices
	sBytes := []byte(s)
	keyBytes := []byte(key)

	// XOR each byte in the string with the corresponding byte in the key
	xoredBytes := make([]byte, len(sBytes))
	for i := 0; i < len(sBytes); i++ {
		xoredBytes[i] = sBytes[i] ^ keyBytes[i%len(keyBytes)]
	}

	// Convert the XORed bytes back to a string
	xoredString := string(xoredBytes)

	return xoredString
}

func main() {

	// Set up the components
	reflector := &Reflector{[26]int{4, 10, 12, 5, 11, 6, 3, 16, 21, 25, 13, 19, 14, 22, 24, 7, 23, 20, 18, 15, 0, 8, 1, 17, 2, 9}}
	plugboard := &Plugboard{[26]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}}
	rightRotor := &Rotor{[26]int{3, 7, 4, 20, 13, 16, 10, 22, 15, 2, 12, 19, 25, 5, 14, 23, 6, 24, 18, 21, 8, 1, 11, 17, 0, 9}, 0}
	middleRotor := &Rotor{[26]int{8, 12, 4, 19, 2, 6, 5, 17, 0, 24, 18, 16, 7, 23, 20, 22, 21, 25, 14, 10, 3, 13, 1, 11, 15, 9}, 0}
	leftRotor := &Rotor{[26]int{2, 4, 6, 8, 10, 12, 3, 16, 18, 20, 24, 22, 26, 14, 25, 5, 9, 23, 7, 1, 11, 13, 21, 19, 17, 15}, 0}
	inputRotor := &InputRotor{[26]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}}

	// Create the rotor set and enigma machine
	rotorSet := &RotorSet{rightRotor, leftRotor, middleRotor}
	enigma := &EnigmaMachine{plugboard, reflector, rotorSet, inputRotor}

	// Input message to encode
	message := "ISTFOURZEROTWO"
	fmt.Println("Original message: ", message)

	// Encodes the message
	encodedMessage := ""
	for i := 0; i < len(message); i++ {
		encodedChar := enigma.encoded(message[i])
		encodedMessage += string(encodedChar)
	}

	//XOR the final string (the special feature)
	encodedMessage = xorString(encodedMessage, "secret")
	fmt.Println("Encoded message: ", encodedMessage)

	//XOR the final string (reversing the initial XOR function)
	encodedMessage = xorString(encodedMessage, "secret")

	// Decodes the message (Shows that it is not the same as the plaintext)
	// Note: We wanted to find a way to reset the rotors so that the decryption would show the actual plaintext, but we got stumped/ran out of time
	decodedMessage := ""
	for i := 0; i < len(encodedMessage); i++ {
		decodedChar := enigma.encoded(encodedMessage[i])
		decodedMessage += string(decodedChar)
	}

	// Print the encoded and decoded messages
	fmt.Println("Decoded message: ", decodedMessage)
}
