package helpers

import "regexp"

const (
	updatedString = ",\"createdat\":\\s?\"(20.*?)\""
	createdString = ",\"updatedat\":\\s?\"(20.*?)\""
)

var (
	UpdatedRegex = regexp.MustCompile(updatedString)
	CreatedRegex = regexp.MustCompile(createdString)
)
