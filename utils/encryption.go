package utils

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
)

func DecodeBase64(uniqueID string, message string) *string {
	data, err := b64.StdEncoding.DecodeString(message)
	if err != nil {
		log.Errorf(uniqueID, "Error occurred "+err.Error())
		return nil
	}
	result := string(data)
	return &result
}

func EncryptionSha256(data string) string {
	hash := sha256.Sum256([]byte(data))
	strResult := fmt.Sprintf("%x", hash[:])
	return strResult
}
