package word_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/word"
)

func ExampleNewWord() {

	w, err := word.NewWord('M')
	if err != nil {
		_ = fmt.Errorf("%s:", err.Error())
		return
	}

	fmt.Printf("%s\n", w.String())

	// Output: M
}

func ExampleWord_String() {
	w, err := word.NewWord('M')
	if err != nil {
		_ = fmt.Errorf("%s:", err.Error())
		return
	}

	fmt.Printf("%s\n", w.String())

	// Output: M
}

func ExampleWord_Value() {
	w, err := word.NewWord('M')
	if err != nil {
		_ = fmt.Errorf("%s:", err.Error())
		return
	}

	fmt.Printf("%v\n", w.Value())

	// Output: 77
}

func TestNewWord(t *testing.T) {
	t.Run("words valids", func(t *testing.T) {
		validCases := []struct {
			value          byte
			expectedValue  byte
			expectedString string
		}{
			{'G', 'G', "G"},
			{'M', 'M', "M"},
			{'T', 'T', "T"},
			{'S', 'S', "S"},
			{'*', '*', "*"},
		}

		for _, tc := range validCases {
			t.Run(fmt.Sprintf("test with %s", string(tc.value)), func(t *testing.T) {
				w, err := word.NewWord(tc.value)
				if err != nil {
					t.Errorf("got %v, want %v and %v", err, tc.expectedValue, tc.expectedString)
				}

				if w.Value() != tc.expectedValue {
					t.Errorf("got %v, want byte %v", w.Value(), tc.expectedValue)
				}

				if w.String() != tc.expectedString {
					t.Errorf("got %v, want string %v", w.Value(), tc.expectedString)
				}
			})
		}
	})

	t.Run("words invalids", func(t *testing.T) {
		invalidCases := []struct {
			value         byte
			expectedWord  *word.Word
			expectedError word.WordInvalidValueError
		}{
			{'+', nil, word.WordInvalidValueError{Value: '+'}},
			{'-', nil, word.WordInvalidValueError{Value: '-'}},
			{' ', nil, word.WordInvalidValueError{Value: ' '}},
			{'\n', nil, word.WordInvalidValueError{Value: '\n'}},
			{'$', nil, word.WordInvalidValueError{Value: '$'}},
			{'"', nil, word.WordInvalidValueError{Value: '"'}},
			{'0', nil, word.WordInvalidValueError{Value: '0'}},
			{'9', nil, word.WordInvalidValueError{Value: '9'}},
		}

		for _, tc := range invalidCases {
			t.Run(fmt.Sprintf("test with %s", string(tc.value)), func(t *testing.T) {
				w, err := word.NewWord(tc.value)

				if w != tc.expectedWord {
					t.Errorf("got %v, want nil", w)
				}

				if err == nil {
					t.Errorf("got nil, want '%s'", tc.expectedError.Error())
				}

				if errors.Is(err, &tc.expectedError) {
					t.Errorf("got %s, want '%s'", err.Error(), tc.expectedError.Error())
				}
			})
		}
	})
}

func TestWordInvalidValueError_Error(t *testing.T) {

	const value = ';'
	const expected = "gcode's word has invalid value: " + string(value)

	t.Run(fmt.Sprintf("with %s", string(value)), func(t *testing.T) {
		err := word.WordInvalidValueError{Value: ';'}
		if err.Error() != expected {
			t.Errorf("got %s, want %s", err.Error(), expected)
		}
	})

}
