package addressablegcode_test

import (
	"fmt"

	"github.com/mauroalderete/gcode-core/gcode"
	"github.com/mauroalderete/gcode-core/gcode/addressablegcode"
	"github.com/mauroalderete/gcode-core/gcode/unaddressablegcode"
)

func ExampleNew() {

	gc, err := addressablegcode.New[float32]('X', 12.3)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%s", gc)

	// Output: X12.300
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

	// save a reference to gcode addressable instance in a gcoder interface
	var gc gcode.Gcoder = gca

	if !gc.HasAddress() {
		fmt.Println("Ups! gcode not contain address")
		return
	}

	// try recovery the reference to gcode addressable instance from a gcoder interface
	// using assertion
	if value, ok := gc.(gcode.AddressableGcoder[int32]); ok {
		add := value.Address()
		fmt.Printf("the int32 address recovered is %v\n", add)
	}

	// Output: the int32 address recovered is 66555
}

func ExampleGcode_Compare() {

	addressableGcode, err := addressablegcode.New('X', "\"Hola mundo!\"")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	anotherAddressableGcode, err := addressablegcode.New('X', "\"Hello World!\"")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	fmt.Printf("%v\n", addressableGcode.Compare(anotherAddressableGcode))

	// Output: false
}

func ExampleGcode_Compare_second() {

	addressableGcode, err := addressablegcode.New('X', "\"Hola mundo!\"")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	unaddressableGcode, err := unaddressablegcode.New('X')
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	fmt.Printf("%v\n", addressableGcode.Compare(unaddressableGcode))

	// Output: false
}

func ExampleGcode_HasAddress() {

	gca, err := addressablegcode.New('X', "\"Hola mundo!\"")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	fmt.Printf("%v\n", gca.HasAddress())

	// Output: true
}

func ExampleGcode_HasAddress_second() {

	gca, err := addressablegcode.New('X', "\"Hola mundo!\"")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	// assign reference from gcode addressable instance to gcoder interface
	var gc gcode.Gcoder = gca

	fmt.Printf("%v\n", gc.HasAddress())

	// Output: true
}

func ExampleGcode_String() {

	gca, err := addressablegcode.New('X', "\"Hola mundo!\"")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	fmt.Println(gca)

	// Output: X"Hola mundo!"
}

func ExampleGcode_String_second() {

	gca, err := addressablegcode.New('X', "\"Hola mundo!\"")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	// assign reference from gcode addressable instance to gcoder interface
	var gc gcode.Gcoder = gca

	fmt.Println(gc)

	// Output: X"Hola mundo!"
}

func ExampleGcode_Word() {
	gca, err := addressablegcode.New[int32]('D', 0)
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	fmt.Println(gca.Word())

	// Output: 68
}
