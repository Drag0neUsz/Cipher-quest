package tests

import (
	"testing"

	ciphers "github.com/Drag0neUsz/Cipher-quest/internal/content"
)

func TestCaesarCipherEncrypt(t *testing.T) {
	tests := []struct {
		shift      int
		plaintext  string
		ciphertext string
	}{
		{1, "Hello, World!", "Ifmmp, Xpsme!"},
		{1, "abcdefghijklmnopqrstuvwxyz", "bcdefghijklmnopqrstuvwxyza"},
		{1, "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "BCDEFGHIJKLMNOPQRSTUVWXYZA"},
		{26, "ABCdef", "ABCdef"},
		{25, "abcDEF", "zabCDE"},
		{13, "hello", "uryyb"},
		{0, "hello", "hello"},
		{27, "hello", "ifmmp"},
	}
	for _, test := range tests {
		cipher := &ciphers.CaesarCipher{Shift: test.shift}
		if cipher.Encrypt(test.plaintext) != test.ciphertext {
			t.Fatalf(">>>>> FAILED: Expected %s, got %s", test.ciphertext, cipher.Encrypt(test.plaintext))
		}
		t.Logf("Expected %s, got %s", test.ciphertext, cipher.Encrypt(test.plaintext))
		t.Logf("Test %s with shift %d passed", test.plaintext, test.shift)
	}
}

func TestCaesarCipherDecrypt(t *testing.T) {
	tests := []struct {
		shift      int
		ciphertext string
		plaintext  string
	}{
		{1, "Ifmmp, Xpsme!", "Hello, World!"},
		{1, "bcdefghijklmnopqrstuvwxyza", "abcdefghijklmnopqrstuvwxyz"},
		{1, "BCDEFGHIJKLMNOPQRSTUVWXYZA", "ABCDEFGHIJKLMNOPQRSTUVWXYZ"},
		{26, "ABCdef", "ABCdef"},
		{25, "zabCDE", "abcDEF"},
		{13, "uryyb", "hello"},
		{0, "hello", "hello"},
		{27, "ifmmp", "hello"},
	}
	for _, test := range tests {
		cipher := &ciphers.CaesarCipher{Shift: test.shift}
		if cipher.Decrypt(test.ciphertext) != test.plaintext {
			t.Fatalf(">>>>> FAILED: Expected %s, got %s", test.plaintext, cipher.Decrypt(test.ciphertext))
		}
		t.Logf("Expected %s, got %s", test.plaintext, cipher.Decrypt(test.ciphertext))
		t.Logf("Test %s with shift %d passed", test.ciphertext, test.shift)
	}
}
