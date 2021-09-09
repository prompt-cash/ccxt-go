package utils

import (
	"strconv"
	"strings"
)

func ParseInt(number string) int64 {
	value, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		return 0
	}
	return value
}

func ParseUint(number string) uint64 {
	value, err := strconv.ParseUint(number, 10, 64)
	if err != nil {
		return 0
	}
	return value
}

func ParseFloat(number string) float64 {
	value, err := strconv.ParseFloat(number, 64)
	if err != nil {
		return 0.0
	}
	return value
}

func ParseBool(booleanExpr string) bool {
	exprLower := strings.TrimSpace(strings.ToLower(booleanExpr))
	if exprLower == "on" { // for HTML form values
		return true
	}
	value, err := strconv.ParseBool(booleanExpr)
	if err != nil {
		return false
	}
	return value
}

// Returns the substring of str with maxLen (starting from beginning).
func MaxString(str string, maxLen int) string {
	if len(str) > maxLen {
		return str[:maxLen]
	}
	return str
}

func IsStringInList(str string, list []string) bool {
	for _, cur := range list {
		if str == cur {
			return true
		}
	}
	return false
}

func StrRepeat(s string, count int) string {
	if count < 0 {
		return ""
	}
	b := make([]byte, len(s)*count)
	bp := copy(b, s)
	for bp < len(b) {
		copy(b[bp:], b[:bp])
		bp *= 2
	}
	return string(b)
}
