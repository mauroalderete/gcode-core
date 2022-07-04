package address_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/address"
)

func TestNewAddress(t *testing.T) {
	t.Run("construct string address", func(t *testing.T) {
		cases := []struct {
			value         string
			errorExpected error
		}{
			{"\"MYROUTER\"", nil},
			{"", &address.AddressStringTooShortError{Value: ""}},
			{"1", &address.AddressStringTooShortError{Value: "1"}},
			{"\"", &address.AddressStringTooShortError{Value: "\""}},
			{"\"\n\"", &address.AddressStringContainInvalidCharsError{Value: "\"\n\""}},
			{"\"\r\"", &address.AddressStringContainInvalidCharsError{Value: "\"\t\""}},
			{"\"\t\"", &address.AddressStringContainInvalidCharsError{Value: "\"\r\""}},
			{"\"\"", nil},
			{"\"ABCD EFG 123\"", nil},
			{"\"ABC'X'Y'Z;\"\" 123\"", nil},
			{"123\"", &address.AddressStringQuoteError{Value: "123\""}},
			{"\"123", &address.AddressStringQuoteError{Value: "\"123"}},
			{"\"ABC'X'Y'Z;\" 123\"", &address.AddressStringQuoteError{Value: "\"ABC'X'Y'Z;\" 123\""}},
			{"\"ABC'X'Y'Z;\"\" 123\"", nil},
			{"\"ABC'X'Y'Z;\"\"\" 123\"", &address.AddressStringQuoteError{Value: "\"ABC'X'Y'Z;\"\"\" 123\""}},
			{"\"ABC'X'Y'Z;\" \"123\"", &address.AddressStringQuoteError{Value: "\"ABC'X'Y'Z;\" \"123\""}},
			{"\"ABC'X'Y'Z;\"\"\" \"\" 123\"", &address.AddressStringQuoteError{Value: "\"ABC'X'Y'Z;\"\"\" \"\" 123\""}},
			{"\"ABC'X'Y'Z;\"\" \"\" 123\"", nil},
			{"\"ABC'X'Y'Z;\"\"\"\" 123\"", nil},
		}

		for i, c := range cases {

			t.Run(fmt.Sprintf("case (%v)", i), func(t *testing.T) {
				add, err := address.NewAddress(c.value)
				if err != nil && c.errorExpected == nil {
					t.Errorf("expected nil but got %v", err)
				}

				if !errors.Is(err, c.errorExpected) {
					t.Errorf("error type not match, got %v, want %v", err, c.errorExpected)
				}

				if !add.CompareValue(c.value) {
					t.Errorf("string address not match, got %v, want %v", add.Value(), c.value)
				}
			})
		}
	})
}
