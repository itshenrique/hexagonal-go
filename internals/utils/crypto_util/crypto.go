package crypto_util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password *string) {
	passwordBytes := []byte(*password)

	hashedData, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	*password = string(hashedData)
}

func ComparePassword(hashedPassword *string, password *string) bool {
	// Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword([]byte(*hashedPassword), []byte(*password))
	return err == nil
}

func HashMd5(text []byte) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func Encrypt(plaintext string, key string) (*string, error) {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	stringValue := string(gcm.Seal(nonce, nonce, []byte(plaintext), nil))

	return &stringValue, nil
}

func Decrypt(ciphertext string, key string) (*string, error) {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	value, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)

	if err != nil {
		return nil, err
	}

	stringValue := string(value)

	return &stringValue, nil
}
