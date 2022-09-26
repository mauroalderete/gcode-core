package addressablegcode

import (
	"fmt"
	"testing"

	"github.com/mauroalderete/gcode-core/gcode"
)

//#region Mockups

type mockGcoder interface {
	fmt.Stringer

	Compare(gcode.Gcoder) bool
	HasAddress() bool
	Word() byte
}

type mockComparerGcoder interface {
	Compare(gcode.Gcoder) bool
}

type mockUnaddressableGcode struct{}

func (ag *mockUnaddressableGcode) Compare(gcode.Gcoder) bool {
	return false
}

func (ag *mockUnaddressableGcode) HasAddress() bool {
	return false
}

func (ag *mockUnaddressableGcode) String() string {
	return ""
}

func (ag *mockUnaddressableGcode) Word() byte {
	return '?'
}

//#endregion
//#region unit tests

func TestNewGcodeAddressable(t *testing.T) {

	t.Run("addressable gcode uint32", func(t *testing.T) {
		cases := map[string]struct {
			word    byte
			address uint32
			valid   bool
		}{
			"eval_W0":   {'W', 0, true},
			"eval_X1":   {'X', 1, true},
			"eval_N2":   {'N', 2, true},
			"eval_+3":   {'+', 3, false},
			"eval_\\t4": {'\t', 4, false},
			"eval_\"5":  {'"', 5, false},
		}

		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				gc, err := New(tc.word, tc.address)
				if tc.valid {
					if err != nil {
						t.Errorf("got error %v, want error nil", err)
						return
					}
					if gc.String() != fmt.Sprintf("%s%d", string(tc.word), tc.address) {
						t.Errorf("got gcode %s, want gcode %s%d", gc, string(tc.word), tc.address)
					}
				} else {
					if err == nil {
						t.Errorf("got error nil, want error not nil")
						return
					}
				}
			})
		}
	})

	t.Run("addressable gcode int32", func(t *testing.T) {
		cases := map[string]struct {
			word    byte
			address int32
			valid   bool
		}{
			"eval_W0":   {'W', -1, true},
			"eval_X1":   {'X', 0, true},
			"eval_N2":   {'N', 1, true},
			"eval_+3":   {'+', 3, false},
			"eval_\\t4": {'\t', 4, false},
			"eval_\"5":  {'"', 5, false},
		}

		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				gc, err := New(tc.word, tc.address)
				if tc.valid {
					if err != nil {
						t.Errorf("got error %v, want error nil", err)
						return
					}
					if gc.String() != fmt.Sprintf("%s%d", string(tc.word), tc.address) {
						t.Errorf("got gcode %s, want gcode %s%d", gc, string(tc.word), tc.address)
					}
				} else {
					if err == nil {
						t.Errorf("got error nil, want error not nil")
						return
					}
				}
			})
		}
	})

	t.Run("addressable gcode float32", func(t *testing.T) {
		cases := map[string]struct {
			word    byte
			address float32
			valid   bool
		}{
			"eval_W0":     {'W', 0, true},
			"eval_X1.1":   {'X', 1.1, true},
			"eval_N2.2":   {'N', 2.2, true},
			"eval_+3.3":   {'+', 3.3, false},
			"eval_\\t4.4": {'\t', 4.4, false},
			"eval_\"5.5":  {'"', 5.5, false},
		}

		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				gc, err := New(tc.word, tc.address)
				if tc.valid {
					if err != nil {
						t.Errorf("got error %v, want error nil", err)
						return
					}
					if gc.String() != fmt.Sprintf("%s%.1f", string(tc.word), tc.address) {
						t.Errorf("got gcode %s, want gcode %s%.1f", gc, string(tc.word), tc.address)
					}
				} else {
					if err == nil {
						t.Errorf("got error nil, want error not nil")
						return
					}
				}
			})
		}
	})

	t.Run("addressable gcode string", func(t *testing.T) {
		cases := map[string]struct {
			word    byte
			address string
			valid   bool
		}{
			"eval_W":   {'W', "\"Hola mundo\"", true},
			"eval_X":   {'X', "\"Hola \"\"mundo\"\"\"", true},
			"eval_N":   {'N', "\"Hola mundo\"", true},
			"eval_+":   {'+', "\"Hola mundo\"", false},
			"eval_\\t": {'\t', "\"Hola mundo\"", false},
			"eval_\"":  {'"', "\"Hola mundo\"", false},
			"eval_W2":  {'W', "Hola mundo\"", false},
			"eval_X2":  {'X', "\"Hola \"mundo\"", false},
			"eval_N2":  {'N', "\"Hola mundo", false},
			"eval_W3":  {'W', "\"\tHola mundo\"", false},
			"eval_X3":  {'X', "?", false},
		}

		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				gc, err := New(tc.word, tc.address)
				if tc.valid {
					if err != nil {
						t.Errorf("got error %v, want error nil", err)
						return
					}
					if gc.String() != fmt.Sprintf("%s%s", string(tc.word), tc.address) {
						t.Errorf("got gcode %s, want gcode %s%s", gc, string(tc.word), tc.address)
					}
				} else {
					if err == nil {
						t.Errorf("got error nil, want error not nil")
						return
					}
				}
			})
		}
	})
}

