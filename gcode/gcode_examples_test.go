package gcode_test

import (
	"fmt"

	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/gcode"
)

func ExampleNewGcode() {

	gc, err := gcode.NewGcode('X')
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%s", gc)

	// Output: X
}

func ExampleNewGcodeAddressable() {

	gc, err := gcode.NewGcodeAddressable[float32]('X', 12.3)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%s", gc)

	// Output: X12.3
}

func ExampleGcode_HasAddress() {

	gc, err := gcode.NewGcode('X')
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	fmt.Printf("%v\n", gc.HasAddress())

	// Output: false
}

func ExampleGcode_HasAddress_second() {

	gca, err := gcode.NewGcodeAddressable('X', "\"Hola mundo!\"")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	fmt.Printf("%v\n", gca.HasAddress())

	// Output: true
}

func ExampleGcode_HasAddress_third() {

	gca, err := gcode.NewGcodeAddressable('X', "\"Hola mundo!\"")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	var gc gcode.Gcoder = gca

	fmt.Printf("%v\n", gc.HasAddress())

	// Output: true
}

func ExampleGcode_String() {

	gc, err := gcode.NewGcode('X')
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	fmt.Println(gc)

	// Output: X
}

func ExampleGcode_String_second() {

	gca, err := gcode.NewGcodeAddressable('X', "\"Hola mundo!\"")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	fmt.Println(gca)

	// Output: X"Hola mundo!"
}

func ExampleGcode_String_third() {

	gca, err := gcode.NewGcodeAddressable('X', "\"Hola mundo!\"")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	var gc gcode.Gcoder = gca

	fmt.Println(gc)

	// Output: X"Hola mundo!"
}

func ExampleGcode_Word() {
	gc, err := gcode.NewGcode('D')
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	fmt.Println(gc.Word().String())

	// Output: D
}

func ExampleGcode_Word_second() {
	gca, err := gcode.NewGcodeAddressable[int32]('D', 0)
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	fmt.Println(gca.Word().String())

	// Output: D
}

func ExampleGcodeAddressable_Address() {
	gc, err := gcode.NewGcodeAddressable[int32]('N', 66555)
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	fmt.Println(gc.Address().String())

	// Output: 66555
}

func ExampleGcodeAddressable_Address_second() {
	gca, err := gcode.NewGcodeAddressable[int32]('N', 66555)
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	var gc gcode.Gcoder = gca

	if !gc.HasAddress() {
		fmt.Println("Ups! gcode not contain address")
		return
	}

	if value, ok := gc.(*gcode.GcodeAddressable[int32]); ok {
		add := value.Address()
		fmt.Printf("the int32 address recovered is %v\n", add.String())
	}

	// Output: the int32 address recovered is 66555
}
