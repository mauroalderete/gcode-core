// checksum package contain a Checksum struct that implement the Check interface
package checksum

import (
	"fmt"

	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/address"
	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/gcode"
)

const (
	// CHECKSUM_WORD contains the word value that identifies a check gcode
	CHECKSUM_WORD = byte('*')
)

// Checksum is an struct that implement Checker interface
type Checksum struct {
	value gcode.Gcoder
}

func (c *Checksum) Value() gcode.Gcoder {
	return c.value
}

// NewChecksum try to create a Checksum instance from a string that must be parsed
func NewChecksum(source string) (*Checksum, error) {

	var checksum int32 = 0

	for _, v := range source {
		checksum ^= int32(byte(v))
	}

	checksum &= 0xff

	address, err := address.NewAddress(checksum)
	if err != nil {
		return nil, fmt.Errorf("checksum constructor failed to generate a checksum valid %v from source %v: %w", checksum, source, err)
	}

	gcode, err := gcode.NewGcodeAddressable(CHECKSUM_WORD, address.Value())
	if err != nil {
		return nil, fmt.Errorf("failed to create a Checksum instance from source %v: %w", source, err)
	}

	return &Checksum{
		value: gcode,
	}, nil
}
