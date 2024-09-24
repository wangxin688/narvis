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

// 将字符串转换为指针类型的字符串，当字符串值为""零值时，转换为nil
// background：基础的sdk统一定义了结构体基础的字段没有指针类型
func StringToPtrString(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}

func PtrStringToString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
