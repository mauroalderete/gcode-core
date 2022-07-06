package block_test

import (
	"fmt"
	"testing"

	block "github.com/mauroalderete/gcode-skew-transform-cli/pkg/block"
)

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
				// if b.ToCommandComplete() != c.source {
				// 	t.Errorf("got %v, want %v", b.ToCommandComplete(), c.source)
				// }
			})
		}
	})
}

// func ExampleParse() {
// 	const source = "N7 G1 X2.0 Y2.0 F3000.0*85"

// 	b, err := block.Parse(source)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	fmt.Printf("the command is: %s\n", b.ToCommandComplete())

// 	// Output: the command is: N7 G1 X2.0 Y2.0 F3000.0*85
// }
