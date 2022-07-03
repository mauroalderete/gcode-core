package gcode

import (
	"regexp"

	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/check"
	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/gcode"
)

type Block struct {
	lineNumber *gcode.Gcode
	code       *gcode.Gcode
	parameters *[]gcode.Gcode
	check      check.Checker
	comment    *string
}

func (b *Block) String() string {
	return b.ToLineComplete()
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
func (b *Block) Checksum() check.Checker {
	return b.check
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

	return trimStringOut(value)
}

func (b *Block) ToCommandWithChecksum() string {
	value := ""

	value += " " + b.lineNumber.String()
	value += " " + b.code.String()

	for _, g := range *b.parameters {
		value += " " + g.String()
	}

	checkGcode := b.check.Value()
	value += " " + checkGcode.String()

	return trimStringOut(value)
}

func (b *Block) ToLineComplete() string {
	value := ""

	value += " " + b.lineNumber.String()
	value += " " + b.code.String()

	for _, g := range *b.parameters {
		value += " " + g.String()
	}

	checkGcode := b.check.Value()
	value += " " + checkGcode.String()
	value += " " + *b.comment

	return trimStringOut(value)
}

func trimStringOut(line string) string {
	rx := regexp.MustCompile(`\s{2,}`)
	return rx.ReplaceAllString(line, " ")
}

func (b *Block) IsChecksumValid() (bool, error) {
	checksum, err := check.NewCheck(check.CHECKSUM, b.ToCommand())
	if err != nil {
		return false, err
	}

	return checksum.Value() == b.check.Value(), nil
}
