package unaddressablegcode

import (
	"fmt"

	"github.com/mauroalderete/gcode-cli/gcode"
)

//#region interfaces

//#endregion
//#region gcode struct

// Gcode struct model a gcode expression with a stand-alone word.
//
// Allow model a gcode that does not contain an address. For this, stores a word.
type Gcode struct {
	word byte
}

// Compare allows checking if the current entity is equal to a Gcoder input
//
// This method is executed when to be called from a Gcode instance or a Gcoder instance that contains a reference to a Gcode object.
// If the input Gcoder does not implement some gcode.Addressable data type then it returns false
func (g *Gcode) Compare(gcode gcode.Gcoder) bool {
	if gcode == nil {
		return false
	}

	if gcode.HasAddress() {
		return false
	}

	return g.word == gcode.Word()
}

// HasAddress indicate if the gcode contains or not an address.
//
// This method is executed when to be called from a Gcode instance or a Gcoder instance that contains a reference to a Gcode object.
// In both cases, it always returns false.
func (g *Gcode) HasAddress() bool {
	return false
}

// String return Gcode formatted
func (g *Gcode) String() string {
	return string(g.word)
}

// Word return a copy of the word struct in the gcode
func (g *Gcode) Word() byte {
	return g.word
}

//#endregion
//#region constructor

// NewGcode is the constructor to instance a Gcode struct that includes an address.
//
// Receive a word that represents the letter of the command and another value that represent the address of gcode.
//
// The value can be string, int32 or float 32 data type.
//
// In any case, this method will verify the format of both parameters and return nil with an error description if necessary.
func New(word byte) (*Gcode, error) {
	err := gcode.IsValidWord(word)
	if err != nil {
		return nil, fmt.Errorf("failed to create an gcode instance when trying to use %v word: %w", word, err)
	}

	return &Gcode{
		word: word,
	}, nil
}

//#endregion
