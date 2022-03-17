package config

import "net/http"

const OK_STATUS = http.StatusOK
const OK_STRING = "OK"

const DB_ERROR = http.StatusBadRequest
const DB_STRING = "dbError"

const SERVER_ERROR = http.StatusInternalServerError
const SERVER_STRING = "apiError"

const NOTFOUND_ERROR = http.StatusNotFound
const NOTFOUND_STRING = "notFound"

const REQUEST_ERROR = http.StatusBadRequest
const REQUEST_STRING = "requestError"

const VALID_ERROR = http.StatusBadRequest
const VALID_STRING = "validationError"

const AUTH_ERROR = http.StatusForbidden
const AUTH_STRING = "notAllowed"

type CustErr struct {
	statusCode int
	statusMsg  string
	err        error
}

func (r *CustErr) Error() string {
	return r.statusMsg
}
func (r *CustErr) Code() int {
	return r.statusCode
}
func (r *CustErr) Unwrap() error {
	return r.err
}

func NewDBError(err error) error {
	return &CustErr{DB_ERROR, DB_STRING, err}
}
func NewServerError(err error) error {
	return &CustErr{SERVER_ERROR, SERVER_STRING, err}
}
func NewNotfoundError(err error) error {
	return &CustErr{NOTFOUND_ERROR, NOTFOUND_STRING, err}
}
func NewRequestError(err error) error {
	return &CustErr{REQUEST_ERROR, REQUEST_STRING, err}
}
func NewValidationError(err error) error {
	return &CustErr{VALID_ERROR, VALID_STRING, err}
}
func NewAuthError(err error) error {
	return &CustErr{AUTH_ERROR, AUTH_STRING, err}
}
