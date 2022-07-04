package address

import (
	"fmt"
	"strings"
)

type AddressType interface {
	string | int32 | float32
}

type Address[T AddressType] struct {
	value T
}

type AddressStringContainInvalidCharsError struct {
	Value string
}

func (a *AddressStringContainInvalidCharsError) Error() string {
	return fmt.Errorf("gcode's address string contain invalid chars: %v", a.Value).Error()
}

type AddressStringQuoteError struct {
	Value string
}

func (a *AddressStringQuoteError) Error() string {
	return fmt.Errorf("gcode's address string has an invalid use of the quotes: %v", a.Value).Error()
}

type AddressStringTooShortError struct {
	Value string
}

func (a *AddressStringTooShortError) Error() string {
	return fmt.Errorf("gcode's address string is too short: %v", a.Value).Error()
}

func (a *Address[T]) String() string {
	if value, ok := any(a.value).(string); ok {
		return fmt.Sprintf("\"%s\"", value)
	}

	return fmt.Sprintf("%v", a.value)
}

func (a *Address[T]) Value() T {
	return a.value
}

func (a *Address[T]) Compare(address Address[T]) bool {
	return a.value == address.value
}

func (a *Address[T]) CompareValue(value T) bool {
	return a.value == value
}

func NewAddress[T AddressType](address T) (*Address[T], error) {

	if value, ok := any(&address).(string); ok {
		err := isAddressStringValid(value)
		if err != nil {
			return nil, err
		}
	}

	newAddress := Address[T]{
		value: address,
	}

	return &newAddress, nil
}

func isAddressStringValid(address string) error {
	if strings.ContainsAny(address, "\t\n\r") {
		return &AddressStringContainInvalidCharsError{Value: address}
	}

	if len(address) <= 1 {
		return &AddressStringTooShortError{Value: address}
	}

	if address[0] != '"' && address[len(address)-1] != '"' {
		return &AddressStringContainInvalidCharsError{}
	}

	return nil
}
