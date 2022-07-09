package check

import (
	"fmt"

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
				return nil, fmt.Errorf("failed to create a Checker instance of Cheksum type with string %s: %w", data, err)
			}
			c := Checker(ret)
			return c, nil
		}
	case CRC:
		{
		}
	}

	return nil, fmt.Errorf("failed to create a Checker instance, kind's check %v is not valid", kind)
}
