package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func Encrypt(stringToEncrypt string) (string, error) {
	var keyString string = ""
	keyStringValue, keyStringPresent := os.LookupEnv("ENCRYPT_KEY")
	if keyStringPresent {
		keyString = keyStringValue
	}

	key, err := hex.DecodeString(keyString)
	if err != nil {
		return "", err
	}

	plainByte := []byte(stringToEncrypt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	encryptedByte := aesGCM.Seal(nonce, nonce, plainByte, nil)
	return fmt.Sprintf("%x", encryptedByte), nil
}

func Decrypt(encryptedString string) (string, error) {
	var keyString string = ""
	keyStringValue, keyStringPresent := os.LookupEnv("ENCRYPT_KEY")
	if keyStringPresent {
		keyString = keyStringValue
	}

	var plainString string

	key, err := hex.DecodeString(keyString)
	if err != nil {
		return plainString, err
	}

	enc, err := hex.DecodeString(encryptedString)
	if err != nil {
		return plainString, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return plainString, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return plainString, err
	}

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]
	plainByte, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return plainString, err
	}

	plainString = string(plainByte)

	return plainString, nil
}
