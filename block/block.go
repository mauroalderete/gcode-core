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
// A block struct be constructed from the Parse function that requires a line of the gcode file with a gcode block.
//
// The block struct can be export in different formats and can be verified if necessary.
package block

import (
	"fmt"
	"hash"

	"github.com/mauroalderete/gcode-cli/gcode"
)

type Blocker interface {
	fmt.Stringer

	CalculateChecksum() (gcode.AddresableGcoder[uint32], error)
	Checksum() gcode.AddresableGcoder[uint32]
	Command() gcode.Gcoder
	Comment() string
	LineNumber() gcode.AddresableGcoder[uint32]
	Parameters() []gcode.Gcoder
	ToLine() string
	ToLineWithCheck() string
	ToLineWithCheckAndComments() string
	UpdateChecksum() error
	VerifyChecksum() (bool, error)
}

// BlockConfigurer contain the minimal aspects configurables that can invoked whe a caller invoke to package function New()
type BlockConfigurer interface {
	SetGcodeFactory(gcodeFactory gcode.GcoderFactory) error
	SetLineNumber(lineNumber gcode.AddresableGcoder[uint32]) error
	SetParameters(parameters []gcode.Gcoder) error
	SetChecksum(checksum gcode.AddresableGcoder[uint32]) error
	SetChecksumGenerator(hash hash.Hash) error
	SetComment(comment string) error
}

// BlockConfigurationCallbackable is the signature of the callbacks that the package function New() waiting receives to configure the new block instance.
type BlockConfigurationCallbackable func(config BlockConfigurer) error

type BlockerFactory interface {
	New(command gcode.Gcoder, options ...BlockConfigurationCallbackable) (*Blocker, error)
	Parse(source string, options ...BlockConfigurationCallbackable) (*Blocker, error)
}
