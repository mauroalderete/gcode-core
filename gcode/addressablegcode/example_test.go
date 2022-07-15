package addressablegcode_test

import (
	"fmt"

	"github.com/mauroalderete/gcode-cli/gcode"
	"github.com/mauroalderete/gcode-cli/gcode/addressablegcode"
)

func ExampleNew() {

	gc, err := addressablegcode.New[float32]('X', 12.3)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%s", gc)

	// Output: X12.3
}

func ExampleGcode_HasAddress_second() {

	gca, err := addressablegcode.New('X', "\"Hola mundo!\"")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	fmt.Printf("%v\n", gca.HasAddress())

	// Output: true
}

func ExampleGcode_HasAddress_third() {

	gca, err := addressablegcode.New('X', "\"Hola mundo!\"")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	var gc gcode.Gcoder = gca

	fmt.Printf("%v\n", gc.HasAddress())

	// Output: true
}

func ExampleGcode_String_second() {

	gca, err := addressablegcode.New('X', "\"Hola mundo!\"")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	fmt.Println(gca)

	// Output: X"Hola mundo!"
}

func ExampleGcode_String_third() {

	gca, err := addressablegcode.New('X', "\"Hola mundo!\"")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	var gc gcode.Gcoder = gca

	fmt.Println(gc)

	// Output: X"Hola mundo!"
}

func ExampleGcode_Word_second() {
	gca, err := addressablegcode.New[int32]('D', 0)
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	fmt.Println(gca.Word())

	// Output: 68
}

func ExampleGcode_Address() {
	gc, err := addressablegcode.New[int32]('N', 66555)
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	fmt.Println(gc.Address())

	// Output: 66555
}

func ExampleGcode_Address_second() {
	gca, err := addressablegcode.New[int32]('N', 66555)
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	var gc gcode.Gcoder = gca

	if !gc.HasAddress() {
		fmt.Println("Ups! gcode not contain address")
		return
	}

	if value, ok := gc.(gcode.AddresableGcoder[int32]); ok {
		add := value.Address()
		fmt.Printf("the int32 address recovered is %v\n", add)
	}

	// Output: the int32 address recovered is 66555
}