func TestGcodeCompare(t *testing.T) {

	gcodeA, err := New[float32]('M', 1)
	if err != nil {
		t.Errorf("got %v, want nil error", err)
	}

	gcodeB, err := New[float32]('X', 1)
	if err != nil {
		t.Errorf("got %v, want nil error", err)
	}

	gcodeC, err := New[float32]('M', 1.0)
	if err != nil {
		t.Errorf("got %v, want nil error", err)
	}

	cases := map[string]struct {
		inputA mockComparerGcoder
		inputB mockGcoder
		want   bool
	}{
		"gca == gca": {gcodeA, gcodeA, true},
		"gca == gcb": {gcodeA, gcodeB, false},
		"gca == gcc": {gcodeA, gcodeC, true},
		"gca == ugc": {gcodeA, &mockUnaddressableGcode{}, false},
		"gca == nil": {gcodeA, nil, false},

		"gcb == gca": {gcodeB, gcodeA, false},
		"gcb == gcb": {gcodeB, gcodeB, true},
		"gcb == gcc": {gcodeB, gcodeC, false},
		"gcb == ugc": {gcodeB, &mockUnaddressableGcode{}, false},
		"gcb == nil": {gcodeB, nil, false},

		"gcc == gca": {gcodeC, gcodeA, true},
		"gcc == gcb": {gcodeC, gcodeB, false},
		"gcc == gcc": {gcodeC, gcodeC, true},
		"gcc == ugc": {gcodeC, &mockUnaddressableGcode{}, false},
		"gcc == nil": {gcodeC, nil, false},

		"ugc == gca": {&mockUnaddressableGcode{}, gcodeA, false},
		"ugc == gcb": {&mockUnaddressableGcode{}, gcodeB, false},
		"ugc == gcc": {&mockUnaddressableGcode{}, gcodeC, false},
		"ugc == nil": {&mockUnaddressableGcode{}, nil, false},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if tc.inputA.Compare(tc.inputB) != tc.want {
				t.Errorf("got %v, want %v", !tc.want, tc.want)
			}
		})
	}
}

func TestAddressableGcodeHasAddress(t *testing.T) {

	t.Run("address integer", func(t *testing.T) {
		gc, err := New[int32]('X', 12)
		if err != nil {
			t.Errorf("got error %v, want error nil", err)
			return
		}
		if !gc.HasAddress() {
			t.Errorf("got HasAddress false, want HasAddress true")
		}
	})
}

func TestAddresableGcodeSetAddress(t *testing.T) {

	t.Run("address integer valid", func(t *testing.T) {
		gc, err := New[int32]('X', 12)
		if err != nil {
			t.Errorf("got %v, want X", err)
			return
		}
		gc.SetAddress(10)
		if gc.Address() != 10 {
			t.Errorf("got %s, want 10", gc)
		}
	})

	t.Run("address string invalid", func(t *testing.T) {
		gc, err := New('X', "\"hola\"")
		if err != nil {
			t.Errorf("got %v, want X", err)
			return
		}
		err = gc.SetAddress("mundo")
		if err == nil {
			t.Errorf("got %s, want error: not nil", err)
		}
	})
}

func TestAddressableGcodeWord(t *testing.T) {

	t.Run("address integer", func(t *testing.T) {
		gc, err := New[int32]('X', 12)
		if err != nil {
			t.Errorf("got error %v, want error nil", err)
			return
		}
		if gc.Word() != 'X' {
			t.Errorf("got Word %s, want Word X", gc)
		}
	})
}

func TestAddressableGcodeString(t *testing.T) {
	t.Run("float32", func(t *testing.T) {

		var a float32 = 1.2
		var b float32 = 0.5

		cases := map[string]struct {
			address float32
			output  string
		}{
			"0": {1, "1.0"},
			"a": {0, "0.0"},
			"b": {1.1, "1.1"},
			"c": {2.2, "2.2"},
			"d": {float32(1.2), "1.2"},
			"e": {float32(0.5), "0.5"},
			"f": {1.2 - 0.5, "0.7"},
			"g": {float32(1.2) - 0.5, "0.7"},
			"h": {1.2 - float32(0.5), "0.7"},
			"i": {float32(1.2) - float32(0.5), "0.7"},
			"j": {float32(float32(1.2) - float32(0.5)), "0.7"},
			"k": {a - b, "0.7"},
			"l": {a / b, "2.4"},
			"m": {a * b, "0.6"},
			"n": {a * 0.5, "0.6"},
			"o": {b / 0.5, "1.0"},
		}

		for name, tc := range cases {
			t.Run(name, func(t *testing.T) {
				gc, err := New('G', tc.address)
				if err != nil {
					t.Errorf("failed prepare mock")
					return
				}

				if gc.String() != fmt.Sprintf("G%s", tc.output) {
					t.Errorf("failed print, expected %s, got %s", fmt.Sprintf("G%s", tc.output), gc.String())
					return
				}
			})
		}
	})
}

//#endregion
