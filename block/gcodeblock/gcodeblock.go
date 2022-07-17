package gcodeblock

import (
	"fmt"
	"hash"
	"regexp"
	"strconv"
	"strings"

	"github.com/mauroalderete/gcode-cli/block"
	"github.com/mauroalderete/gcode-cli/block/internal/gcodefactory"
	"github.com/mauroalderete/gcode-cli/gcode"
)

const (
	// BLOC_SEPARATOR defines a string used to separate the sections of the block when is exported as line string format
	BLOCK_SEPARATOR = " "
)

//#region block struct

// GcodeBlock struct represents a single gcode block.
//
// Stores data and gcode expressions of each section of the block.
//
// His methods allow export of the block as a line string format or check of the integrity.
//
// To be constructed using the Parse function from a line of the gcode file.
type GcodeBlock struct {
	// special gcode that store the value of the verification of the integrity of the block
	checksum gcode.AddresableGcoder[uint32]
	// first gcode expression and main significance of the block. Always is present.
	command gcode.Gcoder
	// expression attached at the block with some comment. Can be empty
	comment string
	// gcode factory
	gcodeFactory gcode.GcoderFactory
	// instance of the hash algorithm to handle the checksum
	hash hash.Hash
	// line number of the block. It can be null. Always has an int32 type address.
	lineNumber gcode.AddresableGcoder[uint32]
	// list of the rest of the gcode expression that adds information to the command. Can be empty.
	parameters []gcode.Gcoder
}

// String returns the block exported as single-line string format including check and comments section.
//
// It is the same invoke ToLineWithCheckAndComments method
func (b *GcodeBlock) String() string {
	return b.ToLineWithCheckAndComments()
}

// LineNumber returns a gcode addressable of the int32 type.
//
// Represent the line number of the block. It can be null. Always has an int32 type address.
func (b *GcodeBlock) LineNumber() gcode.AddresableGcoder[uint32] {
	return b.lineNumber
}

// Command returns a gcoder struct. Can be addressable or not.
//
// Represent the first gcode expression and main significance of the block. Always is present.
func (b *GcodeBlock) Command() gcode.Gcoder {
	return b.command
}

// Parameters return a list of gcoder structs. Each gcoder can be addressable or not.
//
// Parameters is a list of the rest of the gcode expression that adds information to the command. Can be empty.
func (b *GcodeBlock) Parameters() []gcode.Gcoder {
	return b.parameters
}

// Checksum returns a GcodeAddressable[uint32] if the line of the block contained, else returns nil.
//
// Exists two methods of checking: CRC and Checksum. Actually, only Checksum is supported.
func (b *GcodeBlock) Checksum() gcode.AddresableGcoder[uint32] {
	return b.checksum
}

// CalculateChecksum calculates a checksum from the block and returns a new GcodeAddressable[uint32] with the value computed.
func (b *GcodeBlock) CalculateChecksum() (gcode.AddresableGcoder[uint32], error) {

	b.hash.Reset()
	_, err := b.hash.Write([]byte(b.ToLine()))
	if err != nil {
		return nil, fmt.Errorf("failed to calculate hash to block %s: %w", b.ToLine(), err)
	}

	gc, err := b.gcodeFactory.NewAddressableGcodeUint32('*', uint32(b.hash.Sum(nil)[0]))
	if err != nil {
		return nil, fmt.Errorf("failed to create checksum gcode instance with hash %v: %w", uint32(b.hash.Sum(nil)[0]), err)
	}

	return gc, nil
}

// UpdateChecksum calculates a checksum from the block and stores him in as a new checksum gcode.
func (b *GcodeBlock) UpdateChecksum() error {

	gc, err := b.CalculateChecksum()
	if err != nil {
		return fmt.Errorf("failed update checksum of the block %s: %w", b, err)
	}

	b.checksum = gc

	return nil
}

// VerifyChecksum calculates a checksum and compare him with the checksum stored in the block, it returns true if both matches.
func (b *GcodeBlock) VerifyChecksum() (bool, error) {

	if b.checksum == nil {
		return false, fmt.Errorf("the block '%s' hasn't check section", b)
	}

	gc, err := b.CalculateChecksum()
	if err != nil {
		return false, fmt.Errorf("failed to calculate hash to the control of the checksum of the block %s: %w", b, err)
	}

	return b.checksum.Compare(gc), nil
}

