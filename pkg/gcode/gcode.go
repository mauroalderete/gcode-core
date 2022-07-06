// gcode package to represent a simple g-code unit command.
//
// G-code units are a minimal pieces in that a block g-code is formed.
// Of this form, a simple g-code can be interpreted as a command, parameter, or for any other special purpose.
//
// It consists of one letter directly followed by a number, or can be only a stand-alone letter (Word).
// The letter gives information about the meaning of the g-code. Numbers can be integers (128) or fractional numbers (12.42), depending on context.
//
// For example, an X coordinate can take integers (X175) or fractionals (X17.62).
package gcode

import (
	"fmt"

	gcode_address "github.com/mauroalderete/gcode-skew-transform-cli/pkg/address"
	gcode_word "github.com/mauroalderete/gcode-skew-transform-cli/pkg/word"
)

type Gcoder interface {
	fmt.Stringer
	Word() gcode_word.Word
	HasAddress() bool
}

// Gcode implements a single g-code.
//
// Allow model a gcode command from two elements that match with a word struct and an address struct.
// A Gcode can contain only an AddressType data type. Is say, any value of the string, int32 or float32 data type.
//
// Word struct indicates the command.
// Address struct indicates the value used by the command. Can be nil, if the command isn't used any value.
type Gcode struct {
	word *gcode_word.Word
}

// String return G-code formatted
func (g *Gcode) String() string {
	return g.word.String()
}

// Word return the word struct that is contained in the G-code
func (g *Gcode) Word() gcode_word.Word {
	return *g.word
}

// Word return the word struct that is contained in the G-code
func (g *Gcode) HasAddress() bool {
	return false
}

type GcodeAddresser[T gcode_address.AddressType] interface {
	Gcoder
	Address() *gcode_address.Address[T]
}

type GcodeAddressable[T gcode_address.AddressType] struct {
	*Gcode
	address *gcode_address.Address[T]
}

// String return G-code formatted
func (g *GcodeAddressable[T]) String() string {
	return fmt.Sprintf("%s%s", g.word.String(), g.address.String())
}

// Word return the word struct that is contained in the G-code
func (g *GcodeAddressable[T]) Word() gcode_word.Word {
	return *g.word
}

// Word return the word struct that is contained in the G-code
func (g *GcodeAddressable[T]) HasAddress() bool {
	return true
}

// Address return the Address struct that is contained in the Gcode.
// The address returned can be implement an string, int32, float32 data types
func (g *GcodeAddressable[T]) Address() *gcode_address.Address[T] {
	return g.address
}

// NewGcode is the constructor to instance a G-code struct.
//
// Receive a word that represents the letter of the command and another string that represent the address of gcode.
//
// In general, a word consists of a single letter. Instead, an address can be a number (integer or float) or a string of characters between double-quotes.
//
// In any case, this method will verify the format of both parameters and return nil with an error description if necessary.
func NewGcode(word byte) (Gcoder, error) {

	// Try instace Word struct
	wrd, err := gcode_word.NewWord(word)
	if err != nil {
		return nil, err
	}

	return &Gcode{
		word: wrd,
	}, nil
}

// NewGcode is the constructor to instance a G-code struct.
//
// Receive a word that represents the letter of the command and another string that represent the address of gcode.
//
// In general, a word consists of a single letter. Instead, an address can be a number (integer or float) or a string of characters between double-quotes.
//
// In any case, this method will verify the format of both parameters and return nil with an error description if necessary.
func NewGcodeAddressable[T gcode_address.AddressType](word byte, address T) (GcodeAddresser[T], error) {

	// Try instace Word struct
	wrd, err := gcode_word.NewWord(word)
	if err != nil {
		return nil, err
	}

	// Try instace Address struct
	add, err := gcode_address.NewAddress(address)
	if err != nil {
		return nil, err
	}

	return &GcodeAddressable[T]{
		Gcode: &Gcode{
			word: wrd,
		},
		address: add,
	}, nil
}
