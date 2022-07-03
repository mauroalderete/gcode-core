package word_test

import (
	"fmt"

	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/word"
)

func ExampleNewWord() {

	w, err := word.NewWord("M")
	if err != nil {
		_ = fmt.Errorf("%s:", err.Error())
		return
	}

	fmt.Printf("%s\n", w.String())

	// Output: M
}
