// Package tyers provides error wrapping functions to faciltate creating typed errors that can be
// checked using errors.Is. For example, you might declare some error ErrNoSuchFoo then create
// specific instances of it using it as a type rather than wrapping it:
//
//	var ErrNoSuchFoo error = errors.New("no such foo")
//	err := tyers.Errorf(ErrNoSuchFoo, "'%s' is not a foo", fooName)
//	errors.Is(err, ErrNoSuchFoo) // returns true
//
// Using types instead of wrapping errors decouples the message from the ability to use errors.Is
// without requiring you to write your own Unwrap and Is functions.
package tyers

import (
	"errors"
	"fmt"
)

// New returns an error that formats as the given text and causes errors.Is to return true when
// called with errorType as the target.
func New(errorType error, text string) error {
	return &typedError{
		errorType:  errorType,
		errorValue: errors.New(text),
	}
}

// Errorf returns an error that formats according fmt.Errorf applied to the given text and operands
// and causes errors.Is to return true when called with errorType or any wrapped errors as the
// target.
func Errorf(errorType error, text string, a ...any) error {
	return &typedError{
		errorType:  errorType,
		errorValue: fmt.Errorf(text, a...),
	}
}

// A typedError is an error that binds a normal error value with an additional error as its type
// such that using errors.Is can match the value and any errors it wraps as well as a specific
// error used to represent a category of error which may not appear in the wrapped values.
type typedError struct {

	// errorType is an error representing a category to which the TypedError belongs.
	errorType error

	// errorValue contains the value, ie. the message and wrapped errors of a TypedError.
	errorValue error
}

// Error forwards the call to the Error method of the underlying error value.
func (t *typedError) Error() string {
	return t.errorValue.Error()
}

// Unwrap returns the underlying error value.
func (t *typedError) Unwrap() error {
	return t.errorValue
}

// Is returns a bool indicating whether the ErrorType of the TypedError is the target error.
func (t *typedError) Is(target error) bool {
	return errors.Is(t.errorType, target)
}
