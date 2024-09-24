package tasks

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
