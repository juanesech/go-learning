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

//PassValidatorV2 function can validate if password policy is compliant
func PassValidatorV2(r Record) bool {
	r.min--
	r.max--
	if string(r.recordString[r.min]) == r.char && string(r.recordString[r.max]) == r.char {
		return false
	} else if string(r.recordString[r.min]) != r.char && string(r.recordString[r.max]) != r.char {
		return false
	} else if string(r.recordString[r.min]) != r.char || string(r.recordString[r.max]) != r.char {
		return true
	}
	return false
}
