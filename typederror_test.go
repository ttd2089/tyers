package tyers

import (
	"errors"
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

	t.Run("Error()/returns the supplied text", func(t *testing.T) {
		expectedText := "instance of error"
		underTest := New(errors.New("error type"), expectedText)
		actualText := underTest.Error()
		if actualText != expectedText {
			t.Errorf("expected '%s'; got '%s'", expectedText, actualText)
		}
	})
}
