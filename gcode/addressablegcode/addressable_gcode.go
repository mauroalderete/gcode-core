// addressablegcode package implements gcode.AddressableGcoder interface to model a gcode with an address element.
//
// Define a Gcode struct that implement gcode.AddressableGcoder interface.
// This struct contain a word field to store the word value and an address field to store the addres value.
//
// Like AddressableGcoder interface, this Gcode struct use generics with the restrictions defined by AddressType.
//
// A "New" constructor method allow to instance new Gcode[T] object.
// This method use some internal rules combined with gcode.IsValidWord to validate the inputs before create any instance
package addressablegcode

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mauroalderete/gcode-cli/gcode"
)

//#region gcode addressable struct

// Gcode struct that implements GcodeAddresser interface.
//
// Is composed of a gcode struct and includes an address field to store an address instance.
type Gcode[T gcode.AddressType] struct {
	//Gcoder (via embedded gcode.Gcoder interface)
	gcode.Gcoder

	// word stores the word value of the gcode that we are modeling
	word byte

	// address stores the address value of the gcode that we are modeling
	address T
}

// Address return the value of the address
func (g *Gcode[T]) Address() T {
	return g.address
}

// Compare allows checking if the current entity is equal to a Gcoder input
//
// This method is executed when to be called from a Gcode instance or a Gcoder instance that contains a reference to a Gcode object.
// If the input Gcoder does not implement some addressablegcode.Gcode[T] data type then it returns false
func (g *Gcode[T]) Compare(gcode gcode.Gcoder) bool {

	if gca, ok := gcode.(*Gcode[T]); ok && gca != nil {
		return g.word == gca.Word() && g.address == gca.address
	}

	return false
}

// HasAddress indicate if the gcode contain or not an address.
//
// This method is called from a GcodeAddressable instance or a GcodeAddresser instance that contains a reference to a GcodeAddressable object.
// Always return true.
func (g *Gcode[T]) HasAddress() bool {
	return true
}

// SetAddress allow to store a new value
//
// If the address data type is string then the new value is verified.
// If it doesn't satisfy the a string format then SetAddress returns an error.
func (g *Gcode[T]) SetAddress(address T) error {

	if ok, err := isGenericValueAnStringAddressValid(address); ok {
		if err != nil {
			return fmt.Errorf("failed set the value %v at the %T address: %w", address, address, err)
		}
	}

	g.address = address

	return nil
}

// String return gcode formatted
func (g *Gcode[T]) String() string {

	if float32Value, ok := any(g.address).(float32); ok {
		sv := strconv.FormatFloat(float64(float32Value), 'f', -1, 32)
		if !strings.Contains(sv, ".") {
			sv += ".0"
		}
		return fmt.Sprintf("%s%s", string(g.word), sv)
	}

	return fmt.Sprintf("%s%v", string(g.word), g.address)
}

// Word return a copy of the word struct in the gcode
func (g *Gcode[T]) Word() byte {
	return g.word
}

//#endregion
//#region package constructor

// New return a new Gcode[T] instance or error if some inputs are invalids
// word is the letter that compose the gcode
// address is the value of the gcode
func New[T gcode.AddressType](word byte, address T) (*Gcode[T], error) {
	// Try instace Word struct
	err := gcode.IsValidWord(word)
	if err != nil {
		return nil, fmt.Errorf("failed to create an addressable gcode instance of type %T when trying to use %v word: %w", address, word, err)
	}

	// Try instace Address struct
	if ok, err := isGenericValueAnStringAddressValid(address); ok {
		if err != nil {
			return nil, fmt.Errorf("failed to create an string address instance using the expression %v: %w", address, err)
		}
	}

	return &Gcode[T]{
		word:    word,
		address: address,
	}, nil
}

//#endregion
//#region private functions

// isAddressStringValid allow knowing if a string input can be an address value of string data type valid.
//
// Return an error if s string is invalid.
//
// Return nil if s string satisfies the format of address value of string data type.
func isAddressStringValid(address string) error {
	if len(address) <= 1 {
		return fmt.Errorf("gcode address string is too short: %v", address)
	}

	if strings.ContainsAny(address, "\t\n\r") {
		return fmt.Errorf("gcode address string contains invalid chars: %v", address)
	}

	if !(address[0] == '"' && address[len(address)-1] == '"') {
		return fmt.Errorf("gcode address string isn't enclosed in quotes: %v", address)
	}

	for _, v := range strings.Split(address[1:len(address)-1], "\"\"") {
		if strings.ContainsRune(v, '"') {
			return fmt.Errorf("gcode address string hasn't a valid use of the quotes: %v", address)
		}
	}

	return nil
}

// isGenericValueAnStringAddressValid return true if the value is an string and return error if this string value is not string address valid.
//
// It returns false if the value is not of the string data type.
// In this case, it does not be to verify any string, therefore never it returns an error.
func isGenericValueAnStringAddressValid[T gcode.AddressType](value T) (bool, error) {
	if stringValue, ok := any(value).(string); ok {
		err := isAddressStringValid(stringValue)
		if err != nil {
			return true, err
		}

		return true, nil
	}

	return false, nil
}

//#endregion
