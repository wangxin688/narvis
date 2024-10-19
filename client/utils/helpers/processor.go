package helpers

import (
	"crypto/md5"
	"encoding/hex"
)

func StringToMd5(s string) string {
	hash := md5.New()
	_, _ = hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}
