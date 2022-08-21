// This file defines a Blockconfigurer as an object that implement block.BlockerConfigurer
// interface to allow the caller to configure the new blocks.
//
// Improve self-reference function to design options pattern providing the BlockConfigurer struct to set configs.
package gcodeblock

import (
	"fmt"
	"hash"

	"github.com/mauroalderete/gcode-core/gcode"
)

// optionalBlockPropertyCallbackable is a type that define the signature of the callbacks that implement logic to configure a new block instance.
type optionalBlockPropertyCallbackable func(*GcodeBlock) error

// blockConfigurator satisfy block.BlockConfigurer, contains the logic to create and store each optionalBlockPropertyCallbackable instance.
//
// It defines a slice of callbacks that will recive a new block reference that must be configured.
type blockConfigurator struct {
	configurationCallbacks []optionalBlockPropertyCallbackable
}

// SetGcodeFactory loads the gcodeFactory of the block with the input instance. Doesn't accept nil.
// If this method isn't called when a new block is created, by default will store a instance of Hash from the gcodefactory.GCodeFactory internal package.
func (bc *blockConfigurator) SetGcodeFactory(gcodeFactory gcode.GcoderFactory) error {

	if gcodeFactory == nil {
		return fmt.Errorf("failed set gcode factory, it mustn't be nil")
	}

	bc.configurationCallbacks = append(bc.configurationCallbacks, func(gb *GcodeBlock) error {
		gb.gcodeFactory = gcodeFactory
		return nil
	})

	return nil
}

// SetLineNumber loads the linenumber gcode of the block with the input instance. Doesn't accept nil.
// If this method isn't called when a new block is created, by default will to be nil.
func (bc *blockConfigurator) SetLineNumber(lineNumber gcode.AddressableGcoder[uint32]) error {

	if lineNumber == nil {
		return fmt.Errorf("failed set gcode 'line number', it mustn't be nil")
	}

	bc.configurationCallbacks = append(bc.configurationCallbacks, func(gb *GcodeBlock) error {
		gb.lineNumber = lineNumber
		return nil
	})

	return nil
}

// SetParameters loads the parameters gcode slice of the block with the input instance.
// It doesn't accept nil, but It accept a empty slice.
// If this method isn't called when a new block is created, by default will to be nil.
func (bc *blockConfigurator) SetParameters(parameters []gcode.Gcoder) error {

	if parameters == nil {
		return fmt.Errorf("failed set parameters at block, it mustn't be nil")
	}

	bc.configurationCallbacks = append(bc.configurationCallbacks, func(gb *GcodeBlock) error {
		gb.parameters = parameters
		return nil
	})

	return nil
}

// SetChecksum loads a checksum gcode of the block with the input instance. Doesn't accept nil.
// If this method isn't called when a new block is created, by default will to be nil.
func (bc *blockConfigurator) SetChecksum(checksum gcode.AddressableGcoder[uint32]) error {

	if checksum == nil {
		return fmt.Errorf("failed set checksum generator, it mustn't be nil")
	}

	bc.configurationCallbacks = append(bc.configurationCallbacks, func(gb *GcodeBlock) error {
		gb.checksum = checksum
		return nil
	})

	return nil
}

// SetHash loads the an hash algoritgh instance of the block with the input instanced. Doesn't accept nil.
// It require that the gcodeFactory is loaded in the block previously.
// If this method isn't called when a new block is created, by default will store a instance of Hash from the checksum package.
func (bc *blockConfigurator) SetHash(hash hash.Hash) error {

	if hash == nil {
		return fmt.Errorf("failed set checksum generator, it mustn't be nil")
	}

	bc.configurationCallbacks = append(bc.configurationCallbacks, func(gb *GcodeBlock) error {
		if gb.gcodeFactory == nil {
			return fmt.Errorf("failed to config ChecksumGenerator because depende of the gcodeFactory instanced and currently it is nil")
		}
		gb.hash = hash
		return nil
	})

	return nil
}

// SetComment store the block comments. It accept an empty string.
// If this method isn't called when a new block is created, by default is an empty string.
func (bc *blockConfigurator) SetComment(comment string) error {

	bc.configurationCallbacks = append(bc.configurationCallbacks, func(gb *GcodeBlock) error {
		gb.comment = comment
		return nil
	})

	return nil
}
