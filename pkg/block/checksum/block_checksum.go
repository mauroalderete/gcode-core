package checksum

import (
	"strconv"

	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/gcode"
	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/gcode/address"
)

const CHECKSUM_WORD = 42

type Checksum struct {
	checksum gcode.Gcode
}

type Checksumer interface {
	Checksum() gcode.Gcode
}

func (c *Checksum) Checksum() gcode.Gcode {
	return c.checksum
}

func (c *Checksum) String() string {
	return c.checksum.String()
}

func New(checksum int32) (*Checksum, error) {

	address, err := address.New(checksum)

	if err != nil {
		return nil, err
	}

	gcode, err := gcode.New(strconv.QuoteRune(CHECKSUM_WORD), address.String())

	if err != nil {
		return nil, err
	}

	return &Checksum{
		checksum: *gcode,
	}, nil
}

func GenerateChecksum(value string) (*Checksum, error) {

	var cs int32 = 0

	for _, v := range value {
		cs ^= int32(byte(v))
	}

	cs &= 0xff

	block_checksum, err := New(cs)

	if err != nil {
		return nil, err
	}

	return block_checksum, nil
}
