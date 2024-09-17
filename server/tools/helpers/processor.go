package helpers

import (
	"crypto/md5"
	"encoding/hex"
)

func EnsureSliceNotNil[T any](ptr *[]T) []T {
	if ptr == nil {
		var emptySlice []T
		return emptySlice
	}
	return *ptr
}

func StringToMd5(s string) string {
	hash := md5.New()
	_, _ = hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

func ByteToMd5(b []byte) string {
	hash := md5.New()
	_, _ = hash.Write(b)
	return hex.EncodeToString(hash.Sum(nil))
}
