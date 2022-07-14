package block

import (
	"fmt"
	"testing"

	"github.com/mauroalderete/gcode-cli/checksum"
	"github.com/mauroalderete/gcode-cli/gcode"
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
				b, err := Parse(c.source, checksum.New())
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

func TestBlockFields(t *testing.T) {
	var cases = [1]struct {
		source string
	}{
		{"N7 G1 X2.0 Y2.0 F3000.0"}, //*85
	}

	t.Run("command", func(t *testing.T) {
		for i, c := range cases {
			t.Run(fmt.Sprintf("(%d)", i), func(t *testing.T) {
				b, err := Parse(c.source, checksum.New())
				if err != nil {
					t.Errorf("got %v, want nil error", err)
					return
				}
				if b == nil {
					t.Errorf("got nil block, want %v", c.source)
					return
				}

				if b.Command() == nil {
					t.Errorf("got command: nil, want command: not nil")
					return
				}

				if b.Command().HasAddress() {
					if bca, ok := b.Command().(*gcode.GcodeAddressable[int32]); ok {
						jj := bca.Address().Value()
						bca.Address().SetValue(jj + 10)
					}
				}

				if b.Command().String() != "G11" {
					t.Errorf("got command: %v, want command: G11", b.Command().String())
				}
			})
		}
	})
}

func TestBlockChecksumUpdate(t *testing.T) {
	cases := []struct {
		line     string
		checksum uint32
	}{
		{"N3 T0", 57},
		{"N4 G92 E0", 67},
		{"N5 G28", 22},
		{"N6 G1 F1500.0", 82},
		{"N7 G1 X2.0 Y2.0 F3000.0", 85},
		{"N8 G1 X3.0 Y3.0", 33},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("(%d)", i), func(t *testing.T) {

			b, err := Parse(c.line, checksum.New())
			if err != nil {
				t.Errorf("got block error: %v, want block error: nil", err)
			}

			err = b.UpdateChecksum()
			if err != nil {
				t.Errorf("got checksum error: %v, want checksum error: nil", err)
			}

			if b.Checksum().Address().Value() != c.checksum {
				t.Errorf("got %v, want %v", b.Checksum(), c.checksum)
			}
		})
	}
}

func TestBlockChecksumVerify(t *testing.T) {
	cases := []struct {
		line     string
		checksum uint32
	}{
		{"N3 T0", 57},
		{"N4 G92 E0", 67},
		{"N5 G28", 22},
		{"N6 G1 F1500.0", 82},
		{"N7 G1 X2.0 Y2.0 F3000.0", 85},
		{"N8 G1 X3.0 Y3.0", 33},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("(%d)", i), func(t *testing.T) {

			b, err := Parse(c.line, checksum.New())
			if err != nil {
				t.Errorf("got block error: %v, want block error: nil", err)
			}

			err = b.UpdateChecksum()
			if err != nil {
				t.Errorf("got checksum error: %v, want checksum error: nil", err)
			}

			res, err := b.VerifyChecksum()
			if err != nil {
				t.Errorf("got verify error: %v, want verify error: nil", err)
			}

			if !res {
				t.Errorf("got %v, want %v", b.Checksum(), c.checksum)
			}
		})
	}
}
