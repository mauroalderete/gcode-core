// this file containts the word model. Represents the letter of a gcode command that identifies a single action or a parameter.
//
// The words consist of a simple struct that implements a single value string. This value is a letter according to specification to gcode.
//
// Each word contains a single letter in uppercase. Each one of them can come accompanied by an address or not, or it needs other gcodes that give more information in the form of parameters.
//
// Exist, different classes the words, some are commands and others are used likes parameters for the commands.
package gcode

import (
	"fmt"
)

//#region word struct

// Word struct implement a word of a gcode.
//
// It allows the storing of a letter that represent the value of word.
type Word struct {
	value byte
}

// String return the word value like string data type.
func (w Word) String() string {
	return string(w.value)
}

// Value return the word value like a byte data type.
//
// This field is immutable.
//
// Any change on the value field implies that the gcode involved will become a new gcode totally different, with another significate.
//
// When you need to change the gcode, you can instantiate another new gcode with the new word and the new address that you require.
func (w Word) Value() byte {
	return w.value
}

//#endregion

//#region constructors

// NewWord is a constructor to instance a word struct.
//
// Receive a byte that represents the word in the gcode command.
// If the value is a word valid then will return a pointer a new word struct.
// Else, it will return an error.
func NewWord(word byte) (*Word, error) {

	err := isValid(word)
	if err != nil {
		return nil, fmt.Errorf("failed the construction of a new word with the value %v: %w", word, err)
	}

	return &Word{
		value: word,
	}, nil
}

//#endregion

//#region private functions

// isValid allow knowledge if a potential word value contains a value valid according to a specification gcode.
//
// The set valid values are hard coding and they correspond to a [ReRap documentation].
//
// [ReRap documentation]: https://reprap.org/wiki/G-code
func isValid(word byte) error {

	switch word {
	case 'G', 'M', 'T', 'S', 'P', 'X', 'Y', 'Z', 'U', 'V', 'W', 'I', 'J', 'D', 'H', 'F', 'R', 'Q', 'E', 'N', '*':
		return nil
	}

	return fmt.Errorf("gcode's word has invalid value: %v", word)
}

//#endregion
