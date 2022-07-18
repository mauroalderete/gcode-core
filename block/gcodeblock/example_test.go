package gcodeblock_test

import (
	"fmt"

	"github.com/mauroalderete/gcode-cli/block"
	"github.com/mauroalderete/gcode-cli/block/gcodeblock"
	"github.com/mauroalderete/gcode-cli/gcode"
	"github.com/mauroalderete/gcode-cli/gcode/addressablegcode"
)

func ExampleNew() {

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

func ExampleParse() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := gcodeblock.Parse(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("line is: %s\n", b.String())

	// Output: line is: N7 G1 X2.0 Y2.0 F3000.0
}

// func ExampleGcodeBlock_Checksum() {
// 	const source = "N7 G1 X2.0 Y2.0 F3000.0"

// 	b, err := gcodeblock.Parse(source, func(config block.BlockConfigurer) error {

// 		gc, err := addressablegcode.New[uint32]('*', 85)
// 		if err != nil {
// 			return err
// 		}

// 		config.SetChecksum(gc)

// 		return nil
// 	})
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	fmt.Printf("checksum is: %s\n", b.Checksum())

// 	// Output: checksum is: *85
// }

func ExampleGcodeBlock_Command() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := gcodeblock.Parse(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if b.Command() == nil {
		fmt.Println("command isn't available")
		return
	}

	fmt.Printf("command is: %s\n", b.Command().String())

	// Output: command is: G1
}

func ExampleGcodeBlock_Comment() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := gcodeblock.Parse(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("comment len is: %d\n", len(b.Comment()))

	// Output: comment len is: 0
}

func ExampleGcodeBlock_CalculateChecksum() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := gcodeblock.Parse(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	gc, err := b.CalculateChecksum()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	fmt.Printf("the gcode checksum is: %s\n", gc)

	// Output:
	// the gcode checksum is: *85
}

func ExampleGcodeBlock_UpdateChecksum() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := gcodeblock.Parse(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = b.UpdateChecksum()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	fmt.Printf("the block with checksum is: %s\n", b.ToLine("%l %c %p%k"))

	// Output:
	// the block with checksum is: N7 G1 X2.0 Y2.0 F3000.0*85
}

func ExampleGcodeBlock_VerifyChecksum() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := gcodeblock.Parse(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	checked, err := b.VerifyChecksum()

	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	fmt.Printf("the block is verified: %v\n", checked)

	// Output:
	// the block 'N7 G1 X2.0 Y2.0 F3000.0' hasn't check section
}

func ExampleGcodeBlock_LineNumber() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := gcodeblock.Parse(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if b.LineNumber() == nil {
		fmt.Println("line number isn't available")
		return
	}

	fmt.Printf("line number is: %v\n", b.LineNumber().Address())

	// Output: line number is: 7
}

func ExampleGcodeBlock_Parameters() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := gcodeblock.Parse(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if b.Parameters() == nil {
		fmt.Println("parameters aren't available")
		return
	}

	for i, p := range b.Parameters() {
		fmt.Printf("[%v]: %s\n", i, p.String())
	}

	// Output:
	// [0]: X2.0
	// [1]: Y2.0
	// [2]: F3000.0
}

func ExampleGcodeBlock_String() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := gcodeblock.Parse(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("line is: %s\n", b.String())

	// Output: line is: N7 G1 X2.0 Y2.0 F3000.0
}

func ExampleGcodeBlock_ToLine() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := gcodeblock.Parse(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("line is: %s\n", b.ToLine("%c %p"))

	// Output: line is: G1 X2.0 Y2.0 F3000.0
}

func ExampleGcodeBlock_ToLine_second() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := gcodeblock.Parse(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("line is: %s\n", b.ToLine("%l %c %p%k"))

	// Output: line is: N7 G1 X2.0 Y2.0 F3000.0
}

func ExampleGcodeBlock_ToLine_third() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := gcodeblock.Parse(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("line is: %s\n", b.ToLine("%m"))

	// Output: line is:
}
