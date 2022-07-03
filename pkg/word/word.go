package word

import (
	"errors"
)

type Word struct {
	value string
}

func (w *Word) String() string {
	return w.value
}

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
