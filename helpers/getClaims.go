package helpers

import (
	"errors"
	"go-api-starter-kit/utils/vars"
)

func GetClaimsInt64(claimsID interface{}) (int64, error) {
	id, ok := claimsID.(int64)
	if !ok {
		return 0, vars.StatusRequestError(errors.New("bad claims"))
	}

	return id, nil
}

func GetClaimsBool(claimsID interface{}) (bool, error) {
	status, ok := claimsID.(bool)
	if !ok {
		return false, vars.StatusRequestError(errors.New("bad claims"))
	}

	return status, nil
}
