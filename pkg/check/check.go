package check

import (
	"errors"

	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/check/checksum"
	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/gcode"
)

type CheckKind int32

const (
	CHECKSUM CheckKind = iota
	CRC
)

type Checker interface {
	Value() gcode.Gcoder
}

func NewCheck(kind CheckKind, data string) (Checker, error) {

	switch kind {
	case CHECKSUM:
		{
			ret, err := checksum.NewChecksum(data)
			if err != nil {
				return nil, err
			}
			c := Checker(ret)
			return c, nil
		}
	case CRC:
		{
		}
	}

	return nil, errors.New("kind's check is not valid")
}
