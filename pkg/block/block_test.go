package block_test

import (
	"fmt"
	"testing"

	block "github.com/mauroalderete/gcode-skew-transform-cli/pkg/block"
)

//#region unit tests

func TestParse(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		var cases = [1]struct {
			source string
		}{
			{"N7 G1 X2.0 Y2.0 F3000.0"}, //*85
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("(%v)", i), func(t *testing.T) {
				b, err := block.Parse(c.source)
				if err != nil {
					t.Errorf("got %v, want nil error", err)
					return
				}
				if b == nil {
					t.Errorf("got nil block, want %v", c.source)
					return
				}
				if b.ToLineWithCheckAndComments() != c.source {
					t.Errorf("got %v, want %v", b.ToLineWithCheckAndComments(), c.source)
				}
			})
		}
	})
}

//#endregion
//#region examples

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

	fmt.Printf("checksum is: %s\n", b.Checksum().Value().String())

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

	if b.Comment() == nil {
		fmt.Println("comment isn't available")
		return
	}

	fmt.Printf("comment is: %s\n", *b.Comment())

	// Output: comment isn't available
}

func ExampleBlock_IsChecked() {
	const source = "N7 G1 X2.0 Y2.0 F3000.0"

	b, err := block.Parse(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	checked, err := b.IsChecked()

	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return
	}

	fmt.Printf("isChecked: %v\n", checked)

	// Output: this block hasn't check section
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

	fmt.Printf("line number is: %v\n", b.LineNumber().Address().Value)

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

//#endregion
