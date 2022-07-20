// unaddressablegcode package implements gcode.Gcoder interface to model a gcode without address element.
//
// Define a Gcode struct that implement gcode.Gcoder interface.
// This struct contain a word field to store the word value.
//
// A "New" constructor method allow to instance new Gcode objects.
// This method use gcode.IsValidWord to validate the input before create any instance
package unaddressablegcode

import (
	"fmt"

	"github.com/mauroalderete/gcode-core/gcode"
)

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

// New is the constructor to instance a Gcode struct that does not include an address.
//
// Receive a word that represents the letter of the command.
//
// Return nil with an error description of something is bad.
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
