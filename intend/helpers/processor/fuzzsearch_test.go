package processor_test

import (
	"testing"

	"github.com/wangxin688/narvis/intend/helpers/processor"
)

func TestFuzzySearch(t *testing.T) {
	// Test case: Searching for an empty string in an empty map
	data := map[string]any{}
	searchValue := ""
	ignoreCase := false
	keys := []string{}
	expected := false
	actual := processor.FuzzySearch(data, searchValue, ignoreCase, keys)
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	// Test case: Searching for a non-empty string in an empty map
	searchValue = "test"
	expected = false
	actual = processor.FuzzySearch(data, searchValue, ignoreCase, keys)
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	// Test case: Searching for a non-empty string in a map with a matching key
	data = map[string]any{"key": "value"}
	searchValue = "value"
	expected = true
	actual = processor.FuzzySearch(data, searchValue, ignoreCase, keys)
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	// Test case: Searching for a non-empty string in a map with a non-matching key
	searchValue = "test"
	expected = false
	actual = processor.FuzzySearch(data, searchValue, ignoreCase, keys)
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	// Test case: Searching for a non-empty string in a map with a matching nested value
	data = map[string]any{"key": map[string]any{"nestedKey": "nestedValue"}}
	searchValue = "nestedValue"
	expected = true
	actual = processor.FuzzySearch(data, searchValue, ignoreCase, keys)
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	// Test case: Searching for a non-empty string in a map with a non-matching nested value
	searchValue = "test"
	expected = false
	actual = processor.FuzzySearch(data, searchValue, ignoreCase, keys)
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	// Test case: Searching for a non-empty string in a map with a matching key and ignore case
	data = map[string]any{"key": "value"}
	searchValue = "VALUE"
	ignoreCase = true
	expected = true
	actual = processor.FuzzySearch(data, searchValue, ignoreCase, keys)
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	// Test case: Searching for a non-empty string in a map with a non-matching key and ignore case
	searchValue = "test"
	expected = false
	actual = processor.FuzzySearch(data, searchValue, ignoreCase, keys)
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	// Test case: Searching for a non-empty string in a map with a matching nested value and ignore case
	data = map[string]any{"key": map[string]any{"nestedKey": "nestedValue"}}
	searchValue = "NESTEDVALUE"
	expected = true
	actual = processor.FuzzySearch(data, searchValue, ignoreCase, keys)
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

	// Test case: Searching for a non-empty string in a map with a non-matching nested value and ignore case
	searchValue = "test"
	expected = false
	actual = processor.FuzzySearch(data, searchValue, ignoreCase, keys)
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
	data = map[string]any{
		"name": "John Doe",
		"address": map[string]any{
			"street": "123 Main Street",
			"city":   "Springfield",
		},
		"email":   "john.doe@example.com",
		"age":     30,
		"company": "Tech Solutions",
	}
	searchValue = "spring"
	ignoreCase = true
	keys = []string{"name", "address"}
	expected = true
	actual = processor.FuzzySearch(data, searchValue, ignoreCase, keys)
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestMatchAnyRegex(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		regex    []string
		expected bool
	}{
		{"empty regex list", "hello", []string{}, false},
		{"single regex pattern match", "hello", []string{"hello"}, true},
		{"single regex pattern no match", "hello", []string{"world"}, false},
		{"multiple regex patterns match", "hello", []string{"hello", "world"}, true},
		{"multiple regex patterns no match", "hello", []string{"foo", "bar"}, false},
		{"invalid regex pattern", "hello", []string{"["}, false}, // invalid regex pattern
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := processor.MatchAnyRegex(tt.value, tt.regex)
			if actual != tt.expected {
				t.Errorf("MatchAnyRegex(%q, %v) = %v, want %v", tt.value, tt.regex, actual, tt.expected)
			}
		})
	}
}

func TestMatchAnyRegexInvalidRegex(t *testing.T) {
	regex := []string{"["}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MatchAnyRegex did not panic on invalid regex pattern")
		}
	}()
	processor.MatchAnyRegex("hello", regex)
}

func TestDeviceSearch(t *testing.T) {
	data := map[string]any{
		"deviceRole": "WanRouter",
		"description": map[string]string{
			"en": "WanRouter",
			"zh": "出口路由器",
		},
		"weight":        10,
		"abbreviation":  "WRT",
		"productFamily": "Routing",
	}

	searchValue := "WRT"
	expected := true
	actual := processor.FuzzySearch(data, searchValue, true, []string{})
	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}

}
