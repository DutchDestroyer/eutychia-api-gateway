package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

// copied from https://golang.org/src/crypto/cipher/example_test.go

func GetEncryptedData(unencryptedText string, secretKey string) (string, string, error) {
	// Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
	key, err := hex.DecodeString(secretKey)

	if err != nil {
		return "", "", err
	}

	if len(key) != 32 {
		return "", "", errors.New("key length is not 32")
	}
	plaintext := []byte(unencryptedText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", "", err
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", "", err
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return string(ciphertext), string(nonce), nil
}

func ExampleNewGCM_decrypt(encryptedText string, nonceDB string, secretKey string) (string, error) {
	// Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
	key, err := hex.DecodeString(secretKey)
	if err != nil {
		return "", err
	}

	ciphertext, err := hex.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	nonce, err := hex.DecodeString(nonceDB)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
