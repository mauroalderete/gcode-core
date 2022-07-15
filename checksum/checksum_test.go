package checksum

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	cases := map[string]struct {
		line     string
		checksum uint32
	}{
		"1_": {"N3 T0", 57},
		"2_": {"N4 G92 E0", 67},
		"3_": {"N5 G28", 22},
		"4_": {"N6 G1 F1500.0", 82},
		"5_": {"N7 G1 X2.0 Y2.0 F3000.0", 85},
		"6_": {"N8 G1 X3.0 Y3.0", 33},
	}

	h := New()

	for name, tc := range cases {
		t.Run(fmt.Sprintf("%s%s)", name, tc.line), func(t *testing.T) {
			h.Reset()
			_, err := h.Write([]byte(tc.line))
			if err != nil {
				t.Errorf("got error: not nil, want error: %v", err)
			}

			if tc.checksum != uint32(h.Sum(nil)[0]) {
				t.Errorf("got %v, want %v", h.Sum(nil)[0], tc.checksum)
			}
		})
	}
}

func TestConfiguration(t *testing.T) {
	t.Run("size", func(t *testing.T) {
		if New().Size() != 1 {
			t.Errorf("got %v, want 1", New().Size())
		}
	})

	t.Run("block size", func(t *testing.T) {
		if New().BlockSize() != 1 {
			t.Errorf("got %v, want 1", New().BlockSize())
		}
	})
}
