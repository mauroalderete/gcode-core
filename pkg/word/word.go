// word package represents the letter of a gcode command that identifies a single action or a parameter.
//
// The words consist of a simple struct that implements a single value string. This value is a letter according to specification to gcode.
//
// Each word contains a single letter in uppercase. Each one of them can come accompanied by an address or not, or it needs other gcodes that give more information in the form of parameters.
//
// Exist, different classes the words, some are commands and others are used likes parameters for the commands.
package word

import (
	"fmt"
)

// Word struct implement a word of a gcode
//
// Allow store a letter that mean the value of word.
type Word struct {
	value byte
}

type WordInvalidValueError struct {
	Value byte
}

func (e *WordInvalidValueError) Error() string {
	return fmt.Errorf("gcode's word has invalid value: %s", string(e.Value)).Error()
}

// String return the word value like string data type
func (w *Word) String() string {
	return string(w.value)
}

// Value return the word value like a byte data type
//
// This field is immutable.
//
// Any change on the value field implies that the gcode involved will become a new gcode totally different, with another significate.
//
// When you need to change the gcode, you can instantiate another new gcode with the word and address that you require.
func (w *Word) Value() byte {
	return w.value
}

// NewWord is a constructor to instance a word struct
//
// Receive a byte that represents the word within the gcode command.
// If the value is a word valid then will return a pointer a new word struct.
// Else, will return a error message.
func NewWord(word byte) (*Word, error) {

	err := isValid(word)
	if err != nil {
		return nil, err
	}

	newWord := Word{
		value: word,
	}

	return &newWord, nil
}

// isValid allow knowledge if a potential word value contains a value valid according to a specification gcode
//
// The valid values are hard coding and they correspond to a [ReRap documentation]
//
// [ReRap documentation]: https://reprap.org/wiki/G-code
func isValid(word byte) error {

	switch word {
	case 'G', 'M', 'T', 'S', 'P', 'X', 'Y', 'Z', 'U', 'V', 'W', 'I', 'J', 'D', 'H', 'F', 'R', 'Q', 'E', 'N', '*':
		return nil
	}

	return &WordInvalidValueError{Value: word}
}
