package gcode

import (
	"regexp"

	blockchecksum "github.com/mauroalderete/gcode-skew-transform-cli/pkg/block/checksum"
	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/gcode"
)

type Block struct {
	lineNumber *gcode.Gcode
	code       *gcode.Gcode
	parameters *[]gcode.Gcode
	checksum   *blockchecksum.Checksum
	comment    *string
}

type Blocker interface {
	LineNumber() gcode.Gcode
	Code() gcode.Gcode
	Parameters() []gcode.Gcode
	Checksum() blockchecksum.Checksum
	Comment() string

	IsChecksumValid() (bool, error)

	cleanOutputLine(string) string
}

type BlockerStringer interface {
	ToLineComplete() string
	ToCommand() string
	ToCommandWithChecksum() string
}

func (b *Block) LineNumber() gcode.Gcode {
	return *b.lineNumber
}
func (b *Block) Code() gcode.Gcode {
	return *b.code
}
func (b *Block) Parameters() []gcode.Gcode {
	return *b.parameters
}
func (b *Block) Checksum() blockchecksum.Checksum {
	return *b.checksum
}
func (b *Block) Comment() string {
	return *b.comment
}

func (b *Block) ToCommand() string {
	value := ""

	value += " " + b.lineNumber.String()
	value += " " + b.code.String()

	for _, g := range *b.parameters {
		value += " " + g.String()
	}

	return b.cleanOutputLine(value)
}

func (b *Block) ToCommandWithChecksum() string {
	value := ""

	value += " " + b.lineNumber.String()
	value += " " + b.code.String()

	for _, g := range *b.parameters {
		value += " " + g.String()
	}

	value += " " + b.checksum.String()

	return b.cleanOutputLine(value)
}

func (b *Block) ToLineComplete() string {
	value := ""

	value += " " + b.lineNumber.String()
	value += " " + b.code.String()

	for _, g := range *b.parameters {
		value += " " + g.String()
	}

	value += " " + b.checksum.String()
	value += " " + *b.comment

	return b.cleanOutputLine(value)
}

func (b *Block) String() string {
	return b.ToLineComplete()
}

func (b *Block) cleanOutputLine(line string) string {
	rx := regexp.MustCompile(`\s{2,}`)
	return rx.ReplaceAllString(line, " ")
}

func (b *Block) IsChecksumValid() (bool, error) {
	checksum, err := blockchecksum.GenerateChecksum(b.ToCommand())
	if err != nil {
		return false, err
	}

	return checksum.Checksum() == b.checksum.Checksum(), nil
}
