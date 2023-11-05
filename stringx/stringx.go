package stringx

import "unicode"

func IsEmpty(s string) bool {
	return len(s) == 0
}

func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

func IsBlank(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// IsNumeric Checks if the CharSequence contains only Unicode digits. A decimal point is not a Unicode digit and returns false.
// null will return false. An empty CharSequence (len=0) will return false.
// Note that the method does not allow for a leading sign, either positive or negative
func IsNumeric(s string) bool {
	if IsEmpty(s) {
		return false
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
