// word package represents the letter of a gcode command that identifies a single action or a parameter.
//
// The words consist of a simple struct that implements a single value string. This value is a letter according to specification to gcode.
//
// Each word contains a single letter in uppercase. Each one of them can come accompanied by an address or not, or it needs other gcodes that give more information in the form of parameters.
//
// Exist, different classes the words, some are commands and others are used likes parameters for the commands.
package word

import (
	"errors"
)

// Word struct implement a word of a gcode
//
// Allow store a letter that mean the value of word.
type Word struct {
	value string
}

// String return the word value
func (w *Word) String() string {
	return w.value
}

// NewWord is a constructor to instance a word struct
//
// Recive a value of the word that represent a gcode command.
// If the value is a word valid then will return a pointer a new word struct.
// Else, will return a error message.
func NewWord(word string) (*Word, error) {

	err := isValid(word)
	if err != nil {
		return nil, err
	}

	newWord := Word{
		value: word,
	}

	return &newWord, nil
}

// isValid allow knowledge if a potential word value contain a value valid according a specification gcode
//
// The valid values are harcoding and they corresponds a [ReRap documentation]
//
// [ReRap documentation]: https://reprap.org/wiki/G-code
func isValid(word string) error {
	if len(word) != 1 {
		return errors.New("gcode's word has invalid format")
	}

	switch word[0] {
	case 'G', 'M', 'T', 'S', 'P', 'X', 'Y', 'Z', 'U', 'V', 'W', 'I', 'J', 'D', 'H', 'F', 'R', 'Q', 'E', 'N', '*':
		return nil
	}

	return errors.New("gcode's word has invalid value")
}
