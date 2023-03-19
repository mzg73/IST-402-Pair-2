import string

ciphertext = "Mjqqt NXY 402"  # This is a ciphertext version of the plaintext "Hello IST 402"
key = 5  # The key to be guessed is 5
alphabet_lower = string.ascii_lowercase  # lowercase alphabet
alphabet_upper = string.ascii_uppercase  # uppercase alphabet


def decrypt(ciphertext, key): #This function is used to decrypt the ciphertext into plaintext
    plaintext = ""
    for char in ciphertext:
        if char in alphabet_lower:
            plaintext += alphabet_lower[(alphabet_lower.index(char) - key) % 26]
        elif char in alphabet_upper:
            plaintext += alphabet_upper[(alphabet_upper.index(char) - key) % 26]
        else:
            plaintext += char
    return plaintext


for i in range(0, 10): # The key is 5. The program will brute-force and try each number until it finds the key to display the message.
    if i == key:
        plaintext = decrypt(ciphertext, key)
        print(i, " was the key, brute force success")
        print("Decrypted message:", plaintext)
        break
    else:
        print(i, " is not the key")
