package word_test

import (
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
			expectedError string
		}{
			{'+', nil, "gcode's word has invalid value"},
			{'-', nil, "gcode's word has invalid value"},
			{' ', nil, "gcode's word has invalid value"},
			{'\n', nil, "gcode's word has invalid value"},
			{'$', nil, "gcode's word has invalid value"},
			{'"', nil, "gcode's word has invalid value"},
			{'0', nil, "gcode's word has invalid value"},
			{'9', nil, "gcode's word has invalid value"},
		}

		for _, c := range invalidCases {
			t.Run(fmt.Sprintf("test with %s", string(c.value)), func(t *testing.T) {
				w, err := word.NewWord(c.value)

				if w != c.expectedWord {
					t.Errorf("got %v, want nil", w)
				}

				if err == nil {
					t.Errorf("got nil, want '%s'", "gcode's word has invalid value")
				}

				if err.Error() != "gcode's word has invalid value" {
					t.Errorf("got %s, want '%s'", err.Error(), "gcode's word has invalid value")
				}
			})
		}
	})
}
