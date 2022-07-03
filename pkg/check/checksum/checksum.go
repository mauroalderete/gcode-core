package checksum

import (
	"strconv"

	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/address"
	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/gcode"
)

const CHECKSUM_WORD = 42

type Checksum struct {
	value *gcode.Gcode
}

func (c *Checksum) Value() gcode.Gcode {
	return *c.value
}

func NewChecksum(source string) (*Checksum, error) {

	var checksum int32 = 0

	for _, v := range source {
		checksum ^= int32(byte(v))
	}

	checksum &= 0xff

	address, err := address.NewAddress(checksum)

	if err != nil {
		return nil, err
	}

	gcode, err := gcode.NewGcode(strconv.QuoteRune(CHECKSUM_WORD), address.String())

	if err != nil {
		return nil, err
	}

	return &Checksum{
		value: gcode,
	}, nil
}
