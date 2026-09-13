package tests

import (
	"slices"
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

func TestAtbashCipherEncrypt(t *testing.T) {
	tests := []struct {
		plaintext  string
		ciphertext string
	}{
		{"Hello, World!", "Svool, Dliow!"},
		{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba"},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "ZYXWVUTSRQPONMLKJIHGFEDCBA"},
	}
	for _, test := range tests {
		cipher := &ciphers.AtbashCipher{}
		if cipher.Encrypt(test.plaintext) != test.ciphertext {
			t.Fatalf(">>>>> FAILED: Expected %s, got %s", test.ciphertext, cipher.Encrypt(test.plaintext))
		}
		t.Logf("Expected %s, got %s", test.ciphertext, cipher.Encrypt(test.plaintext))
		t.Logf("Test %s passed", test.plaintext)
	}
}

func TestAtbashCipherDecrypt(t *testing.T) {
	tests := []struct {
		ciphertext string
		plaintext  string
	}{
		{"Svool, Dliow!", "Hello, World!"},
		{"zyxwvutsrqponmlkjihgfedcba", "abcdefghijklmnopqrstuvwxyz"},
		{"ZYXWVUTSRQPONMLKJIHGFEDCBA", "ABCDEFGHIJKLMNOPQRSTUVWXYZ"},
	}
	for _, test := range tests {
		cipher := &ciphers.AtbashCipher{}
		if cipher.Decrypt(test.ciphertext) != test.plaintext {
			t.Fatalf(">>>>> FAILED: Expected %s, got %s", test.plaintext, cipher.Decrypt(test.ciphertext))
		}
		t.Logf("Expected %s, got %s", test.plaintext, cipher.Decrypt(test.ciphertext))
		t.Logf("Test %s passed", test.ciphertext)
	}
}

func TestNewKeywordCipher(t *testing.T) {
	tests := []struct {
		keyword    string
		dictionary []rune
	}{
		{"keyword", []rune("keywordabcfghijlmnpqstuvxz")},
		{"grandmother", []rune("grandmothebcfijklpqsuvwxyz")},
	}
	for _, test := range tests {
		cipher, err := ciphers.NewKeywordCipher(test.keyword)
		if err != nil {
			t.Fatalf(">>>>> FAILED: Expected nil, got %s", err)
		}
		if !slices.Equal(cipher.Dictionary, test.dictionary) {
			t.Fatalf(">>>>> FAILED: Expected %v, got %v", test.dictionary, cipher.Dictionary)
		}
		t.Logf("Expected %v, got %v", test.dictionary, cipher.Dictionary)
		t.Logf("Test %s passed", test.keyword)
	}
}

func TestKeywordCipherEncrypt(t *testing.T) {
	tests := []struct {
		keyword    string
		plaintext  string
		ciphertext string
	}{
		{"zebras", "Flee at once. WE are discovered!", "Siaa zq lkba. VA zoa rfpbluaoar!"},
		{"grandmother", "flee at OnCe. wE aRE diScovered!", "mcdd gs JiAd. wD gPD nhQajvdpdn!"},
	}
	for _, test := range tests {
		cipher, err := ciphers.NewKeywordCipher(test.keyword)
		if err != nil {
			t.Fatalf(">>>>> FAILED: Expected nil, got %s", err)
		}
		if cipher.Encrypt(test.plaintext) != test.ciphertext {
			t.Fatalf(">>>>> FAILED: Expected %s, got %s", test.ciphertext, cipher.Encrypt(test.plaintext))
		}
		t.Logf("Expected %s, got %s", test.ciphertext, cipher.Encrypt(test.plaintext))
		t.Logf("Test %s passed", test.plaintext)
	}
}

func TestKeywordCipherDecrypt(t *testing.T) {
	tests := []struct {
		keyword    string
		ciphertext string
		plaintext  string
	}{
		{"zebras", "Siaa zq lkba. VA zoa rfpbluaoar!", "Flee at once. WE are discovered!"},
		{"grandmother", "mcdd gs JiAd. wD gPD nhQajvdpdn!", "flee at OnCe. wE aRE diScovered!"},
	}
	for _, test := range tests {
		cipher, err := ciphers.NewKeywordCipher(test.keyword)
		if err != nil {
			t.Fatalf(">>>>> FAILED: Expected nil, got %s", err)
		}
		if cipher.Decrypt(test.ciphertext) != test.plaintext {
			t.Fatalf(">>>>> FAILED: Expected %s, got %s", test.plaintext, cipher.Decrypt(test.ciphertext))
		}
		t.Logf("Expected %s, got %s", test.plaintext, cipher.Decrypt(test.ciphertext))
		t.Logf("Test %s passed", test.ciphertext)
	}
}
