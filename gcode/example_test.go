package gcode_test

import (
	"fmt"
	"log"

	"github.com/mauroalderete/gcode-cli/gcode"
)

func ExampleIsValidWord() {

	const word byte = 'M'

	err := gcode.IsValidWord(word)
	if err != nil {
		log.Fatalf("word %s is invalid: %v", string(word), err)
	}

	fmt.Printf("word %s is valid", string(word))

	// Output: word M is valid
}

func ExampleIsValidWord_second() {

	const word byte = ';'

	err := gcode.IsValidWord(word)
	if err != nil {
		fmt.Printf("word %s is invalid: %v", string(word), err)
		return
	}

	fmt.Printf("word %s is valid", string(word))

	// Output: word ; is invalid: gcode's word has invalid value: 59
}