// Comment returns the string with the comment of the block. Or nil if there isn't one.
//
// Is an expression attached at the block with some comment. Can be empty.
func (b *GcodeBlock) Comment() string {
	return b.comment
}

// ToLine export the block as a single-line string format with minimal data for can be executed in a machine
//
// The line generated depends on the content of the gcode line in the file line used to be parsed when the block was constructed.
//
// This method tries to export the line number of the block, the command and the parameters. Sometimes a block can haven't available the line number. In these cases, this is ignored.
//
// Sometimes a block could haven't a command (and any parameters) if the line used to parse the block didn't have either command originally.
func (b *GcodeBlock) ToLine() string {
	var values []string

	if b.lineNumber != nil {
		values = append(values, b.lineNumber.String())
	}

	if b.command != nil {
		values = append(values, b.command.String())
	}

	if b.parameters != nil {
		for _, g := range b.parameters {
			values = append(values, g.String())
		}
	}

	return strings.Join(values, BLOCK_SEPARATOR)
}

// ToLineWithCheck exports the block as a single-line string format adding the check section.
//
// Use ToLine output and attach the check section if is available
func (b *GcodeBlock) ToLineWithCheck() string {
	line := b.ToLine()

	if b.checksum != nil {
		line = strings.Join([]string{
			line,
			b.checksum.String()}, BLOCK_SEPARATOR)
	}

	return line
}

// ToLineWithCheckAndComments exports the block as a single-line string format adding the check section and comment section.
//
// Use ToLineWithCheck output and attach the comment section if is available.
//
// The string output is the expression more early at the original gcode line used to parse the block
func (b *GcodeBlock) ToLineWithCheckAndComments() string {
	line := b.ToLineWithCheck()

	if len(b.comment) > 0 {
		line = strings.Join([]string{line, b.comment}, BLOCK_SEPARATOR)
	}

	return line
}

//#endregion
//#region constructor

func New(command gcode.Gcoder, options ...block.BlockConfigurationCallbackable) (*GcodeBlock, error) {

	// command is required
	if command == nil {
		return nil, fmt.Errorf("command parameter is required")
	}

	// prepare a new GcodeBlock instance with some values by default
	gcodeBlock := &GcodeBlock{
		command: command,
	}

	// prepare an instance of the BlockConfigurer interface to store each configuration callback received
	configurator := &blockConfigurator{}

	// call each options to load configurations callback at configurator instance
	for _, option := range options {
		err := option(configurator)
		if err != nil {
			return nil, fmt.Errorf("failed to load configuration: %w", err)
		}
	}

	// apply each configuration callback that modify the new gcodeBlock instance
	for _, action := range configurator.configurationCallbacks {
		err := action(gcodeBlock)
		if err != nil {
			return nil, fmt.Errorf("failed to apply configuration: %w", err)
		}
	}

	// if is necesary, can validate that gcodeBlock is in valid state
	if gcodeBlock.hash != nil {
		if ok, err := gcodeBlock.VerifyChecksum(); !ok || err != nil {
			return gcodeBlock, fmt.Errorf("gcode block %s is invalid, checksum result is %v, error: %w ", gcodeBlock, ok, err)
		}
	}

	// can apply post-processing, like auto-calculate the checksum

	return gcodeBlock, nil
}

//#endregion

//#region package functions

