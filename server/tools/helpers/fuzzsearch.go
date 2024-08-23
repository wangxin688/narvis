package helpers

import (
	"regexp"
	"strings"
)

// FuzzySearch 在 map[string]any 中递归搜索指定值
// - `data` 是待搜索的 map
// - `searchValue` 是要搜索的值
// - `ignoreCase` 指定是否忽略大小写
// - `keys` 是指定搜索的键列表，如果为空，则搜索所有键
func FuzzySearch(data map[string]any, searchValue string, ignoreCase bool, keys []string) bool {
	matchingKeys := []string{}

	// 如果忽略大小写，将搜索值转换为小写
	if ignoreCase {
		searchValue = strings.ToLower(searchValue)
	}

	// 如果 keys 列表为空，搜索所有键
	if len(keys) == 0 {
		for key := range data {
			keys = append(keys, key)
		}
	}

	for _, key := range keys {
		if value, exists := data[key]; exists {
			matchingKeys = append(matchingKeys, searchInValue(value, searchValue, ignoreCase, key)...)
		}
	}

	return len(matchingKeys) > 0
}

// searchInValue 递归搜索单个值
func searchInValue(value any, searchValue string, ignoreCase bool, currentKey string) []string {
	matchingKeys := []string{}

	switch v := value.(type) {
	case string:
		// 将字符串值转换为小写以进行大小写不敏感的搜索
		if ignoreCase {
			v = strings.ToLower(v)
		}
		// 模糊搜索
		if strings.Contains(v, searchValue) {
			matchingKeys = append(matchingKeys, currentKey)
		}
	case map[string]any:
		// 递归搜索嵌套的 map
		for k, nestedValue := range v {
			matchingKeys = append(matchingKeys, searchInValue(nestedValue, searchValue, ignoreCase, currentKey+"."+k)...)
		}
	}

	return matchingKeys
}

// MatchAnyRegex 检查字符串是否匹配任意正则表达式, 如果传递的正则表达式无效，则返回 false
func MatchAnyRegex(value string, regex []string) bool {
	compiledRegex := make([]*regexp.Regexp, len(regex))
	for i, r := range regex {
		if !checkRegexValid(r) {
			return false
		}
		compiledRegex[i] = regexp.MustCompile(r)
	}
	for _, r := range compiledRegex {
		if r.MatchString(value) {
			return true
		}
	}
	return false
}

func checkRegexValid(pattern string) bool {
	_, err := regexp.Compile(pattern)
	return err == nil
}
