package digit_test

import (
	"testing"

	"github.com/pallat/digit"
)

func TestNotDigit(t *testing.T) {
	t.Run("-1", func(t *testing.T) {
		given := -1

		want := "-1 is not digit"

		s := digit.Pronunciation(given)
		if want != s {
			t.Errorf(s)
		}
	})
	t.Run("10", func(t *testing.T) {
		given := 10

		want := "10 is not digit"

		s := digit.Pronunciation(given)
		if want != s {
			t.Errorf(s)
		}
	})
}
