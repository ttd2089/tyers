package tyers

import (
	"errors"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {

	t.Run("errors.Is()/returns true for error type", func(t *testing.T) {
		errorType := errors.New("error type")
		underTest := New(errorType, "instance of error")
		if !errors.Is(underTest, errorType) {
			t.Errorf("expected '%s' to be '%s'", underTest, errorType)
		}
	})

	t.Run("Error()/formats as the supplied text", func(t *testing.T) {
		expectedText := "instance of error"
		underTest := New(errors.New("error type"), expectedText)
		actualText := underTest.Error()
		if actualText != expectedText {
			t.Errorf("expected '%s'; got '%s'", expectedText, actualText)
		}
	})
}

func TestErrorf(t *testing.T) {

	t.Run("errors.Is()/returns true for error type", func(t *testing.T) {
		errorType := errors.New("error type")
		underTest := Errorf(errorType, "instance of %s", "error")
		if !errors.Is(underTest, errorType) {
			t.Errorf("expected '%s' to be '%s'", underTest, errorType)
		}
	})

	t.Run("errors.Is()/returns true for a wrapped error", func(t *testing.T) {
		wrappedError := errors.New("wrapped error")
		errorType := errors.New("error type")
		underTest := Errorf(errorType, "error: %w", wrappedError)
		if !errors.Is(underTest, wrappedError) {
			t.Errorf("expected '%s' to be '%s'", underTest, wrappedError)
		}
	})

	t.Run("errors.Is()/returns true for a nested wrapped error", func(t *testing.T) {
		wrappedError := errors.New("wrapped error")
		errorType := errors.New("error type")
		underTest := Errorf(errorType, "error: %w", fmt.Errorf("nested error: %w", wrappedError))
		if !errors.Is(underTest, wrappedError) {
			t.Errorf("expected '%s' to be '%s'", underTest, wrappedError)
		}
	})

	t.Run("Error()/formats according to fmt.Errorf", func(t *testing.T) {
		formatSpecifier := "%s %d %+v"
		operands := []any{"string", 10, map[string]bool{"true": true}}
		expectedText := fmt.Errorf(formatSpecifier, operands...).Error()
		underTest := Errorf(errors.New("error type"), formatSpecifier, operands...)
		actualText := underTest.Error()
		if actualText != expectedText {
			t.Errorf("expected '%s'; got '%s'", expectedText, actualText)
		}
	})
}
