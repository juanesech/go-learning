package lib

import (
	"strings"
)

//PassValidator function can validate if password policy is compliant
func PassValidator(r Record) bool {
	charCount := strings.Count(r.recordString, r.char)
	if charCount >= r.min && charCount <= r.max {
		return true
	}
	return false
}
