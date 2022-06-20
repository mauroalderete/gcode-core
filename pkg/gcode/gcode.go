package gcode

import (
	"errors"
	"fmt"
	"strings"

	gcode_address "github.com/mauroalderete/gcode-skew-transform-cli/pkg/gcode/address"
)

type Gcode struct {
	word    string
	address gcode_address.Address
}

type Gcoder interface {
	Word() string
	Address() gcode_address.Address
}

func (g *Gcode) String() string {
	return fmt.Sprintf("%s%s", g.word, g.address)
}

func (g *Gcode) Word() string {
	return g.word
}

func (g *Gcode) Address() gcode_address.Address {
	return g.address
}

func New(word string, address string) (*Gcode, error) {
	if len(word) != 1 {
		return nil, errors.New("gcode's word is invalid")
	}

	if strings.ContainsAny(address, " \t\n\r") {
		return nil, errors.New("gcode's address is invalid")
	}

	add, err := gcode_address.New(address)

	if err != nil {
		return nil, err
	}

	return &Gcode{
		word:    word,
		address: *add,
	}, nil
}

//equal
