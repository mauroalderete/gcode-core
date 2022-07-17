package gcodeblock

import (
	"fmt"
	"hash"

	"github.com/mauroalderete/gcode-cli/gcode"
)

// BlockOptionalPropertiesCallback is a type that define the signature of the callbacks that implement logic to configure a new block instance.
type optionalBlockPropertyCallbackable func(*GcodeBlock) error

type BlockConfigurer interface {
	SetGcodeFactory(gcodeFactory gcode.GcoderFactory) error
	SetLineNumber(lineNumber gcode.AddresableGcoder[uint32]) error
	SetParameters(parameters []gcode.Gcoder) error
	UseChecksum(hash hash.Hash) error
	UseComment() error
	UseLineNumber() error
}

// BlockConfiguration implement a slice of callbacks that will recive a new block reference that must be configured.
type blockConfigurator struct {
	configurationCallbacks []optionalBlockPropertyCallbackable
}

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

func (bc *blockConfigurator) SetLineNumber(lineNumber gcode.AddresableGcoder[uint32]) error {

	if lineNumber == nil {
		return fmt.Errorf("failed set gcode 'line number', it mustn't be nil")
	}

	bc.configurationCallbacks = append(bc.configurationCallbacks, func(gb *GcodeBlock) error {
		gb.lineNumber = lineNumber
		return nil
	})

	return nil
}

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

func (bc *blockConfigurator) SetChecksum(checksum gcode.AddresableGcoder[uint32]) error {

	if checksum == nil {
		return fmt.Errorf("failed set checksum generator, it mustn't be nil")
	}

	bc.configurationCallbacks = append(bc.configurationCallbacks, func(gb *GcodeBlock) error {
		gb.checksum = checksum
		return nil
	})

	return nil
}

func (bc *blockConfigurator) SetChecksumGenerator(hash hash.Hash) error {

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

func (bc *blockConfigurator) SetComment(comment string) error {

	bc.configurationCallbacks = append(bc.configurationCallbacks, func(gb *GcodeBlock) error {
		gb.comment = comment
		return nil
	})

	return nil
}
