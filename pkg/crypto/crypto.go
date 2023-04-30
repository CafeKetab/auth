package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"strings"
)

type Crypto interface {
	Encrypt(plainText string) (string, error)
	Decrypt(cipherText string) (string, error)
}

type crypto struct {
	config *Config
}

func New(cfg *Config) Crypto {
	return &crypto{config: cfg}
}

// procedure is as follow:
//
// 1. plainText + salt -> binary format
//
// 2. encrypt aes in CTR mode
//
// 3. base64 encode
func (c *crypto) Encrypt(plainText string) (string, error) {
	binaryText := []byte(plainText + c.config.Salt)
	binarySecret := []byte(c.config.Secret)

	// Create new AES cipher block
	block, err := aes.NewCipher(binarySecret)
	if err != nil {
		return "", err
	}

	// The IV (Initialization Vector) need to be unique, but not secure.
	// Therefore, it's common to include it at the beginning of the cipher text.
	cipherText := make([]byte, aes.BlockSize+len(binaryText))

	// Create IV
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Encrypt
	encryptStream := cipher.NewCTR(block, iv)
	encryptStream.XORKeyStream(cipherText[aes.BlockSize:], binaryText)

	return base64.URLEncoding.EncodeToString(cipherText), nil
}

// procedure is as follow:
//
// 1. base64 decode
//
// 2. decrypt aes in CTR mode
//
// 3. remove salt
func (c *crypto) Decrypt(cipherText string) (string, error) {
	binaryCipherText, err := base64.URLEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	binarySecret := []byte(c.config.Secret)

	// Create new AES cipher block
	block, err := aes.NewCipher(binarySecret)
	if err != nil {
		return "", err
	}

	// Decrpt
	decryptedText := make([]byte, len(binaryCipherText[aes.BlockSize:]))
	decryptStream := cipher.NewCTR(block, binaryCipherText[:aes.BlockSize])
	decryptStream.XORKeyStream(decryptedText, binaryCipherText[aes.BlockSize:])

	return strings.TrimSuffix(string(decryptedText), c.config.Salt), nil
}
