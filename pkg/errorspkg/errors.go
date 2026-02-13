package errorspkg

import "fmt"

// ErrorCode – ...
type ErrorCode int8

const (
	// NotFound – ...
	NotFound ErrorCode = iota
	// Internal – ...
	Internal
	// BadRequest – ...
	BadRequest
)

// NestedError – struct to wrap error
type NestedError struct {
	Err  error
	Msg  string
	Code ErrorCode
}

// Error – implements error interface
func (e *NestedError) Error() string {
	if e.Err == nil {
		return e.Msg
	}

	return fmt.Sprintf("%s: %s", e.Msg, e.Err.Error())
}

// New – ...
func New(msg string) error {
	return &NestedError{
		Msg: msg,
	}
}

// NewC – ...
func NewC(msg string, code ErrorCode) error {
	return &NestedError{
		Msg:  msg,
		Code: code,
	}
}

// Wrap – ...
func Wrap(err error, msg string) error {
	if err == nil {
		return New(msg)
	}

	return &NestedError{
		Msg: msg,
		Err: err,
	}
}

// WrapC – ...
func WrapC(err error, msg string, code ErrorCode) error {
	if err == nil {
		return NewC(msg, code)
	}

	return &NestedError{
		Msg:  msg,
		Err:  err,
		Code: code,
	}
}

// Unwrap – ...
func (e *NestedError) Unwrap() error {
	if e.Err == nil {
		return nil
	}

	return e.Err
}
