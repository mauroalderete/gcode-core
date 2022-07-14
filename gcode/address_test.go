package gcode

import (
	"fmt"
	"math"
	"testing"
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
					add, err := NewAddress(c.value)
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
					add, err := NewAddress(c.value)

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
				add, err := NewAddress(c)
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
				add, err := NewAddress(c)
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
					add, err := NewAddress("\"something\"")
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
					add, err := NewAddress("\"something\"")
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
				add, err := NewAddress[int32](0)
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
				add, err := NewAddress[float32](0)
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
