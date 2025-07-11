package validation

import "regexp"

func IsEmail(s string) bool {
	// Simple email regex pattern
	re := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	return re.MatchString(s)
}

func IsPassLength(s string, expectedLen int) bool {
	// Simple email regex pattern
	strLength := len(s)

	if strLength >= expectedLen {
		return true
	} else {
		return false
	}
}
