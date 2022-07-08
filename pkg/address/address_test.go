package address_test

import (
	"errors"
	"fmt"
	"math"
	"testing"

	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/address"
)

func ExampleIsAddressStringValid() {
	const s = "\"Hola \"\"Mundo\"\" \""

	err := address.IsAddressStringValid(s)
	if err != nil {
		var containInvalidChars *address.AddressStringContainInvalidCharsError
		if !errors.As(err, &containInvalidChars) {
			_ = fmt.Errorf("invalid format: %v", containInvalidChars.Error())
			return
		}

		var quoteError *address.AddressStringQuoteError
		if !errors.As(err, &quoteError) {
			_ = fmt.Errorf("invalid format: %v", quoteError.Error())
			return
		}

		var tooShort *address.AddressStringTooShortError
		if !errors.As(err, &tooShort) {
			_ = fmt.Errorf("invalid format: %v", tooShort.Error())
			return
		}
	}

	fmt.Println("string has valid format")

	// Output: string has valid format
}

func ExampleAddress() {
	add, err := address.NewAddress[float32](12)

	if err != nil {
		_ = fmt.Errorf("unexpected error: %v", err)
		return
	}

	var value float32 = add.Value
	fmt.Printf("address value is: %v\n", value)

	// Output: address value is: 12
}

func ExampleAddress_second() {
	add, err := address.NewAddress[float32](math.Pi)

	if err != nil {
		_ = fmt.Errorf("unexpected error: %v", err)
		return
	}

	add.Value = 3.1415

	fmt.Printf("address value is: %v\n", add.String())

	// Output: address value is: 3.1415
}

func ExampleNewAddress() {
	add, err := address.NewAddress[int32](-23)
	if err != nil {
		_ = fmt.Errorf("unexpected error: %v", err)
		return
	}

	fmt.Printf("address value is: %s\n", add)

	// Output: address value is: -23
}

func ExampleNewAddress_second() {
	add, err := address.NewAddress[float32](math.Pi)
	if err != nil {
		_ = fmt.Errorf("unexpected error: %v", err)
		return
	}

	fmt.Printf("address value is: %s\n", add)

	// Output: address value is: 3.1415927
}

func ExampleNewAddress_third() {
	add, err := address.NewAddress("\"Hola Mundo!\"")
	if err != nil {
		_ = fmt.Errorf("unexpected error: %v", err)
		return
	}

	fmt.Printf("address value is: %s\n", add)

	// Output: address value is: "Hola Mundo!"
}

func ExampleAddress_Compare() {
	addBase, err := address.NewAddress[float32](math.Pi)
	if err != nil {
		_ = fmt.Errorf("unexpected error: %v", err)
		return
	}

	addTarget, err := address.NewAddress[float32](3.1415)
	if err != nil {
		_ = fmt.Errorf("unexpected error: %v", err)
		return
	}

	if addBase.Compare(*addTarget) {
		fmt.Println("both addresses are equals")
	} else {
		fmt.Println("both addresses are diferents")
	}

	// Output: both addresses are diferents
}

func ExampleAddress_String() {
	add, err := address.NewAddress("\"Hola Mundo!\"")

	if err != nil {
		_ = fmt.Errorf("unexpected error: %v", err)
		return
	}

	fmt.Printf("address value is: %s\n", add.String())

	// Output: address value is: "Hola Mundo!"
}

func ExampleAddress_String_second() {
	add, err := address.NewAddress[float32](12)

	if err != nil {
		_ = fmt.Errorf("unexpected error: %v", err)
		return
	}

	fmt.Printf("address value is: %s\n", add.String())

	// Output: address value is: 12.0
}

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
						return
					}

					if add.Value != c.value {
						t.Errorf("address string not match, got %v, want %v", add.Value, c.value)
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

	t.Run("constructor address integer", func(t *testing.T) {

		cases := [5]int32{-111, -1, 0, 1, 111}

		for i, c := range cases {
			t.Run(fmt.Sprintf("(%v)", i), func(t *testing.T) {
				add, err := address.NewAddress(c)
				if err != nil {
					t.Errorf("expected nil error but got %v", err)
				}

				if add == nil {
					t.Errorf("got nil address, want %v", c)
					return
				}

				if add.Value != c {
					t.Errorf("got %v address, want %v", add.Value, c)
				}
			})
		}
	})

	t.Run("constructor address float", func(t *testing.T) {

		cases := [10]float32{-111, -1, 0, 1, 111, 7.5, -0.0002, math.Pi, math.E, math.MaxFloat32}

		for i, c := range cases {
			t.Run(fmt.Sprintf("(%v)", i), func(t *testing.T) {
				add, err := address.NewAddress(c)
				if err != nil {
					t.Errorf("expected nil error but got %v", err)
				}

				if add == nil {
					t.Errorf("got nil address, want %v", c)
					return
				}

				if add.Value != c {
					t.Errorf("got %v address, want %v", add.Value, c)
				}
			})
		}
	})
}

func TestAddress_AddressStringContainInvalidCharsError(t *testing.T) {
	t.Run("valid message", func(t *testing.T) {
		const value = "\"\n\""
		const expected = "gcode's address string contain invalid chars: " + value

		_, err := address.NewAddress(value)
		if err == nil {
			t.Errorf("expected nil error but got %v", err)
			return
		}

		if err.Error() != expected {
			t.Errorf("got %v, want %v", err, expected)
		}
	})
}

func TestAddress_AddressStringQuoteError(t *testing.T) {
	t.Run("valid message", func(t *testing.T) {
		const value = "some"
		const expected = "gcode's address string has an invalid use of the quotes: "

		_, err := address.NewAddress(value)
		if err == nil {
			t.Errorf("expected nil error but got %v", err)
			return
		}

		if err.Error() != expected {
			t.Errorf("got %v, want %v", err, expected)
		}
	})
}

func TestAddress_AddressStringTooShortError(t *testing.T) {
	t.Run("valid message", func(t *testing.T) {
		const value = ""
		const expected = "gcode's address string is too short: " + value

		_, err := address.NewAddress(value)
		if err == nil {
			t.Errorf("expected nil error but got %v", err)
			return
		}

		if err.Error() != expected {
			t.Errorf("got %v, want %v", err, expected)
		}
	})
}
