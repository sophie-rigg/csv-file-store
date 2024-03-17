package utils

import (
	"regexp"
)

var (
	// _emailRegex is a regex to match an email, not complete but good enough for this purpose
	_emailRegex = regexp.MustCompile(`[[:alnum:]]*@{1}[[:alnum:]]*[.][[:alnum:]]*`)
	// _idFormatRegex is a regex to match the format of a uuid
	_idFormatRegex = regexp.MustCompile(`^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`)
)

// RecordContainsEmail Checks a record to see if it contains an email
func RecordContainsEmail(record []string) bool {
	for _, value := range record {
		if containsEmail(value) {
			return true
		}
	}
	return false
}

func containsEmail(value string) bool {
	return _emailRegex.MatchString(value)
}

// CheckIDFormat checks if the given id is in the correct format of uuid
func CheckIDFormat(id string) bool {
	return _idFormatRegex.MatchString(id)
}
