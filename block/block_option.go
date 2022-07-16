package block

import (
	"fmt"
	"hash"

	"github.com/mauroalderete/gcode-cli/gcode"
)

// BlockConfigurerCallback is the signature of the callbacks that the package function New() waiting receives to configure the new block instance.
type BlockConfigurerCallback func(BlockConfigurer) error

// BlockConfigurer contain the minimal aspects configurables that can invoked whe a caller invoke to package function New()
type BlockConfigurer interface {
	Checksum(hash hash.Hash)
	GcodeFactory(gcodeFactory gcode.GcoderFactory)
	LineNumber(lineNumber gcode.AddresableGcoder[uint32])
}

// BlockOptionalPropertiesCallback is a type that define the signature of the callbacks that implement logic to configure a new block instance.
type BlockOptionalPropertiesCallback func(*Block) error

// BlockConfiguration implement a slice of callbacks that will recive a new block reference that must be configured.
type BlockConfiguration struct {
	configurationCallbacks []BlockOptionalPropertiesCallback
}

// Checksum recives a Hash instance and create a callback that stores him at new block instance.
func (bc *BlockConfiguration) Checksum(hash hash.Hash) {
	bc.configurationCallbacks = append(bc.configurationCallbacks, func(b *Block) error {
		if hash == nil {
			return fmt.Errorf("hash option should not be nil")
		}
		b.hash = hash
		return nil
	})
}

// GcodeFactory recives a GcoderFactory instance and create a callback that stores him at new block instance.
func (bc *BlockConfiguration) GcodeFactory(gcodeFactory gcode.GcoderFactory) {
	bc.configurationCallbacks = append(bc.configurationCallbacks, func(b *Block) error {
		if gcodeFactory == nil {
			return fmt.Errorf("gcodeFactory option should not be nil")
		}
		b.gcodeFactory = gcodeFactory
		return nil
	})
}

// LineNumber recives a lineNumber gcode instance and create a callback that stores him at new block instance.
func (bc *BlockConfiguration) LineNumber(lineNumber gcode.AddresableGcoder[uint32]) {
	fmt.Println("aaaaaaaaaaaaaa")
	bc.configurationCallbacks = append(bc.configurationCallbacks, func(b *Block) error {
		if lineNumber == nil {
			return fmt.Errorf("gcodeFactory option should not be nil")
		}
		b.lineNumber = lineNumber
		fmt.Println("final line number: ", lineNumber, b)
		return nil
	})
	fmt.Println("actions...", bc.configurationCallbacks)
}

// ***********************************
// extension of configurer with a new method to add block parameters.
// ***********************************

type BlockConfigurerParameters interface {
	Parameters(parameters []gcode.Gcoder)
}

type BlockConfigurationParameter struct {
	BlockConfiguration
}

func (bc *BlockConfigurationParameter) Parameters(parameters []gcode.Gcoder) {
	bc.configurationCallbacks = append(bc.configurationCallbacks, func(b *Block) error {
		if parameters == nil {
			return fmt.Errorf("gcodeFactory option should not be nil")
		}
		b.parameters = parameters
		return nil
	})
}
