package errors

import (
	"fmt"
	"gin-starter/common/constant"
	"net/http"
	"strconv"
	"strings"
)

var (
	// ErrRecordNotFound represents error when record is not found.
	ErrRecordNotFound = NewError(http.StatusNotFound, "record not found")
	// ErrInternalServerError represents error when internal server error occurs.
	ErrInternalServerError = NewError(http.StatusInternalServerError, "internal server error")
	// ErrWrongLoginCredentials represents error when login credentials are wrong.
	ErrWrongLoginCredentials = NewError(http.StatusBadRequest, "username atau password salah")
	// ErrTooFarFromOffice represents error when user is too far from office.
	ErrTooFarFromOffice = NewError(http.StatusBadRequest, "anda terlalu jauh dari kantor")
	// ErrInvalidArgument represents error when invalid argument is passed.
	ErrInvalidArgument = NewError(http.StatusBadRequest, "invalid argument")
	// ErrWrongPasswordConfirmation define error if password confirmation is wrong
	ErrWrongPasswordConfirmation = NewError(http.StatusBadRequest, "Konfirmasi password kamu tidak sesuai.")
	// ErrOTPMismatch represents error when otp is mismatched.
	ErrOTPMismatch    = NewError(http.StatusBadRequest, "Kode OTP Salah")
	ErrDuplicateEntry = NewError(http.StatusBadRequest, "Data sudah ada")
	ErrEmptyData      = NewError(http.StatusBadRequest, "Data mandatory kosong")
)

// Error represents a data structure for error.
// It implements golang error interface.
type Error struct {
	// Code represents error code.
	Code int `json:"code"`
	// Message represents error message.
	// This is the message that exposed to the user.
	Message string `json:"message"`
}

// NewError creates an instance of Error.
func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// Error returns internal message in one string.
func (err *Error) Error() error {
	return fmt.Errorf("%d:%s", err.Code, err.Message)
}

// ParseError parses error message and returns an instance of Error.
func ParseError(err error) *Error {
	if err == nil {
		return nil
	}

	split := strings.Split(err.Error(), ":")

	fmt.Println(err)

	code, err := strconv.ParseInt(split[0], constant.Ten, constant.ThirtyTwo)

	if err != nil {
		return ErrInternalServerError
	}

	return NewError(int(code), split[1])
}
