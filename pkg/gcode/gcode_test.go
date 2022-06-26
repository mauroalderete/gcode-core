package gcode_test

import (
	"fmt"

	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/address"
	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/gcode"
	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/word"
)

func ExampleNewGcode() {

	gc, err := gcode.NewGcode("X", "12.3")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%s", gc)

	// Output: X12.3
}

func ExampleGcode_String() {

	gc, err := gcode.NewGcode("X", "12.3")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	fmt.Println(gc)

	// Output: X12.3
}

func ExampleGcode_Word() {
	gc, err := gcode.NewGcode("D", "0")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	var w word.Word = gc.Word()

	fmt.Println(w.String())

	// Output: D
}

func ExampleGcode_Address() {
	gc, err := gcode.NewGcode("N", "66555")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	var a address.Address = gc.Address()

	fmt.Println(a.String())

	// Output: 66555
}

func ExampleGcode_Compare() {
	gc1, err := gcode.NewGcode("G", "33")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	gc2, err := gcode.NewGcode("G", "34")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	fmt.Printf("%t", gc1.Compare(gc2))

	// Output: false
}
