package vars

import "net/http"

const STATUS_OK_CODE = http.StatusOK
const STATUS_OK_STRING = "OK"

const STATUS_DB_ERROR_CODE = http.StatusBadRequest
const STATUS_DB_ERROR_STRING = "dbError"

const STATUS_SERVER_ERROR_CODE = http.StatusInternalServerError
const STATUS_SERVER_ERROR_STRING = "apiError"

const STATUS_NOTFOUND_ERROR_CODE = http.StatusNotFound
const STATUS_NOTFOUND_ERROR_STRING = "notFound"

const STATUS_NOTFOUND_USER_ERROR_CODE = http.StatusForbidden
const STATUS_NOTFOUND_USER_ERROR_STRING = "noUserFound"

const STATUS_REQUEST_ERROR_CODE = http.StatusBadRequest
const STATUS_REQUEST_ERROR_STRING = "requestError"

const STATUS_VALIDATION_ERROR_CODE = http.StatusBadRequest
const STATUS_VALIDATION_ERROR_STRING = "validationError"

const STATUS_AUTH_ERROR_CODE = http.StatusForbidden
const STATUS_AUTH_ERROR_STRING = "notAllowed"

const STATUS_AUTH_LOCKED_ERROR_CODE = http.StatusForbidden
const STATUS_AUTH_LOCKED_ERROR_STRING = "userLocked"

type StatusErr struct {
	statusCode int
	statusMsg  string
	err        error
}

func (r *StatusErr) Error() string {
	return r.statusMsg
}
func (r *StatusErr) Code() int {
	return r.statusCode
}
func (r *StatusErr) Unwrap() error {
	return r.err
}

func StatusDBError(err error) error {
	return &StatusErr{STATUS_DB_ERROR_CODE, STATUS_DB_ERROR_STRING, err}
}
func StatusServerError(err error) error {
	return &StatusErr{STATUS_SERVER_ERROR_CODE, STATUS_SERVER_ERROR_STRING, err}
}
func StatusNotfoundError(err error) error {
	return &StatusErr{STATUS_NOTFOUND_ERROR_CODE, STATUS_NOTFOUND_ERROR_STRING, err}
}
func StatusNotfoundUserError(err error) error {
	return &StatusErr{STATUS_NOTFOUND_USER_ERROR_CODE, STATUS_NOTFOUND_USER_ERROR_STRING, err}
}
func StatusRequestError(err error) error {
	return &StatusErr{STATUS_REQUEST_ERROR_CODE, STATUS_REQUEST_ERROR_STRING, err}
}
func StatusValidationError(err error) error {
	return &StatusErr{STATUS_VALIDATION_ERROR_CODE, STATUS_VALIDATION_ERROR_STRING, err}
}
func StatusAuthError(err error) error {
	return &StatusErr{STATUS_AUTH_ERROR_CODE, STATUS_AUTH_ERROR_STRING, err}
}
