package checksum_test

import (
	"fmt"
	"testing"

	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/checksum"
)

func TestNew(t *testing.T) {
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

	h := checksum.New()

	for i, c := range cases {
		t.Run(fmt.Sprintf("(%d)", i), func(t *testing.T) {
			h.Reset()
			_, err := h.Write([]byte(c.line))
			if err != nil {
				t.Errorf("got error: not nil, want error: %v", err)
			}

			if c.checksum != uint32(h.Sum(nil)[0]) {
				t.Errorf("got %v, want %v", h.Sum(nil)[0], c.checksum)
			}
		})
	}
}

func TestConfiguration(t *testing.T) {
	t.Run("size", func(t *testing.T) {
		if checksum.New().Size() != 1 {
			t.Errorf("got %v, want 1", checksum.New().Size())
		}
	})

	t.Run("block size", func(t *testing.T) {
		if checksum.New().BlockSize() != 1 {
			t.Errorf("got %v, want 1", checksum.New().BlockSize())
		}
	})
}
