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
	"fmt"
	"strconv"
	"strings"

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

// Parse recives a string expression and tries convert a gcode.Gcoder object.
// source is a string expression of a gcode valid.
// if the expression is not recognited then returns an error.
// The orden to evaluate is N or checksum gcode first, string gcode second, nexto the float gcode and int gcode to end.
func (g *GcodeFactory) Parse(source string) (gcode.Gcoder, error) {

	if source == "" {
		return nil, fmt.Errorf("it is not possible to parse an empty string")
	}

	var gcode gcode.Gcoder
	var err error

	// is an unaddressable gcode
	if len(source) == 1 {

		gcode, err = g.NewUnaddressableGcode(source[0])
		if err != nil {
			return nil, fmt.Errorf("failed to parse %s, error to instance a new unaddresable gcode: %w", source, err)
		}

		return gcode, nil
	}

	// contains a linenumber or checksum gcode
	if source[0] == 'N' || source[0] == '*' {

		val, err := strconv.ParseInt(source[1:], 10, 32)
		if err != nil {
			return nil, fmt.Errorf("failed to try parse uint32 value from %s gcode: %w", source, err)
		}

		if val < 0 {
			return nil, fmt.Errorf("failed to try parse %d to uint32 value, it must be positive", val)
		}

		gcode, err = g.NewAddressableGcodeUint32(source[0], uint32(val))
		if err != nil {
			return nil, fmt.Errorf("try generate uint32 gcode from %s: %w", source, err)
		}

		return gcode, nil
	}

	// contains a string address
	if strings.Contains(source, "\"") {

		gcode, err = g.NewAddressableGcodeString(source[0], source[1:])
		if err != nil {
			return nil, fmt.Errorf("failed to parse %s, error to instance a new string addressable gcode: %w", source, err)
		}

		return gcode, nil
	}

	// contains a float address
	if strings.Contains(source, ".") {

		val, err := strconv.ParseFloat(source[1:], 32)
		if err != nil {
			return nil, fmt.Errorf("failed to parse %s, error to try get float address: %w", source, err)
		}

		gcode, err = g.NewAddressableGcodeFloat32(source[0], float32(val))
		if err != nil {
			return nil, fmt.Errorf("failed to parse %s, error to instance a new float32 addressable gcode: %w", source, err)
		}

		return gcode, nil
	}

	val, err := strconv.ParseInt(source[1:], 10, 32)
	if err != nil {
		return nil, fmt.Errorf("failed to try parse int value from %s gcode: %w", source, err)
	}
	gcode, err = g.NewAddressableGcodeInt32(source[0], int32(val))
	if err != nil {
		return nil, fmt.Errorf("try generate uint32 gcode from %s: %w", source, err)
	}

	return gcode, nil
}
