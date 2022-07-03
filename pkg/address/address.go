package address

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Address struct {
	value string
}

type AddressType interface {
	string | int32 | float32
}

func (a *Address) String() string {
	return a.value
}

func (a *Address) ToInteger() (int, error) {
	return toInteger(a.value)
}

func (a *Address) ToFloat() (float32, error) {
	return toFloat(a.value)
}

func (a *Address) IsInteger() bool {
	_, err := a.ToInteger()

	return err == nil
}

func (a *Address) IsFloat() bool {
	_, err := a.ToFloat()
	return err == nil
}

func (a *Address) IsNumber() bool {

	return a.IsInteger() || a.IsFloat()
}

func (a *Address) Compare(address fmt.Stringer) bool {
	return a.value == address.String()
}

func NewAddress[T AddressType](address T) (*Address, error) {

	var addressValue string

	switch a := any(&address).(type) {
	case *string:
		{
			addressValue = *a
		}
	case *int32:
		{
			addressValue = strconv.FormatInt(int64(*a), 10)
		}
	case *float32:
		{
			addressValue = strconv.FormatFloat(float64(*a), 'f', -1, 32)
		}
	default:
		{
			return nil, errors.New("address is invalid type")
		}
	}

	err := isValid(addressValue)
	if err != nil {
		return nil, err
	}

	newAddress := Address{
		value: addressValue,
	}

	return &newAddress, nil
}

func isValid(address string) error {

	if strings.ContainsAny(address, "\t\n\r") {
		return errors.New("gcode's address contain invalid chars")
	}

	_, err := toInteger(address)
	_, err1 := toFloat(address)

	if err == nil || err1 == nil {
		return nil
	}

	if len(address) == 1 {
		return fmt.Errorf("when a gcode's address isn't a numeric value, it must contain a string close in between double quotes, %s", address)
	}

	if address[0:1] != "\"" && address[len(address)-1:] != "\"" {
		return fmt.Errorf("gcode's address value could not be regconocide like integer, float number or string, %s", address)
	}

	return nil
}

func toInteger(value string) (int, error) {
	value64, err := strconv.ParseInt(value, 10, 32)
	return int(value64), err
}

func toFloat(value string) (float32, error) {
	value64, err := strconv.ParseFloat(value, 32)
	return float32(value64), err
}
