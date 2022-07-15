package gcodefactory

import (
	"github.com/mauroalderete/gcode-cli/gcode"
	"github.com/mauroalderete/gcode-cli/gcode/addressablegcode"
	"github.com/mauroalderete/gcode-cli/gcode/unaddressablegcode"
)

type GcodeFactory struct{}

// NewGcode is the constructor to instance a Gcode struct.
//
// Receive a word that represents the letter of the command of a gcode.
//
// If the word is an unknown symbol it returns nil with an error description.
func (g *GcodeFactory) NewUnaddressableGcode(word byte) (gcode.Gcoder, error) {
	ng, err := unaddressablegcode.New(word)
	if err != nil {
		return nil, err
	}

	return ng, nil
}

func (g *GcodeFactory) NewAddressableGcodeUint32(word byte, address uint32) (gcode.AddresableGcoder[uint32], error) {

	ng, err := addressablegcode.New(word, address)
	if err != nil {
		return nil, err
	}

	return ng, nil
}

func (g *GcodeFactory) NewAddressableGcodeInt32(word byte, address int32) (gcode.AddresableGcoder[int32], error) {

	ng, err := addressablegcode.New(word, address)
	if err != nil {
		return nil, err
	}

	return ng, nil
}

func (g *GcodeFactory) NewAddressableGcodeFloat32(word byte, address float32) (gcode.AddresableGcoder[float32], error) {

	ng, err := addressablegcode.New(word, address)
	if err != nil {
		return nil, err
	}

	return ng, nil
}

func (g *GcodeFactory) NewAddressableGcodeString(word byte, address string) (gcode.AddresableGcoder[string], error) {

	ng, err := addressablegcode.New(word, address)
	if err != nil {
		return nil, err
	}

	return ng, nil
}
