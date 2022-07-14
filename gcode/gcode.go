package gcode

import (
	"fmt"
)

//#region gcode struct

// Gcoder interface allows getting the word that gives meaning to the gcode and knowing if includes an address or not
type Gcoder interface {
	fmt.Stringer
	Word() Word
	Compare(Gcoder) bool
	HasAddress() bool
}

// Gcode struct model a gcode expression with a stand-alone word.
//
// Allow model a gcode that does not contain an address. For this, stores a word.
type Gcode struct {
	word Word
}

// String return Gcode formatted
func (g *Gcode) String() string {
	return g.word.String()
}

// Word return a copy of the word struct in the gcode
func (g *Gcode) Word() Word {
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
type GcodeAddressable[T AddressType] struct {
	*Gcode
	address *Address[T]
}

// String return gcode formatted
func (g *GcodeAddressable[T]) String() string {
	return fmt.Sprintf("%s%s", g.word.String(), g.address.String())
}

// Word return a copy of the word struct in the gcode
func (g *GcodeAddressable[T]) Word() Word {
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
func (g *GcodeAddressable[T]) Address() *Address[T] {
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
	wrd, err := NewWord(word)
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
func NewGcodeAddressable[T AddressType](word byte, address T) (*GcodeAddressable[T], error) {

	// Try instace Word struct
	wrd, err := NewWord(word)
	if err != nil {
		return nil, fmt.Errorf("failed to create an addressable gcode instance of type %T when trying to use %v word: %w", address, word, err)
	}

	// Try instace Address struct
	add, err := NewAddress(address)
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
