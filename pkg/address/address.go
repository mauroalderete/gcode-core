// address package allows store and management the representation of the part assigned to the value of a gcode.
//
// A gcode can have or doesn't have an address. When it has, the address can be of the int32, float32 or string data type.
//
// This package contains a constructor that returns an address of some of these data types defined by the AddressType interface.
//
// An address struct is bound with a series of methods and functions that allow you to operate with the value of the address.
//
// At the same time, this package defines some error interfaces to handle the incidences that occur during the construction.
package address

import (
	"fmt"
	"strconv"
	"strings"
)

// AddressType interface defines the restriction type used as type generic to Address model
type AddressType interface {
	string | int32 | float32
}

// Address[T AddressType] struct model a address of a gcode.
//
// An address can to be the int32, float32, string data type. It is defined by the restriction with AddressType interface.
//
// Expose a Value field that stores the useful data.
type Address[T AddressType] struct {
	Value T
}

// AddressStringContainInvalidCharsError struct model an error that happens when a potential address value of the string data type contains chars doesn't allowed
type AddressStringContainInvalidCharsError struct {
	Value string
}

func (a *AddressStringContainInvalidCharsError) Error() string {
	return fmt.Errorf("gcode's address string contain invalid chars: %v", a.Value).Error()
}

// AddressStringQuoteError struct model an error that happens when a potential address value of the string data type contains chars doesn't allowed
type AddressStringQuoteError struct {
	Value string
}

func (a *AddressStringQuoteError) Error() string {
	return fmt.Errorf("gcode's address string has an invalid use of the quotes: %v", a.Value).Error()
}

// AddressStringTooShortError struct model an error that happens when a potential address value of the string data type is too short to be a valid data
type AddressStringTooShortError struct {
	Value string
}

func (a *AddressStringTooShortError) Error() string {
	return fmt.Errorf("gcode's address string is too short: %v", a.Value).Error()
}

func (a *Address[T]) String() string {
	if value, ok := any(a.Value).(float32); ok {
		sv := strconv.FormatFloat(float64(value), 'f', -1, 32)
		if !strings.Contains(sv, ".") {
			sv += ".0"
		}
		return sv
	}
	return fmt.Sprintf("%v", a.Value)
}

// Compare allow knowing if an address is equal to the other address object
func (add *Address[T]) Compare(a Address[T]) bool {
	return add.Value == a.Value
}

// NewAddress[T AddressType] return a pointer to a new instance of an address struct.
//
// Return an error when the value does not correspond to a format valid
func NewAddress[T AddressType](value T) (*Address[T], error) {

	if value, ok := any(value).(string); ok {
		err := IsAddressStringValid(value)
		if err != nil {
			return nil, err
		}
	}

	newAddress := Address[T]{
		Value: value,
	}

	return &newAddress, nil
}

// IsAddressStringValid allow knowing if a string input can be an address value of string data type valid.
//
// Return an of the error types defined in this package if s string is invalid.
//
// Return nil if s string satisfies the format of address value of string data type.
func IsAddressStringValid(s string) error {
	if len(s) <= 1 {
		return &AddressStringTooShortError{Value: s}
	}

	if strings.ContainsAny(s, "\t\n\r") {
		return &AddressStringContainInvalidCharsError{Value: s}
	}

	if !(s[0] == '"' && s[len(s)-1] == '"') {
		return &AddressStringQuoteError{}
	}

	for _, v := range strings.Split(s[1:len(s)-1], "\"\"") {
		if strings.ContainsRune(v, '"') {
			return &AddressStringQuoteError{}
		}
	}

	return nil
}
