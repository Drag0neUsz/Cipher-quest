package content

import (
	"fmt"
	"unicode"
)

type Cipher interface {
	Name() string
	Key() string
	Encrypt(plaintext string) string
	Decrypt(ciphertext string) string
}

type CaesarCipher struct {
	Shift int
}

func (c CaesarCipher) Name() string {
	return "Caesar Cipher"
}

func (c CaesarCipher) Key() string {
	return fmt.Sprintf("Shift: %d", c.Shift)
}

func (c CaesarCipher) Encrypt(plaintext string) string {
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

func (c CaesarCipher) Decrypt(ciphertext string) string {
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

type AtbashCipher struct {
}

func (c AtbashCipher) Name() string {
	return "Atbash Cipher"
}

func (c AtbashCipher) Description() string {
	return `Atbash Cipher is a substitution cipher where each letter in the plaintext is replaced with the letter that is the same distance from the end of the alphabet.

	Example:
	Plaintext: HELLO
	Ciphertext: SVOOL
	`
}

func (c AtbashCipher) Key() string {
	return "No key required"
}

func (c AtbashCipher) Encrypt(plaintext string) string {
	plaintextRunes := []rune(plaintext)
	for index, char := range plaintextRunes {
		if char >= 'A' && char <= 'Z' {
			plaintextRunes[index] = 'A' + ('Z'-char)%26
		} else if char >= 'a' && char <= 'z' {
			plaintextRunes[index] = 'a' + ('z'-char)%26
		}
	}
	return string(plaintextRunes)
}

func (c AtbashCipher) Decrypt(ciphertext string) string {
	return c.Encrypt(ciphertext) // Atbash is symmetric
}
