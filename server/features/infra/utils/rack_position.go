package infra_utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wangxin688/narvis/server/tools/errors"
)

func ParseUint8s(s string) ([]uint8, error) {
	// 去除首尾空格
	s = strings.TrimSpace(s)
	// 使用逗号分割字符串
	parts := strings.Split(s, ",")
	var result []uint8

	for _, part := range parts {
		// 去除每部分前后的空格
		part = strings.TrimSpace(part)
		// 尝试将每部分转换为uint8
		num, err := strconv.ParseUint(part, 10, 8) // 第三个参数8指定了结果必须小于256（uint8的范围）
		if err != nil {
			return nil, fmt.Errorf("invalid uint8 value: %q", part)
		}
		// 将uint64转换为uint8（这里假设转换是安全的，因为我们已经限制了范围）
		result = append(result, uint8(num))
	}

	return result, nil
}

func SliceUint8ToString(s []uint8) (string, error) {
	if !isConsecutive(s) {
		return "", errors.NewError(errors.CodeRackPositionInconsecutive, errors.MsgRackPositionInconsecutive)
	}
	var result []string
	for _, v := range s {
		result = append(result, fmt.Sprintf("%d", v))
	}
	return strings.Join(result, ","), nil
}

func isConsecutive(slice []uint8) bool {
	if len(slice) <= 1 {
		return true // 空切片或只有一个元素，视为既不递增也不递减
	}

	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(slice); i++ {
		if slice[i] == slice[i-1]+1 {
			isDecreasing = false
		} else if slice[i] == slice[i-1]-1 {
			isIncreasing = false
		} else {
			// 如果当前元素既不满足递增也不满足递减，则直接返回false
			return false
		}
	}

	if isIncreasing {
		return true
	} else if isDecreasing {
		return true
	}

	return false
}
