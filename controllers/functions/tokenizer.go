package functions

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var secretKey = []byte("my32digitkey12345678901234567890")

// var (
// 	// We're using a 32 byte long secret key.
// 	// This is probably something you generate first
// 	// then put into and environment variable.
// 	secretKey string = "N1PCdw3M2B1TfJhoaY2mL736p2vCUc47"
// )

func GetAESDecrypted(encrypted string, nonce string) (string, error) {
	// key := "my32digitkey12345678901234567890"
	// iv := "my16digitIvKey12"

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	decodedCipherText, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	decodedNonce, err := base64.StdEncoding.DecodeString(nonce)
	if err != nil {
		return "", err
	}

	plainText, err := aesGCM.Open(nil, decodedNonce, decodedCipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
