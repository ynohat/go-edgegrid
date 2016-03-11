package edgegrid

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io"
)

// ComputeHmac256 encrypts message with hmac sha 256
func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	io.WriteString(h, message)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// Compute256 encrypts message with hmac sha 256
func Compute256(message string) string {
	h := sha256.New()
	io.WriteString(h, message)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
