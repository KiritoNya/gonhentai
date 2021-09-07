package raw

import (
	"errors"
	"unicode/utf8"
)

const tagName string = "raw"

// j => jpg, p => png, g => gif
func normalizeExt(ext string) (string, error) {
	switch ext {
	case "j":
		return "jpg", nil
	case "p":
		return "png", nil
	case "g":
		return "gif", nil
	default:
		return "", errors.New("Image type not found")
	}
}

// decodeUnicodeString is a function that decode all unicode rune in the string
func decodeUnicodeString(str string) (result string) {
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		result += string(r)
		str = str[size:]
	}
	return result
}
