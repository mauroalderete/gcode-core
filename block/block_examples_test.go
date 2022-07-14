package block_test

import (
	"fmt"

	"github.com/mauroalderete/gcode-cli/block"
)

func ExampleParse() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := block.Parse(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("line is: %s\n", b.ToLineWithCheckAndComments())

	// Output: line is: N7 G1 X2.0 Y2.0 F3000.0
}

func ExampleBlock_Checksum() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := block.Parse(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if b.Checksum() == nil {
		fmt.Println("checksum isn't available")
		return
	}

	fmt.Printf("checksum is: %s\n", b.Checksum())

	// Output: checksum isn't available
}

func ExampleBlock_Command() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := block.Parse(source)
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

func ExampleBlock_Comment() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := block.Parse(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("comment len is: %d\n", len(b.Comment()))

	// Output: comment len is: 0
}

func ExampleBlock_CalculateChecksum() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := block.Parse(source)
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

func ExampleBlock_UpdateChecksum() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := block.Parse(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = b.UpdateChecksum()
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	fmt.Printf("the block with checksum is: %s\n", b)

	// Output:
	// the block with checksum is: N7 G1 X2.0 Y2.0 F3000.0 *85
}

func ExampleBlock_VerifyChecksum() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := block.Parse(source)
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

func ExampleBlock_LineNumber() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := block.Parse(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if b.LineNumber() == nil {
		fmt.Println("line number isn't available")
		return
	}

	fmt.Printf("line number is: %v\n", b.LineNumber().Address().Value())

	// Output: line number is: 7
}

func ExampleBlock_Parameters() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := block.Parse(source)
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

func ExampleBlock_String() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := block.Parse(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("line is: %s\n", b.String())

	// Output: line is: N7 G1 X2.0 Y2.0 F3000.0
}

func ExampleBlock_ToLine() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := block.Parse(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("line is: %s\n", b.ToLine())

	// Output: line is: N7 G1 X2.0 Y2.0 F3000.0
}

func ExampleBlock_ToLineWithCheck() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := block.Parse(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("line is: %s\n", b.ToLineWithCheck())

	// Output: line is: N7 G1 X2.0 Y2.0 F3000.0
}

func ExampleBlock_ToLineWithCheckAndComments() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := block.Parse(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("line is: %s\n", b.ToLineWithCheckAndComments())

	// Output: line is: N7 G1 X2.0 Y2.0 F3000.0
}
