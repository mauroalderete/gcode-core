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
	"regexp"
	"strconv"
	"strings"

	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/checksum"
	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/gcode"
)

const (
	// BLOC_SEPARATOR defines a string used to separate the sections of the block when is exported as line string format
	BLOCK_SEPARATOR = " "
)

//#region block struct

// Block struct represents a single gcode block.
//
// Stores data and gcode expressions of each section of the block.
//
// His methods allow export of the block as a line string format or check of the integrity.
//
// To be constructed using the Parse function from a line of the gcode file.
type Block struct {
	// line number of the block. It can be null. Always has an int32 type address.
	lineNumber *gcode.GcodeAddressable[uint32]
	// first gcode expression and main significance of the block. Always is present.
	command gcode.Gcoder
	// list of the rest of the gcode expression that adds information to the command. Can be empty.
	parameters []gcode.Gcoder
	// special gcode that store the value of the verification of the integrity of the block
	checksum *gcode.GcodeAddressable[uint32]
	// expression attached at the block with some comment. Can be empty
	comment string
	// instance of the hash algorithm to handle the checksum
	hash hash.Hash
}

// String returns the block exported as single-line string format including check and comments section.
//
// It is the same invoke ToLineWithCheckAndComments method
func (b *Block) String() string {
	return b.ToLineWithCheckAndComments()
}

// LineNumber returns a gcode addressable of the int32 type.
//
// Represent the line number of the block. It can be null. Always has an int32 type address.
func (b *Block) LineNumber() *gcode.GcodeAddressable[uint32] {
	return b.lineNumber
}

// Command returns a gcoder struct. Can be addressable or not.
//
// Represent the first gcode expression and main significance of the block. Always is present.
func (b *Block) Command() gcode.Gcoder {
	return b.command
}

// Parameters return a list of gcoder structs. Each gcoder can be addressable or not.
//
// Parameters is a list of the rest of the gcode expression that adds information to the command. Can be empty.
func (b *Block) Parameters() []gcode.Gcoder {
	return b.parameters
}

// Checksum returns a GcodeAddressable[uint32] if the line of the block contained, else returns nil.
//
// Exists two methods of checking: CRC and Checksum. Actually, only Checksum is supported.
func (b *Block) Checksum() *gcode.GcodeAddressable[uint32] {
	return b.checksum
}

// CalculateChecksum calculates a checksum from the block and returns a new GcodeAddressable[uint32] with the value computed.
func (b *Block) CalculateChecksum() (*gcode.GcodeAddressable[uint32], error) {

	b.hash.Reset()
	_, err := b.hash.Write([]byte(b.ToLine()))
	if err != nil {
		return nil, fmt.Errorf("failed to calculate hash to block %s: %w", b.ToLine(), err)
	}

	gc, err := gcode.NewGcodeAddressable('*', uint32(b.hash.Sum(nil)[0]))
	if err != nil {
		return nil, fmt.Errorf("failed to create checksum gcode instance with hash %v: %w", uint32(b.hash.Sum(nil)[0]), err)
	}

	return gc, nil
}

// UpdateChecksum calculates a checksum from the block and stores him in as a new checksum gcode.
func (b *Block) UpdateChecksum() error {

	gc, err := b.CalculateChecksum()
	if err != nil {
		return fmt.Errorf("failed update checksum of the block %s: %w", b, err)
	}

	b.checksum = gc

	return nil
}

// VerifyChecksum calculates a checksum and compare him with the checksum stored in the block, it returns true if both matches.
func (b *Block) VerifyChecksum() (bool, error) {

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
func (b *Block) Comment() string {
	return b.comment
}

// ToLine export the block as a single-line string format with minimal data for can be executed in a machine
//
// The line generated depends on the content of the gcode line in the file line used to be parsed when the block was constructed.
//
// This method tries to export the line number of the block, the command and the parameters. Sometimes a block can haven't available the line number. In these cases, this is ignored.
//
// Sometimes a block could haven't a command (and any parameters) if the line used to parse the block didn't have either command originally.
func (b *Block) ToLine() string {
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
func (b *Block) ToLineWithCheck() string {
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
func (b *Block) ToLineWithCheckAndComments() string {
	line := b.ToLineWithCheck()

	if len(b.comment) > 0 {
		line = strings.Join([]string{line, b.comment}, BLOCK_SEPARATOR)
	}

	return line
}

//#endregion
//#region package functions

// Parse return a new block instance using the data available in a single gcode line from gcode file
//
// Recive a string that must contain a gcode line valid.
// Try to extract each section from de block line to stores.
//
// Return an error if was a problem.
func Parse(s string) (*Block, error) {

	pblock := prepareSourceToParse(s)

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
				gca, err = gcode.NewGcodeAddressable(pword, int32(valueInt))
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
				gca, err = gcode.NewGcodeAddressable(pword, float32(valueFloat))
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
			gca, err = gcode.NewGcodeAddressable(pword, paddress)
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
			gc, err := gcode.NewGcode(pword)
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

	var b *Block

	if len(gcodes) == 1 {

		ww := gcodes[0].Word()
		if ww.String() == "N" {

			//convert
			var ln *gcode.GcodeAddressable[uint32]
			var ln2 *gcode.GcodeAddressable[int32]
			var ok bool
			if ln2, ok = gcodes[0].(*gcode.GcodeAddressable[int32]); !ok {
				return nil, fmt.Errorf("line number gcode found, but it was not possible to parse it")
			}

			ln, _ = gcode.NewGcodeAddressable('N', uint32(ln2.Address().Value()))

			b = &Block{
				lineNumber: ln,
				command:    nil,
				parameters: nil,
				checksum:   nil,
				comment:    "",
				hash:       checksum.New(),
			}

		} else {
			b = &Block{
				lineNumber: nil,
				command:    gcodes[0],
				parameters: nil,
				checksum:   nil,
				comment:    "",
				hash:       checksum.New(),
			}
		}
	} else {
		ww := gcodes[0].Word()
		if ww.String() == "N" {

			//convert
			var ln *gcode.GcodeAddressable[uint32]
			var ln2 *gcode.GcodeAddressable[int32]
			var ok bool
			if ln2, ok = gcodes[0].(*gcode.GcodeAddressable[int32]); !ok {
				return nil, fmt.Errorf("line number gcode found, but it was not possible to parse it: %v %v", ok, gcodes[0])
			}

			ln, _ = gcode.NewGcodeAddressable('N', uint32(ln2.Address().Value()))

			b = &Block{
				lineNumber: ln,
				command:    gcodes[1],
				parameters: gcodes[2:], //out of index warning
				checksum:   nil,
				comment:    "",
				hash:       checksum.New(),
			}

		} else {
			b = &Block{
				lineNumber: nil,
				command:    gcodes[0],
				parameters: gcodes[1:],
				checksum:   nil,
				comment:    "",
				hash:       checksum.New(),
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
