package gcode

import (
	"fmt"

	gcode_address "github.com/mauroalderete/gcode-skew-transform-cli/pkg/address"
	gcode_word "github.com/mauroalderete/gcode-skew-transform-cli/pkg/word"
)

type Gcode struct {
	word    *gcode_word.Word
	address *gcode_address.Address
}

func (g *Gcode) String() string {
	return fmt.Sprintf("%s%s", g.word, g.address)
}

func (g *Gcode) Word() gcode_word.Word {
	return *g.word
}

func (g *Gcode) Address() gcode_address.Address {
	return *g.address
}

func (g *Gcode) Compare(gcode fmt.Stringer) bool {
	return g.String() == gcode.String()
}

func NewGcode(word string, address string) (*Gcode, error) {
	wrd, err := gcode_word.NewWord(word)
	if err != nil {
		return nil, err
	}

	add, err := gcode_address.NewAddress(address)

	if err != nil {
		return nil, err
	}

	return &Gcode{
		word:    wrd,
		address: add,
	}, nil
}
