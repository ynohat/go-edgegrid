package edgegrid

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

// ComputeHmac256 encrypts message with hmac sha 256
func ComputeHmac256(message string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// Compute256 encrypts message with hmac sha 256
func Compute256(message string) string {
	h := sha256.New()
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
