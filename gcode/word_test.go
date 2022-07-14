package word_test

import (
	"fmt"
	"testing"

	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/word"
)

//region unit tests

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

		for i, tc := range validCases {
			t.Run(fmt.Sprintf("(%d)", i), func(t *testing.T) {
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
			value byte
		}{
			{'+'},
			{'-'},
			{' '},
			{'\n'},
			{'$'},
			{'"'},
			{'0'},
			{'9'},
		}

		for i, tc := range invalidCases {
			t.Run(fmt.Sprintf("(%d)", i), func(t *testing.T) {
				w, err := word.NewWord(tc.value)

				if w != nil {
					t.Errorf("got %v word, want word nil", w)
				}

				if err == nil {
					t.Errorf("got nil error, want error not nil")
				}
			})
		}
	})
}

//#endregion
//#region examples

func ExampleNewWord() {

	w, err := word.NewWord('M')
	if err != nil {
		_ = fmt.Errorf("%s:", err.Error())
		return
	}

	fmt.Printf("%s\n", w)

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

//#endregion
