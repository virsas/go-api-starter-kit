package helpers

import (
	"regexp"

	"github.com/go-playground/validator"
)

const (
	AlphaNumSpaceString = "^[a-zA-Z0-9\\s]+$"
	AlphaSpaceString    = "^[a-zA-Z\\s]+$"
)

var (
	AlphaNumSpaceRegex = regexp.MustCompile(AlphaNumSpaceString)
	AlphaSpaceRegex    = regexp.MustCompile(AlphaSpaceString)
)

func AlphaNumSpaceValid(fl validator.FieldLevel) bool {
	return AlphaNumSpaceRegex.MatchString(fl.Field().String())
}
func AlphaSpaceValid(fl validator.FieldLevel) bool {
	return AlphaSpaceRegex.MatchString(fl.Field().String())
}
