package helpers

import (
	"errors"
	"go-api-starter-kit/utils/vars"
)

func GetID(claimsID interface{}) (int64, error) {
	id, ok := claimsID.(int64)
	if !ok {
		return 0, vars.StatusRequestError(errors.New("bad claims"))
	}

	return id, nil
}
