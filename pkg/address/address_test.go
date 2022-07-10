package address_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/address"
)

//#region unit tests

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
						t.Errorf("got error: %v, want error: not nil", err)
					}

					if add == nil {
						t.Errorf("got nil address, want string: %v", c.value)
						return
					}

					if add.Value() != c.value {
						t.Errorf("got string: %v, want string: %v", add.Value(), c.value)
					}
				})
			}
		})

		t.Run("invalid cases", func(t *testing.T) {
			cases := []struct {
				value string
			}{
				{""},
				{"1"},
				{"\""},
				{"\"\n\""},
				{"\"\r\""},
				{"\"\t\""},
				{"123\""},
				{"\"123"},
				{"\"ABC'X'Y'Z;\" 123\""},
				{"\"ABC'X'Y'Z;\"\"\" 123\""},
				{"\"ABC'X'Y'Z;\" \"123\""},
				{"\"ABC'X'Y'Z;\"\"\" \"\" 123\""},
			}

			for i, c := range cases {
				t.Run(fmt.Sprintf("(%v)", i), func(t *testing.T) {
					add, err := address.NewAddress(c.value)

					if add != nil {
						t.Errorf("got address: %v, want nil address", add)
					}

					if err == nil {
						t.Errorf("got error: nil, want error: not nil")
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
					t.Errorf("got error: %v, want error: nil", err)
				}

				if add == nil {
					t.Errorf("got address: nil, want address value: %v", c)
					return
				}

				if add.Value() != c {
					t.Errorf("got address value: %v, want address value: %v", add.Value(), c)
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
					t.Errorf("got error: %v, want error: nil", err)
				}

				if add == nil {
					t.Errorf("got address: nil, want address value: %v", c)
					return
				}

				if add.Value() != c {
					t.Errorf("got address value: %v, want address value: %v", add.Value(), c)
				}
			})
		}
	})
}

func TestAddress_SetValue(t *testing.T) {
	t.Run("set value at string address", func(t *testing.T) {

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
					add, err := address.NewAddress("\"something\"")
					if err != nil {
						t.Errorf("got error: %v, want error: not nil", err)
					}
					if add == nil {
						t.Errorf("got nil address, want string: %v", c.value)
						return
					}

					err = add.SetValue(c.value)
					if err != nil {
						t.Errorf("got error: %v, want error: not nil", err)
					}
					if add.Value() != c.value {
						t.Errorf("got string: %v, want string: %v", add.Value(), c.value)
					}
				})
			}
		})

		t.Run("invalid cases", func(t *testing.T) {
			cases := []struct {
				value string
			}{
				{""},
				{"1"},
				{"\""},
				{"\"\n\""},
				{"\"\r\""},
				{"\"\t\""},
				{"123\""},
				{"\"123"},
				{"\"ABC'X'Y'Z;\" 123\""},
				{"\"ABC'X'Y'Z;\"\"\" 123\""},
				{"\"ABC'X'Y'Z;\" \"123\""},
				{"\"ABC'X'Y'Z;\"\"\" \"\" 123\""},
			}

			for i, c := range cases {
				t.Run(fmt.Sprintf("(%v)", i), func(t *testing.T) {
					add, err := address.NewAddress("\"something\"")
					if err != nil {
						t.Errorf("got error: %v, want error: not nil", err)
					}
					if add == nil {
						t.Errorf("got nil address, want string: %v", c.value)
						return
					}

					err = add.SetValue(c.value)
					if err == nil {
						t.Errorf("got error: nil, want error: not nil")
					}
				})
			}
		})
	})

	t.Run("set value at integer address", func(t *testing.T) {

		cases := [5]int32{-111, -1, 0, 1, 111}

		for i, c := range cases {
			t.Run(fmt.Sprintf("(%v)", i), func(t *testing.T) {
				add, err := address.NewAddress[int32](0)
				if err != nil {
					t.Errorf("got error: %v, want error: nil", err)
				}
				if add == nil {
					t.Errorf("got address: nil, want address value: %v", c)
					return
				}

				err = add.SetValue(c)
				if err != nil {
					t.Errorf("got error: %v, want error: nil", err)
				}
				if add.Value() != c {
					t.Errorf("got address value: %v, want address value: %v", add.Value(), c)
				}
			})
		}
	})

	t.Run("set value at float32 address", func(t *testing.T) {

		cases := [10]float32{-111, -1, 0, 1, 111, 7.5, -0.0002, math.Pi, math.E, math.MaxFloat32}

		for i, c := range cases {
			t.Run(fmt.Sprintf("(%v)", i), func(t *testing.T) {
				add, err := address.NewAddress[float32](0)
				if err != nil {
					t.Errorf("got error: %v, want error: nil", err)
				}
				if add == nil {
					t.Errorf("got address: nil, want address value: %v", c)
					return
				}

				err = add.SetValue(c)
				if err != nil {
					t.Errorf("got error: %v, want error: nil", err)
				}
				if add.Value() != c {
					t.Errorf("got address value: %v, want address value: %v", add.Value(), c)
				}
			})
		}
	})
}

//#endregion
//#region examples

func ExampleIsAddressStringValid() {
	const s = "\"Hola \"\"Mundo\"\" \""

	err := address.IsAddressStringValid(s)
	if err != nil {
		_ = fmt.Errorf("invalid format: %v", err)
		return
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

	var value float32 = add.Value()
	fmt.Printf("address value is: %v\n", value)

	// Output: address value is: 12
}

func ExampleAddress_second() {
	add, err := address.NewAddress[float32](math.Pi)

	if err != nil {
		_ = fmt.Errorf("unexpected error: %v", err)
		return
	}

	add.SetValue(3.1415)

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

func ExampleAddress_Value() {
	add, err := address.NewAddress[float32](0)

	if err != nil {
		_ = fmt.Errorf("unexpected error: %v", err)
		return
	}

	fmt.Printf("address value is: %v\n", add.Value())

	// Output: address value is: 0
}

func ExampleAddress_SetValue() {
	add, err := address.NewAddress[float32](0)

	if err != nil {
		_ = fmt.Errorf("unexpected error: %v", err)
		return
	}

	err = add.SetValue(12)
	if err != nil {
		_ = fmt.Errorf("unexpected error: %v", err)
		return
	}

	fmt.Printf("address value is: %v\n", add.Value())

	// Output: address value is: 12
}

//#endregion
