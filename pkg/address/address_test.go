package address_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/address"
)

func TestNewAddress(t *testing.T) {
	t.Run("construct string address", func(t *testing.T) {

		t.Run("valid cases", func(t *testing.T) {
			cases := []struct {
				value string
			}{
				{"\"MYROUTER\""},
				{"\"\""},
				{"\"ABCD EFG 123\""},
				{"\"ABC'X'Y'Z;\"\" 123\""},
				{"\"ABC'X'Y'Z;\"\" 123\""},
				{"\"ABC'X'Y'Z;\"\" \"\" 123\""},
				{"\"ABC'X'Y'Z;\"\"\"\" 123\""},
			}

			for i, c := range cases {

				t.Run(fmt.Sprintf("(%v)", i), func(t *testing.T) {
					add, err := address.NewAddress(c.value)
					if err != nil {
						t.Errorf("expected nil but got %v", err)
					}

					if add == nil {
						t.Errorf("got nil address string, want %v", c.value)
					}

					if !add.CompareValue(c.value) {
						t.Errorf("address string not match, got %v, want %v", add.Value(), c.value)
					}
				})
			}
		})

		t.Run("too short", func(t *testing.T) {
			cases := []struct {
				value string
			}{
				{""},
				{"1"},
				{"\""},
			}

			for i, c := range cases {
				t.Run(fmt.Sprintf("(%v)", i), func(t *testing.T) {
					expected := address.AddressStringTooShortError{Value: c.value}
					add, err := address.NewAddress(c.value)

					if add != nil {
						t.Errorf("expected nil address, but got %v", add)
					}

					if err == nil {
						t.Errorf("got nil error, want %v", expected)
					}

					var tooShortError *address.AddressStringTooShortError
					if !errors.As(err, &tooShortError) {
						t.Errorf("got %v error, want %v", err, expected)
					}
				})
			}
		})

		t.Run("contain invalid chars", func(t *testing.T) {
			cases := []struct {
				value string
			}{
				{"\"\n\""},
				{"\"\r\""},
				{"\"\t\""},
			}

			for i, c := range cases {
				t.Run(fmt.Sprintf("(%v)", i), func(t *testing.T) {
					expected := address.AddressStringContainInvalidCharsError{Value: c.value}
					add, err := address.NewAddress(c.value)

					if add != nil {
						t.Errorf("expected nil address, but got %v", add)
					}

					if err == nil {
						t.Errorf("got nil error, want %v", expected)
					}

					var containInvalidCharsError *address.AddressStringContainInvalidCharsError
					if !errors.As(err, &containInvalidCharsError) {
						t.Errorf("got %v error, want %v", err, expected)
					}
				})
			}
		})

		t.Run("invalid use of the quotes", func(t *testing.T) {
			cases := []struct {
				value string
			}{
				{"123\""},
				{"\"123"},
				{"\"ABC'X'Y'Z;\" 123\""},
				{"\"ABC'X'Y'Z;\"\"\" 123\""},
				{"\"ABC'X'Y'Z;\" \"123\""},
				{"\"ABC'X'Y'Z;\"\"\" \"\" 123\""},
			}

			for i, c := range cases {
				t.Run(fmt.Sprintf("(%v)", i), func(t *testing.T) {
					expected := address.AddressStringQuoteError{Value: c.value}
					add, err := address.NewAddress(c.value)

					if add != nil {
						t.Errorf("expected nil address, but got %v", add)
					}

					if err == nil {
						t.Errorf("got nil error, want %v", expected)
					}

					var quoteError *address.AddressStringQuoteError
					if !errors.As(err, &quoteError) {
						t.Errorf("got %v error, want %v", err, expected)
					}
				})
			}
		})
	})
}
