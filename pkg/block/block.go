package block

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/check"
	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/gcode"
)

type Block struct {
	lineNumber *gcode.GcodeAddressable[int32]
	command    gcode.Gcoder
	parameters []gcode.Gcoder
	check      check.Checker
	comment    *string
}

func (b *Block) String() string {
	return b.ToCommandComplete()
}

func (b *Block) LineNumber() gcode.GcodeAddressable[int32] {
	return *b.lineNumber
}
func (b *Block) Command() gcode.Gcoder {
	return b.command
}
func (b *Block) Parameters() []gcode.Gcoder {
	return b.parameters
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
	value += " " + b.command.String()

	for _, g := range b.parameters {
		value += " " + g.String()
	}

	return removeDuplicateSpaces(value)
}

func (b *Block) ToCommandWithCheck() string {
	value := ""

	value += " " + b.lineNumber.String()
	value += " " + b.command.String()

	for _, g := range b.parameters {
		value += " " + g.String()
	}

	checkGcode := b.check.Value()
	value += " " + checkGcode.String()

	return removeDuplicateSpaces(value)
}

func (b *Block) ToCommandComplete() string {
	value := ""

	if b.lineNumber != nil {
		value += " " + b.lineNumber.String()
	}
	if b.command != nil {
		value += " " + b.command.String()
	}

	if b.parameters != nil {
		for _, g := range b.parameters {
			value += " " + g.String()
		}
	}

	if b.check != nil {
		value += " " + b.check.Value().String()
	}

	if b.comment != nil {
		value += " " + *b.comment
	}

	return removeDuplicateSpaces(value)
}

func (b *Block) IsChecked() (bool, error) {
	checksum, err := check.NewCheck(check.CHECKSUM, b.ToCommand())
	if err != nil {
		return false, err
	}

	return checksum.Value() == b.check.Value(), nil
}

func Parse(s string) (*Block, error) {

	pblock := prepareSourceToParse(s)

	fmt.Printf("from %v, to %v\n", s, pblock)

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
		fmt.Printf("\n**********\nindex: %v\n", i)

		if i == 0 {
			pblock = pblock[1:]
			continue
		}

		var pgcode string = ""
		var pword byte
		var paddress string = ""

		if i <= -1 {
			fmt.Printf("ultimo!!!\n")
			pgcode = pblock
			pword = pgcode[0]
		} else {
			pgcode = pblock[:i]
			pword = pgcode[0]
		}

		if len(pgcode) > 1 {
			paddress = pgcode[1:]
		}

		fmt.Printf("pblock(%v): %v\n", len(pblock), pblock)
		fmt.Printf("pgcode(%v): %v\n", len(pgcode), pgcode)
		fmt.Printf("pword(%v): %v\n", 1, string(pword))
		fmt.Printf("paddress(%v): %v\n", len(paddress), paddress)
		// fmt.Printf("remain(%v): %v\n", len(pblock[i:]), pblock[i:])

		if len(pgcode) > 1 {
			fmt.Printf("parsing address from %v\n", pgcode[1:])
			var gca gcode.Gcoder
			//tiene address
			//es int?
			valueInt, err := strconv.ParseInt(pgcode[1:], 10, 32)
			if err == nil {
				gca, err = gcode.NewGcodeAddressable(pword, int32(valueInt))
				if err != nil {
					return nil, err
				}
				gcodes = append(gcodes, gca)
				fmt.Printf("add int32 ok: %v\n", gca.String())
				if i <= -1 {
					break loop
				}
				pblock = pblock[i:]
				continue
			}

			//es float?
			valueFloat, err := strconv.ParseFloat(pgcode[1:], 32)
			if err == nil {
				gca, err = gcode.NewGcodeAddressable(pword, float32(valueFloat))
				if err != nil {
					return nil, err
				}
				gcodes = append(gcodes, gca)
				fmt.Printf("add float32 ok: %v\n", gca.String())
				if i <= -1 {
					fmt.Printf("aaaaaaaaaaa\n")
					break loop
				}
				pblock = pblock[i:]
				continue
			}

			//asumo string
			gca, err = gcode.NewGcodeAddressable(pword, pgcode[1:])
			if err != nil {
				return nil, err
			}
			gcodes = append(gcodes, gca)
			fmt.Printf("add string ok: %v\n", gca.String())
			if i <= -1 {
				break loop
			}
			pblock = pblock[i:]
			continue
		} else {
			fmt.Printf("single!\n")
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
		b = &Block{
			lineNumber: nil,
			command:    gcodes[0],
			parameters: nil,
			check:      nil,
			comment:    nil,
		}
	} else {
		b = &Block{
			lineNumber: nil,
			command:    gcodes[0],
			parameters: gcodes[1:],
			check:      nil,
			comment:    nil,
		}
	}

	return b, nil
}

func removeDuplicateSpaces(s string) string {
	rx := regexp.MustCompile(`\s{2,}`)
	return rx.ReplaceAllString(s, " ")
}

func removeSpecialChars(s string) string {
	rx := regexp.MustCompile(`[\n\t\r]`)
	return rx.ReplaceAllString(s, " ")
}

func prepareSourceToParse(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	s = removeDuplicateSpaces(s)
	s = removeSpecialChars(s)

	return s
}
