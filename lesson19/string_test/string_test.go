package string_test

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	s := "abcd1234"
	prefix := hasPrefix(s, "ab")
	fmt.Println(prefix)
	suffix := hasSuffix(s, "1234")
	fmt.Println(suffix)
}

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func hasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s) - len(suffix):] == suffix
}
// 判断是否为子串
func Contains(s, substr string) bool {
	for i:=0;i<len(s);i++ {
		if hasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}