package helpers

import (
	"crypto/rand"
	"encoding/hex"
)

func RandomHexString(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}
