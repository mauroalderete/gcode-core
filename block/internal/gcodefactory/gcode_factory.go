// gcodefactory implements gcode.GcoderFactory interface to create new instances Gcoder and AddressableGcoder using the implementations from unaddressablegcode and addressablegcode packages
//
// This package is only to internal use by gcodeblock package.
//
// Is used to establish a factory pattern by default when a caller required a gcodeblock instance
// but does not provide his self-gcodefactory.
//
// gcodefactory depends of unaddressablegcode and addressablegcode packages.
// His implementations are instanced by the GcodeFactory struct define in this package.
package gcodefactory

import (
	"github.com/mauroalderete/gcode-cli/gcode"
	"github.com/mauroalderete/gcode-cli/gcode/addressablegcode"
	"github.com/mauroalderete/gcode-cli/gcode/unaddressablegcode"
)

type GcodeFactory struct{}

// NewUnaddressableGcode is the constructor to instance a unaddressablegcode.Gcode struct.
//
// word represents the letter of the gcode command.
//
// If the word is an unknown symbol it returns nil with an error description.
func (g *GcodeFactory) NewUnaddressableGcode(word byte) (gcode.Gcoder, error) {
	ng, err := unaddressablegcode.New(word)
	if err != nil {
		return nil, err
	}

	return ng, nil
}

// NewAddressableGcodeUint32 is the constructor to instance a addressablegcode.Gcode[uint32] struct.
//
// word represents the letter of the gcode command.
// address is the value of the gcode command.
//
// If the word is an unknown symbol it returns nil with an error description.
func (g *GcodeFactory) NewAddressableGcodeUint32(word byte, address uint32) (gcode.AddressableGcoder[uint32], error) {

	ng, err := addressablegcode.New(word, address)
	if err != nil {
		return nil, err
	}

	return ng, nil
}

// NewAddressableGcodeInt32 is the constructor to instance a addressablegcode.Gcode[int32] struct.
//
// word represents the letter of the gcode command.
// address is the value of the gcode command.
//
// If the word is an unknown symbol it returns nil with an error description.
func (g *GcodeFactory) NewAddressableGcodeInt32(word byte, address int32) (gcode.AddressableGcoder[int32], error) {

	ng, err := addressablegcode.New(word, address)
	if err != nil {
		return nil, err
	}

	return ng, nil
}

// NewAddressableGcodeFloat32 is the constructor to instance a addressablegcode.Gcode[float32] struct.
//
// word represents the letter of the gcode command.
// address is the value of the gcode command.
//
// If the word is an unknown symbol it returns nil with an error description.
func (g *GcodeFactory) NewAddressableGcodeFloat32(word byte, address float32) (gcode.AddressableGcoder[float32], error) {

	ng, err := addressablegcode.New(word, address)
	if err != nil {
		return nil, err
	}

	return ng, nil
}

// NewAddressableGcodeString is the constructor to instance a addressablegcode.Gcode[string] struct.
//
// word represents the letter of the gcode command.
// address is the value of the gcode command.
//
// If the word is an unknown symbol it returns nil with an error description.
func (g *GcodeFactory) NewAddressableGcodeString(word byte, address string) (gcode.AddressableGcoder[string], error) {

	ng, err := addressablegcode.New(word, address)
	if err != nil {
		return nil, err
	}

	return ng, nil
}
