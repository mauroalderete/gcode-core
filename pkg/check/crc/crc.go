package crc

import (
	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/address"
	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/gcode"
)

const CHECKSUM_WORD = 42

type CRC struct {
	value *gcode.Gcode
}

func (c *CRC) Value() gcode.Gcode {
	return *c.value
}

func NewCRC(source string) (*CRC, error) {

	var crc int32 = 0

	for _, v := range source {
		crc ^= int32(byte(v))
	}

	crc &= 0xff

	address, err := address.NewAddress(crc)

	if err != nil {
		return nil, err
	}

	gcode, err := gcode.NewGcode(CHECKSUM_WORD, address.String())

	if err != nil {
		return nil, err
	}

	return &CRC{
		value: gcode,
	}, nil
}
