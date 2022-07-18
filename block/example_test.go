package block_test

import (
	"fmt"

	"github.com/mauroalderete/gcode-cli/block"
	"github.com/mauroalderete/gcode-cli/block/gcodeblock"
	"github.com/mauroalderete/gcode-cli/gcode"
	"github.com/mauroalderete/gcode-cli/gcode/addressablegcode"
)

func ExampleBlocker() {

	// create a new command
	command, err := addressablegcode.New[int32]('G', 1)
	if err != nil {
		fmt.Printf("got error not nil, want error nil: %v", err)
		return
	}

	// try create a new block with command G1
	b, err := gcodeblock.New(command, func(config block.BlockConfigurer) error {

		// set line number N7
		lineNumber, err := addressablegcode.New[uint32]('N', 7)
		if err != nil {
			return fmt.Errorf("got error not nil, want error nil: %v", err)
		}
		config.SetLineNumber(lineNumber)

		// set all parameters X2 Y2 F3000
		p1, err := addressablegcode.New[float32]('X', 2)
		if err != nil {
			return fmt.Errorf("got error not nil, want error nil: %v", err)
		}
		p2, err := addressablegcode.New[float32]('Y', 2)
		if err != nil {
			return fmt.Errorf("got error not nil, want error nil: %v", err)
		}
		p3, err := addressablegcode.New[float32]('F', 3000)
		if err != nil {
			return fmt.Errorf("got error not nil, want error nil: %v", err)
		}
		params := []gcode.Gcoder{p1, p2, p3}
		config.SetParameters(params)

		// set a comment
		config.SetComment(";lorem ipsum")

		return nil
	})

	// calculate the checksum and update the block with the new value
	err = b.UpdateChecksum()
	if err != nil {
		fmt.Printf("failed to update: %v", err)
		return
	}

	fmt.Println(b.ToLine("%l %c %p%k %m"))

	// Output: N7 G1 X2.0 Y2.0 F3000.0*85 ;lorem ipsum
}
