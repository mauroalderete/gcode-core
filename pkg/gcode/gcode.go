// gcode package to represent a single gcode expression.
//
// Gcode expressions are the minimal pieces in a block (line with a command at the gcode file).
//
// A gcode can be interpreted as a command, parameter, or for any other special purpose.
//
// It consists of only a stand-alone letter, named word.
// Or the word directly followed by a numeric or alphanumeric value named address.
// The word gives information about the meaning of the gcode.
// Instead, the address asserts a value to gcode.
// It can be integers (128) or fractional numbers (12.42) or strings depending on the context.
//
// For example, an X coordinate can take integers (X175) or fractionals (X17.62).
//
// gcode package implements two constructor methods to instantiate a Gcode struct or a GcodeAddressable struct.
//
// One of them is used to model a gcode with a stand-alone word.
// Is said, allows to get a gcode without an address component.
//
// The other method is a generic method that allows constructing a typical gcode object but, also,
// includes an address struct of the data type defined by the AddressType interface in the address package.
// This constructor we allow to get a gcode with a string address, or address int32 or address float32.
//
// All gcode types implement a Gcoder interface.
// Instead, the gcodes that contain an address implement the GcodeAddresser interface.
//
// GcodeAddresser interface wrap Gcoder interface.
// This means, that all the GcodeAddresser objects are as well Gcoder objects.
package gcode

import (
	"fmt"

	gcode_address "github.com/mauroalderete/gcode-skew-transform-cli/pkg/address"
	gcode_word "github.com/mauroalderete/gcode-skew-transform-cli/pkg/word"
)

//#region gcode struct

// Gcoder interface allows getting the word that gives meaning to the gcode and knowing if includes an address or not
type Gcoder interface {
	fmt.Stringer
	Word() gcode_word.Word
	Compare(Gcoder) bool
	HasAddress() bool
}

// Gcode struct model a gcode expression with a stand-alone word.
//
// Allow model a gcode that does not contain an address. For this, stores a word.
type Gcode struct {
	word gcode_word.Word
}

// String return Gcode formatted
func (g *Gcode) String() string {
	return g.word.String()
}

// Word return a copy of the word struct in the gcode
func (g *Gcode) Word() gcode_word.Word {
	return g.word
}

// Compare allows checking if the current entity is equal to a Gcoder input
//
// This method is executed when to be called from a Gcode instance or a Gcoder instance that contains a reference to a Gcode object.
// If the input Gcoder does not implement some gcode.Addressable data type then it returns false
func (g *Gcode) Compare(gcode Gcoder) bool {

	if gcode == nil {
		return false
	}

	if gcode.HasAddress() {
		return false
	}

	return g.word == gcode.Word()
}

// HasAddress indicate if the gcode contains or not an address.
//
// This method is executed when to be called from a Gcode instance or a Gcoder instance that contains a reference to a Gcode object.
// In both cases, it always returns false.
func (g *Gcode) HasAddress() bool {
	return false
}

//#endregion
//#region gcode addressable struct

// GcodeAddressable struct that implements GcodeAddresser interface
//
// Is composed of a gcode struct and includes an address field to store an address instance
type GcodeAddressable[T gcode_address.AddressType] struct {
	*Gcode
	address *gcode_address.Address[T]
}

// String return gcode formatted
func (g *GcodeAddressable[T]) String() string {
	return fmt.Sprintf("%s%s", g.word.String(), g.address.String())
}

// Word return a copy of the word struct in the gcode
func (g *GcodeAddressable[T]) Word() gcode_word.Word {
	return g.word
}

// Compare allows checking if the current entity is equal to a Gcoder input
//
// This method is executed when to be called from a Gcode instance or a Gcoder instance that contains a reference to a Gcode object.
// If the input Gcoder does not implement some gcode.Addressable data type then it returns false
func (g *GcodeAddressable[T]) Compare(gcode Gcoder) bool {

	if gca, ok := gcode.(*GcodeAddressable[T]); ok && gca != nil {
		return g.word == gca.word && g.address.Compare(*gca.address)
	}

	return false
}

// HasAddress indicate if the gcode contain or not an address.
//
// This method is called from a GcodeAddressable instance or a GcodeAddresser instance that contains a reference to a GcodeAddressable object.
// Always return true.
func (g *GcodeAddressable[T]) HasAddress() bool {
	return true
}

// Address return the Address struct that is contained in the Gcode.
func (g *GcodeAddressable[T]) Address() *gcode_address.Address[T] {
	return g.address
}

//#endregion
//#region constructors

// NewGcode is the constructor to instance a Gcode struct.
//
// Receive a word that represents the letter of the command of a gcode.
//
// If the word is an unknown symbol it returns nil with an error description.
func NewGcode(word byte) (Gcoder, error) {

	// Try instace Word struct
	wrd, err := gcode_word.NewWord(word)
	if err != nil {
		return nil, fmt.Errorf("failed to create an gcode instance when trying to use %v word: %w", word, err)
	}

	return &Gcode{
		word: *wrd,
	}, nil
}

// NewGcode is the constructor to instance a Gcode struct that includes an address.
//
// Receive a word that represents the letter of the command and another value that represent the address of gcode.
//
// The value can be string, int32 or float 32 data type.
//
// In any case, this method will verify the format of both parameters and return nil with an error description if necessary.
func NewGcodeAddressable[T gcode_address.AddressType](word byte, address T) (*GcodeAddressable[T], error) {

	// Try instace Word struct
	wrd, err := gcode_word.NewWord(word)
	if err != nil {
		return nil, fmt.Errorf("failed to create an addressable gcode instance of type %T when trying to use %v word: %w", address, word, err)
	}

	// Try instace Address struct
	add, err := gcode_address.NewAddress(address)
	if err != nil {
		return nil, fmt.Errorf("failed to create an addressable gcode instance of type %T when trying to use %v address: %w", address, address, err)
	}

	return &GcodeAddressable[T]{
		Gcode: &Gcode{
			word: *wrd,
		},
		address: add,
	}, nil
}

//#endregion
