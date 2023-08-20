package appraisal

import (
	"regexp"
)

type CodeType uint8

const (
	UnknownCode CodeType = 0
	BuybackCode CodeType = 1
	ShopCode    CodeType = 2

	ReStr string = "[us]{1}[0-9a-f]{16}"
	// ReStr string = "[uUsS]{1}[0-9a-fA-F]{16}"
)

var Re *regexp.Regexp = regexp.MustCompile(ReStr)

func ParseCode(txt string) (string, CodeType) {
	code := Re.FindString(txt)
	if code == "" {
		return "", UnknownCode
	} else if code[0] == 'u' {
		return code, BuybackCode
	} else {
		return code, ShopCode
	}
}

// func isLowercase(code string) bool {
// 	if code[0] == 'U' || code[0] == 'S' {
// 		return false
// 	}
// 	for i := 1; i < len(code); i++ {
// 		if code[i] >= 'A' && code[i] <= 'F' {
// 			return false
// 		}
// 	}
// 	return true
// }