package internal

import (
	"fmt"
	"unicode"
)

type Cipher interface {
	Name() string
	Description() string
	Key() string
	Encrypt(plaintext string) string
	Decrypt(ciphertext string) string
}

type CaesarCipher struct {
	Shift int
}

func (c *CaesarCipher) Name() string {
	return "Caesar Cipher"
}

func (c *CaesarCipher) Description() string {
	return "Caesar Cipher is a type of substitution cipher where each letter in the plaintext is replaced with a letter some fixed number of positions down the alphabet."
}

func (c *CaesarCipher) Key() string {
	return fmt.Sprintf("Shift: %d", c.Shift)
}

func (c *CaesarCipher) Encrypt(plaintext string) string {
	plaintextRunes := []rune(plaintext)
	for index, char := range plaintextRunes {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			if unicode.IsUpper(char) {
				plaintextRunes[index] = 'A' + (char-'A'+rune(c.Shift))%26
			} else {
				plaintextRunes[index] = 'a' + (char-'a'+rune(c.Shift))%26
			}
		}
	}
	return string(plaintextRunes)
}

func (c *CaesarCipher) Decrypt(ciphertext string) string {
	ciphertextRunes := []rune(ciphertext)
	for index, char := range ciphertextRunes {
		if char >= 'A' && char <= 'Z' {
			ciphertextRunes[index] = 'A' + (char-'A'-rune(c.Shift)+26)%26
		} else if char >= 'a' && char <= 'z' {
			ciphertextRunes[index] = 'a' + (char-'a'-rune(c.Shift)+26)%26
		}
	}
	return string(ciphertextRunes)
}