// Parse return a new block instance using the data available in a single gcode line from gcode file
//
// Recive a string that must contain a gcode line valid.
// Try to extract each section from de block line to stores.
//
// Return an error if was a problem.
func Parse(s string, checksum hash.Hash, gcodeFactory gcode.GcoderFactory) (*GcodeBlock, error) {

	pblock := prepareSourceToParse(s)

	if gcodeFactory == nil {
		gcodeFactory = &gcodefactory.GcodeFactory{}
	}

	const separator = ' '

	var gcodes []gcode.Gcoder
	var i int = 0

loop:
	for {
		if i <= -1 {
			break loop
		}
		if len(pblock) == 0 {
			break loop
		}
		i = strings.IndexRune(pblock, separator)

		if i == 0 {
			pblock = pblock[1:]
			continue
		}

		var pgcode string = ""
		var pword byte
		var paddress string = ""

		if i <= -1 {
			pgcode = pblock
			pword = pgcode[0]
		} else {
			pgcode = pblock[:i]
			pword = pgcode[0]
		}

		if len(pgcode) > 1 {
			paddress = pgcode[1:]
		}

		if len(pgcode) > 1 {
			var gca gcode.Gcoder
			//tiene address
			//es int?
			valueInt, err := strconv.ParseInt(paddress, 10, 32)
			if err == nil {
				gca, err = gcodeFactory.NewAddressableGcodeInt32(pword, int32(valueInt))
				if err != nil {
					return nil, err
				}
				gcodes = append(gcodes, gca)
				if i <= -1 {
					break loop
				}
				pblock = pblock[i:]
				continue
			}

			//es float?
			valueFloat, err := strconv.ParseFloat(paddress, 32)
			if err == nil {
				gca, err = gcodeFactory.NewAddressableGcodeFloat32(pword, float32(valueFloat))
				if err != nil {
					return nil, err
				}
				gcodes = append(gcodes, gca)
				if i <= -1 {
					break loop
				}
				pblock = pblock[i:]
				continue
			}

			//asumo string
			gca, err = gcodeFactory.NewAddressableGcodeString(pword, paddress)
			if err != nil {
				return nil, err
			}
			gcodes = append(gcodes, gca)
			if i <= -1 {
				break loop
			}
			pblock = pblock[i:]
			continue
		} else {
			gc, err := gcodeFactory.NewUnaddressableGcode(pword)
			if err != nil {
				return nil, err
			}
			gcodes = append(gcodes, gc)
			if i <= -1 {
				break loop
			}
		}
		if i <= -1 {
			break loop
		}
		pblock = pblock[i:]
	}

	var b *GcodeBlock

	if len(gcodes) == 1 {

		ww := gcodes[0].Word()
		if ww == 'N' {

			//convert
			var ln gcode.AddresableGcoder[uint32]
			var ln2 gcode.AddresableGcoder[int32]
			var ok bool
			if ln2, ok = gcodes[0].(gcode.AddresableGcoder[int32]); !ok {
				return nil, fmt.Errorf("line number gcode found, but it was not possible to parse it")
			}

			ln, _ = gcodeFactory.NewAddressableGcodeUint32('N', uint32(ln2.Address()))

			b = &GcodeBlock{
				lineNumber:   ln,
				command:      nil,
				parameters:   nil,
				checksum:     nil,
				comment:      "",
				hash:         checksum,
				gcodeFactory: gcodeFactory,
			}

		} else {
			b = &GcodeBlock{
				lineNumber:   nil,
				command:      gcodes[0],
				parameters:   nil,
				checksum:     nil,
				comment:      "",
				hash:         checksum,
				gcodeFactory: gcodeFactory,
			}
		}
	} else {
		ww := gcodes[0].Word()
		if ww == 'N' {

			//convert
			var ln gcode.AddresableGcoder[uint32]
			var ln2 gcode.AddresableGcoder[int32]
			var ok bool
			if ln2, ok = gcodes[0].(gcode.AddresableGcoder[int32]); !ok {
				return nil, fmt.Errorf("line number gcode found, but it was not possible to parse it: %v %v", ok, gcodes[0])
			}

			ln, _ = gcodeFactory.NewAddressableGcodeUint32('N', uint32(ln2.Address()))

			b = &GcodeBlock{
				lineNumber:   ln,
				command:      gcodes[1],
				parameters:   gcodes[2:], //out of index warning
				checksum:     nil,
				comment:      "",
				hash:         checksum,
				gcodeFactory: gcodeFactory,
			}

		} else {
			b = &GcodeBlock{
				lineNumber:   nil,
				command:      gcodes[0],
				parameters:   gcodes[1:],
				checksum:     nil,
				comment:      "",
				hash:         checksum,
				gcodeFactory: gcodeFactory,
			}
		}
	}

	return b, nil
}

//#endregion
//#region private functions

// removeDuplicateSpaces remove all space char consecutive two or more times
func removeDuplicateSpaces(s string) string {
	rx := regexp.MustCompile(`\s{2,}`)
	return rx.ReplaceAllString(s, " ")
}

// removeSpecialChars remove all escape characters
func removeSpecialChars(s string) string {
	rx := regexp.MustCompile(`[\n\t\r]`)
	return rx.ReplaceAllString(s, " ")
}

// prepareSourceToParse modify a string to can be parsed for the Parse function
//
// It doesn't verify if s strings is a gcode line valid
func prepareSourceToParse(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	s = removeDuplicateSpaces(s)
	s = removeSpecialChars(s)

	return s
}

//#endregion