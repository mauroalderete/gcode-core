package gcode

import (
	"fmt"
	"strconv"
	"strings"
)

//#region address struct

// AddressType interface defines the restriction type used as type generic to Address model
type AddressType interface {
	string | int32 | float32 | uint32
}

// Address[T AddressType] struct model a address of a gcode.
//
// An address can to be the int32, float32, string data type.
// It is defined by the restriction with AddressType interface.
//
// Expose a Value field that stores the useful data.
type Address[T AddressType] struct {
	value T
}

// Value return the value of the address
func (a *Address[T]) Value() T {
	return a.value
}

// SetValue allow to store a new value
//
// If the address data type is string then the new value is verified.
// If it doesn't satisfy the a string format then SetValue returns an error.
func (a *Address[T]) SetValue(value T) error {

	if ok, err := isGenericValueAnStringAddressValid(value); ok {
		if err != nil {
			return fmt.Errorf("failed set the value %v at the %T address: %w", value, value, err)
		}
	}

	a.value = value

	return nil
}

// String return the value in address as string format
//
// If the address data type is a float32 then the format returned will contain at least one decimal,
// regardless of whether the value corresponds to an integer.
//
// For example, for the float32 value 13, the string returned will be "13.0".
func (a *Address[T]) String() string {
	if float32Value, ok := any(a.Value()).(float32); ok {
		sv := strconv.FormatFloat(float64(float32Value), 'f', -1, 32)
		if !strings.Contains(sv, ".") {
			sv += ".0"
		}
		return sv
	}
	return fmt.Sprintf("%v", a.Value())
}

// Compare allow knowing if an address is equal to the other address object
func (a *Address[T]) Compare(address Address[T]) bool {
	return a.Value() == address.value
}

//#endregion
//#region constructors

// NewAddress[T AddressType] return a pointer to a new instance of an address struct.
//
// Return an error when the value does not correspond to a format valid
func NewAddress[T AddressType](value T) (*Address[T], error) {

	if ok, err := isGenericValueAnStringAddressValid(value); ok {
		if err != nil {
			return nil, fmt.Errorf("failed to create an string address instance using the expression %v: %w", value, err)
		}
	}

	return &Address[T]{
		value: value,
	}, nil
}

//#endregion
//#region package functions

// IsAddressStringValid allow knowing if a string input can be an address value of string data type valid.
//
// Return an error if s string is invalid.
//
// Return nil if s string satisfies the format of address value of string data type.
func IsAddressStringValid(s string) error {
	if len(s) <= 1 {
		return fmt.Errorf("gcode address string is too short: %v", s)
	}

	if strings.ContainsAny(s, "\t\n\r") {
		return fmt.Errorf("gcode address string contains invalid chars: %v", s)
	}

	if !(s[0] == '"' && s[len(s)-1] == '"') {
		return fmt.Errorf("gcode address string isn't enclosed in quotes: %v", s)
	}

	for _, v := range strings.Split(s[1:len(s)-1], "\"\"") {
		if strings.ContainsRune(v, '"') {
			return fmt.Errorf("gcode address string hasn't a valid use of the quotes: %v", s)
		}
	}

	return nil
}

//#endregion

//#region private functions

// isGenericValueAnStringAddressValid return true if the value is an string and return error if this string value is not string address valid.
//
// It returns false if the value is not of the string data type.
// In this case, it does not be to verify any string, therefore never it returns an error.
func isGenericValueAnStringAddressValid[T AddressType](value T) (bool, error) {
	if stringValue, ok := any(value).(string); ok {
		err := IsAddressStringValid(stringValue)
		if err != nil {
			return true, err
		}

		return true, nil
	}

	return false, nil
}

//#endregion
