//package Aes_crypt
package Aes_crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

// func encryptAes1(data []byte) (a string, b string, err error) {
	
// }

func EncryptByAes(data []byte) (a string, b string, err error) {
	// Generate a random key
	key := make([]byte, sha256.Size)
	if _, err := rand.Read(key); err != nil {
		fmt.Println(err)
		return "", "", err
	}

	// Hash the key with SHA-256 for added security
	Hashedkey := sha256.Sum256(key)

	// Create a new AES cipher
	fmt.Println(Hashedkey)
	block, err := aes.NewCipher([]byte(Hashedkey[:]))
	if err != nil {
		return "", "", err
	}

	// Create a new GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", "", err
	}

	// Create a nonce
	nonce := make([]byte, gcm.NonceSize())
	fmt.Println("nonce size :",nonce)
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", "", err
	}

	// Encrypt the data
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	fmt.Println(ciphertext)

	// Return the encrypted data as a base64 encoded string
	return base64.StdEncoding.EncodeToString(ciphertext), base64.StdEncoding.EncodeToString(Hashedkey[:]), nil
}

func DecryptAES(encryptedData string, Base64key string) ([]byte,string, error) {
	// Decode the encrypted data from base64
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil,"error in decode the message", err
	}
	fmt.Println(ciphertext)

	key, err := base64.StdEncoding.DecodeString(Base64key)
	if err != nil {
		return nil,"", err
	}

	// Create a new AES cipher
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil,"error in new AES", err
	}

	// Create a new GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil,"error in new GCM", err
	}

	// Extract the nonce from the ciphertext
	nonceSize := gcm.NonceSize()
	nonce := ciphertext[:nonceSize]
	ciphertext = ciphertext[nonceSize:]

	// Decrypt the data
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil,"error in decrpt the message", err
	}

	return plaintext,"", nil
}
