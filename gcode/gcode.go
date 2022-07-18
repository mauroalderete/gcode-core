// gcode package contains the model to represent a single gcode expression.
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
// The address model allows store and management the representation of the part assigned to the value of a gcode.
//
// A gcode can have or doesn't have an address. When it has, the address must be of either int32, uin32, float32 or string data type.
// This package contains a constructor that returns an address of some of these data types defined by the AddressType interface.
//
// gcode package define a series of interfaces to implement mainly two kind gcode structs.
//
// Gcoder interface is used to model a gcode with a stand-alone word.
// Is said, allows to get a gcode without an address component.
//
// AddressableGcoder interfaces, join to AddressType interface, allows constructing a typical gcode object but, also,
// includes an address struct of the data type defined by the AddressType interface in the address package.
// This we allow to handle a gcode with a string address, address int32, address uint32 or address float32.
//
// AddressableGcoder interface wrap Gcoder interface.
// This means, that all the AddressableGcoder objects are as well Gcoder objects.
//
// All Gcoder instances implement a word model to represent the letter of a gcode command that identifies a single action or a parameter.
// The Word method allow get a byte that represent this command of gcode.
// This value is a letter in uppercase according to specification to gcode.
// Each one of them can come accompanied by an address or not,
//
// Exist, different classes the words, some are commands and others are used likes parameters for the commands.
//
// An address struct that implement AddressableGcoder is bound with a series of methods that allow you to operate with the value of the address.
//
// Foremor, gcode package includes GcoderFactory.
// It is a interace that contains some methods that return a new Gcoder instance or a new AddressableGcoder instance.
//
// gcode package provides two packages that implement all interfaces ready to use.
package gcode

import "fmt"

//#region interfaces

// Gcoder interface allows getting the word that gives meaning to the gcode and knowing if includes an address or not.
type Gcoder interface {
	// Stringer (via the embedded fmt.Stringer interface) return the Gcoder value in string format.
	fmt.Stringer

	// Compare allows comparing the values of the current Gcoder with another.
	Compare(Gcoder) bool

	// HasAddress indicates if the current Gcoder has address element or not.
	HasAddress() bool

	// Word returns the word value that defines the gcode.
	Word() byte
}

// AddressType interface defines the restriction type used as type generic to Address model.
// The address can be numeric or some cases strings.
// The numeric values supported changes depending on the word value.
type AddressType interface {
	string | int32 | float32 | uint32
}

// AddressableGcoder compose the Gcoder interface and add methods to handle the address value.
type AddressableGcoder[T AddressType] interface {
	// Gcoder (via the embedded gcode.Gcoder interface) allow converts AddressableGcoder in a Gcoder element.
	Gcoder

	// Address return addres value of the data type T.
	Address() T

	// SetAddress stores addres value of the data type T.
	// Return an error if the format the value is not valid.
	SetAddress(value T) error
}

// GcoderFactory defines constructors to instance each kind of gcode object.
type GcoderFactory interface {
	// Create a Gcoder instance that not use address element.
	NewUnaddressableGcode(word byte) (Gcoder, error)

	// Create a AddressableGcoder instance of a specific type.
	NewAddressableGcodeUint32(word byte, address uint32) (AddressableGcoder[uint32], error)
	NewAddressableGcodeInt32(word byte, address int32) (AddressableGcoder[int32], error)
	NewAddressableGcodeFloat32(word byte, address float32) (AddressableGcoder[float32], error)
	NewAddressableGcodeString(word byte, address string) (AddressableGcoder[string], error)
}

//#endregion
//#region package functions

// IsValid allow knowledge if a potential word value contains a value valid according to a specification gcode.
//
// The set valid values are hard coding and they correspond to a [ReRap documentation].
//
// [ReRap documentation]: https://reprap.org/wiki/G-code
func IsValidWord(word byte) error {

	switch word {
	case 'G', 'M', 'T', 'S', 'P', 'X', 'Y', 'Z', 'U', 'V', 'W', 'I', 'J', 'D', 'H', 'F', 'R', 'Q', 'E', 'N', '*':
		return nil
	}

	return fmt.Errorf("gcode's word has invalid value: %v", word)
}

//#endregion
