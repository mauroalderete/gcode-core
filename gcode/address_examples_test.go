package gcode_test

import (
	"fmt"
	"math"

	"github.com/mauroalderete/gcode-cli/gcode"
)

//#region examples

func ExampleIsAddressStringValid() {
	const s = "\"Hola \"\"Mundo\"\" \""

	err := gcode.IsAddressStringValid(s)
	if err != nil {
		_ = fmt.Errorf("invalid format: %v", err)
		return
	}

	fmt.Println("string has valid format")

	// Output: string has valid format
}

func ExampleAddress() {
	add, err := gcode.NewAddress[float32](12)

	if err != nil {
		_ = fmt.Errorf("unexpected error: %v", err)
		return
	}

	var value float32 = add.Value()
	fmt.Printf("address value is: %v\n", value)

	// Output: address value is: 12
}

func ExampleAddress_second() {
	add, err := gcode.NewAddress[float32](math.Pi)

	if err != nil {
		_ = fmt.Errorf("unexpected error: %v", err)
		return
	}

	add.SetValue(3.1415)

	fmt.Printf("address value is: %v\n", add.String())

	// Output: address value is: 3.1415
}

func ExampleNewAddress() {
	add, err := gcode.NewAddress[int32](-23)
	if err != nil {
		_ = fmt.Errorf("unexpected error: %v", err)
		return
	}

	fmt.Printf("address value is: %s\n", add)

	// Output: address value is: -23
}

func ExampleNewAddress_second() {
	add, err := gcode.NewAddress[float32](math.Pi)
	if err != nil {
		_ = fmt.Errorf("unexpected error: %v", err)
		return
	}

	fmt.Printf("address value is: %s\n", add)

	// Output: address value is: 3.1415927
}

func ExampleNewAddress_third() {
	add, err := gcode.NewAddress("\"Hola Mundo!\"")
	if err != nil {
		_ = fmt.Errorf("unexpected error: %v", err)
		return
	}

	fmt.Printf("address value is: %s\n", add)

	// Output: address value is: "Hola Mundo!"
}

func ExampleAddress_Compare() {
	addBase, err := gcode.NewAddress[float32](math.Pi)
	if err != nil {
		_ = fmt.Errorf("unexpected error: %v", err)
		return
	}

	addTarget, err := gcode.NewAddress[float32](3.1415)
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
	add, err := gcode.NewAddress("\"Hola Mundo!\"")

	if err != nil {
		_ = fmt.Errorf("unexpected error: %v", err)
		return
	}

	fmt.Printf("address value is: %s\n", add.String())

	// Output: address value is: "Hola Mundo!"
}

func ExampleAddress_String_second() {
	add, err := gcode.NewAddress[float32](12)

	if err != nil {
		_ = fmt.Errorf("unexpected error: %v", err)
		return
	}

	fmt.Printf("address value is: %s\n", add.String())

	// Output: address value is: 12.0
}

func ExampleAddress_Value() {
	add, err := gcode.NewAddress[float32](0)

	if err != nil {
		_ = fmt.Errorf("unexpected error: %v", err)
		return
	}

	fmt.Printf("address value is: %v\n", add.Value())

	// Output: address value is: 0
}

func ExampleAddress_SetValue() {
	add, err := gcode.NewAddress[float32](0)

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
