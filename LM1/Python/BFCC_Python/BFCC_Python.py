# Project: Brute Force using Python
# Author: Marcus Gonzalez
# Course: IST 402
# Date Developed: 3/19/2023
# Date Last Changed: 3/19/2023

import string

ciphertext = "Mjqqt NXY 402"  # This is a ciphertext version of the plaintext "Hello IST 402"
alphabet_lower = string.ascii_lowercase  # lowercase alphabet
alphabet_upper = string.ascii_uppercase  # uppercase alphabet


def decrypt(ciphertext, shift):
    plaintext = ""
    for char in ciphertext:
        if char in alphabet_lower:
            shifted = ((ord(char) - ord('a') + shift) % 26) + ord('a')
            plaintext += chr(shifted)
        elif char in alphabet_upper:
            shifted = ((ord(char) - ord('A') + shift) % 26) + ord('A')
            plaintext += chr(shifted)
        else:
            plaintext += char
    return plaintext


for shift in range(26):
    plaintext = decrypt(ciphertext, shift)
    print(f"Shift {shift}: {plaintext}")
