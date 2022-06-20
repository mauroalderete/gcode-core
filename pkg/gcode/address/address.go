package address

import (
	"errors"
	"fmt"
	"strconv"
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
	value64, err := strconv.ParseInt(a.value, 10, 32)
	return int(value64), err
}

func (a *Address) ToFloat() (float32, error) {
	value64, err := strconv.ParseFloat(a.value, 32)
	return float32(value64), err
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

func NewAddress[A AddressType](address A) (*Address, error) {

	var addressValue string

	switch a := any(&address).(type) {
	case *string:
		addressValue = *a
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

	newAddress := Address{
		value: addressValue,
	}

	return &newAddress, nil
}
