// block package contains structs and methods to manage a gcode block
//
// A gcode block is a single line in the gcode file that contains almost a gcode expression.
// It is an operation that a machine must execute.
//
// A block contains different sections that can be completed or not.
//
// - line number of the block. It can be null. Always has an int32 type address.
//
// - first gcode expression and main significance of the block. Always is present.
//
// - list of the rest of the gcode expression that adds information to the command. Can be empty.
//
// - special gcode that store the value of the verification of the integrity of the block
//
// - expression attached at the block with some comment. Can be empty
//
// This package allows storing the data that define a single gcode block.
package block

import (
	"fmt"
	"hash"

	"github.com/mauroalderete/gcode-cli/gcode"
)

// Blocker defines the minimal methods to handle each element that compose a block
type Blocker interface {
	fmt.Stringer

	CalculateChecksum() (gcode.AddressableGcoder[uint32], error)
	Checksum() gcode.AddressableGcoder[uint32]
	Command() gcode.Gcoder
	Comment() string
	LineNumber() gcode.AddressableGcoder[uint32]
	Parameters() []gcode.Gcoder
	ToLine(format string) string
	UpdateChecksum() error
	VerifyChecksum() (bool, error)
}

// BlockConfigurer contains the configurable options that define a block when is constructed.
// Allows loading of the elements that compose a block during the creation.
//
// A block is determined by his gcode command. The rest of the parts might are included or not.
// Some blocks will require parameters, others not.
type BlockConfigurer interface {
	// Set a gcode.GcodeFactory instances that the block will use to handle his internal gcode elements
	SetGcodeFactory(gcodeFactory gcode.GcoderFactory) error

	// Set a line number of the block
	SetLineNumber(lineNumber gcode.AddressableGcoder[uint32]) error

	// Set all parameters that compose the block
	SetParameters(parameters []gcode.Gcoder) error

	// Set the checksum value of the block
	SetChecksum(checksum gcode.AddressableGcoder[uint32]) error

	// Set the hash instance that implement the algorith to execute checksum
	SetHash(hash hash.Hash) error

	// Set the comments from the block
	SetComment(comment string) error
}

// BlockConfigurationCallbackable is the signature of the callbacks that the package function New() waiting receives to configure the new block instance.
//
// Each callback provide a BlockConfigurer instance that implement a set of methods to configure the new block instance.
type BlockConfigurationCallbackable func(config BlockConfigurer) error

// BlockerFactory define the methods to create new Block instances
type BlockerFactory interface {
	// New return a new block instance with the configurations wishes.
	//
	// command is a gcode with address or not that define the block command.
	// options are a series of configuration callbacks to allow set different aspects of the block.
	// each option provides a config object that can be used to load the values that define the block.
	New(command gcode.Gcoder, options ...BlockConfigurationCallbackable) (*Blocker, error)

	// Parse returns a new block instance with the configurations we wish, from a single block line.
	// The block line must contain the correct format, on the contrary, the parsing process will end with an error.
	//
	// source is the string line to parse.
	// options are a series of configuration callbacks to allow set different aspects of the block.
	// each option provides a config object that can be used to load the values that define the block.
	Parse(source string, options ...BlockConfigurationCallbackable) (*Blocker, error)
}
