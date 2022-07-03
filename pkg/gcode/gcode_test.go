package gcode_test

import (
	"fmt"
	"testing"

	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/address"
	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/gcode"
	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/word"
)

func TestNewGcode(t *testing.T) {
	t.Run("Gcode valids", func(t *testing.T) {

		t.Run("Gcode with integer address", func(t *testing.T) {
			gc, err := gcode.NewGcode('X', "12")
			if err != nil {
				t.Errorf("got %v, want X12", err)
				return
			}
			if gc.String() != "X12" {
				t.Errorf("got %s, want X12", gc)
			}
		})

		t.Run("Gcode with float address", func(t *testing.T) {
			gc, err := gcode.NewGcode('X', "12.3")
			if err != nil {
				t.Errorf("got %v, want X12.3", err)
				return
			}
			if gc.String() != "X12.3" {
				t.Errorf("got %s, want X12.3", gc)
			}
		})

		t.Run("Gcode with string address", func(t *testing.T) {
			gc, err := gcode.NewGcode('X', "\"lorem ipsu\"")
			if err != nil {
				t.Errorf("got %v, want X\"lorem ipsu\"", err)
				return
			}
			if gc.String() != "X\"lorem ipsu\"" {
				t.Errorf("got %s, want X\"lorem ipsu\"", gc)
			}
		})
	})

	t.Run("Gcode invalids", func(t *testing.T) {
		t.Run("word invalid value", func(t *testing.T) {
			_, err := gcode.NewGcode('+', "12")
			if err == nil {
				t.Errorf("got error == nil, want gcode's word has invalid value")
				return
			}
			if err.Error() != "gcode's word has invalid value" {
				t.Errorf("got %v, want gcode's word has invalid value", err)
			}
		})

		t.Run("address with invalid characters", func(t *testing.T) {
			const msg = "gcode's address contain invalid chars"
			_, err := gcode.NewGcode('X', "\n")
			if err == nil {
				t.Errorf("got error == nil, %s", msg)
				return
			}
			if err.Error() != msg {
				t.Errorf("got %v, want %s", err, msg)
			}
		})

		t.Run("address with single character not numeric", func(t *testing.T) {
			const msg = "when a gcode's address isn't a numeric value, it must contain a string close in between double quotes, a"
			_, err := gcode.NewGcode('X', "a")
			if err == nil {
				t.Errorf("got error == nil, %s", msg)
				return
			}
			if err.Error() != msg {
				t.Errorf("got %v, want %s", err, msg)
			}
		})

		t.Run("address with data type invalid", func(t *testing.T) {
			const msg = "gcode's address value could not be regconocide like integer, float number or string, 12.12.12"
			_, err := gcode.NewGcode('X', "12.12.12")
			if err == nil {
				t.Errorf("got error == nil, %s", msg)
				return
			}
			if err.Error() != msg {
				t.Errorf("got %v, want %s", err, msg)
			}
		})
	})
}

func ExampleNewGcode() {

	gc, err := gcode.NewGcode('X', "12.3")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%s", gc)

	// Output: X12.3
}

func ExampleGcode_String() {

	gc, err := gcode.NewGcode('X', "12.3")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	fmt.Println(gc)

	// Output: X12.3
}

func ExampleGcode_Word() {
	gc, err := gcode.NewGcode('D', "0")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	var w word.Word = gc.Word()

	fmt.Println(w.String())

	// Output: D
}

func ExampleGcode_Address() {
	gc, err := gcode.NewGcode('N', "66555")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	var a address.Address = gc.Address()

	fmt.Println(a.String())

	// Output: 66555
}

func ExampleGcode_Compare() {
	gc1, err := gcode.NewGcode('G', "33")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	gc2, err := gcode.NewGcode('G', "34")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	fmt.Printf("%t", gc1.Compare(gc2))

	// Output: false
}
