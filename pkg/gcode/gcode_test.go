package gcode_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/address"
	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/gcode"
	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/word"
)

func TestNewGcode(t *testing.T) {
	t.Run("valids", func(t *testing.T) {

		t.Run("address integer", func(t *testing.T) {
			gc, err := gcode.NewGcode[int32]('X', 12)
			if err != nil {
				t.Errorf("got %v, want X12", err)
				return
			}
			if gc.String() != "X12" {
				t.Errorf("got %s, want X12", gc)
			}
		})

		t.Run("address float", func(t *testing.T) {
			gc, err := gcode.NewGcode[float32]('X', 12.3)
			if err != nil {
				t.Errorf("got %v, want X12.3", err)
				return
			}
			if gc.String() != "X12.3" {
				t.Errorf("got %s, want X12.3", gc)
			}
		})

		t.Run("address string", func(t *testing.T) {
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

	t.Run("invalids", func(t *testing.T) {

		t.Run("word", func(t *testing.T) {
			cases := []struct {
				word    byte
				address int32
				err     error
			}{
				{'+', 2, &word.WordInvalidValueError{Value: '+'}},
				{'\t', 2, &word.WordInvalidValueError{Value: '\t'}},
				{'"', 2, &word.WordInvalidValueError{Value: '"'}},
			}

			for i, c := range cases {
				t.Run(fmt.Sprintf("(%v)", i), func(t *testing.T) {
					_, err := gcode.NewGcode(c.word, c.address)
					if err.Error() != c.err.Error() {
						t.Errorf("got %v, want %v", err, c.err)
					}
				})
			}
		})

		t.Run("address string", func(t *testing.T) {
			cases := []struct {
				word    byte
				address string
				err     error
			}{
				{'X', "", &address.AddressStringTooShortError{Value: ""}},
				{'X', "\"\t\"", &address.AddressStringContainInvalidCharsError{Value: "\"\t\""}},
				{'X', "\"\"\"", &address.AddressStringQuoteError{Value: ""}},
			}

			for i, c := range cases {
				t.Run(fmt.Sprintf("(%v)", i), func(t *testing.T) {
					_, err := gcode.NewGcode(c.word, c.address)
					if err.Error() != c.err.Error() {
						t.Errorf("got %v, want %v", err, c.err)
					}
				})
			}
		})
	})
}

func ExampleNewGcode() {

	gc, err := gcode.NewGcode[float32]('X', 12.3)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%s", gc)

	// Output: X12.3
}

func ExampleGcode_String() {

	gc, err := gcode.NewGcode('X', "\"Hola mundo!\"")
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	fmt.Println(gc)

	// Output: X"Hola mundo!"
}

func ExampleGcode_Word() {
	gc, err := gcode.NewGcode[int32]('D', 0)
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	var w word.Word = gc.Word()

	fmt.Println(w.String())

	// Output: D
}

func ExampleGcode_Address() {
	gc, err := gcode.NewGcode[int32]('N', 66555)
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	var a address.Address[int32] = gc.Address()

	fmt.Println(a.String())

	// Output: 66555
}

func ExampleGcode_Compare() {
	gc1, err := gcode.NewGcode[float32]('G', math.Pi)
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	gc2, err := gcode.NewGcode[float32]('X', math.Pi)
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	fmt.Printf("%t", gc1.Compare(gc2))

	// Output: false
}
