package unaddressablegcode

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

type mockHasAddressGcoder interface {
	HasAddress() bool
}

type mockWordGcoder interface {
	Word() byte
}

type mockAddressableGcode struct{}

func (ag *mockAddressableGcode) Compare(gcode.Gcoder) bool {
	return false
}

func (ag *mockAddressableGcode) HasAddress() bool {
	return true
}

func (ag *mockAddressableGcode) String() string {
	return ""
}

func (ag *mockAddressableGcode) Word() byte {
	return '?'
}

//#endregion
//#region unit test

func TestNewGcode(t *testing.T) {

	cases := map[string]struct {
		input byte
		valid bool
	}{
		"eval_W":   {'W', true},
		"eval_X":   {'X', true},
		"eval_N":   {'N', true},
		"eval_+":   {'+', false},
		"eval_\\t": {'\t', false},
		"eval_\"":  {'"', false},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			gc, err := New(tc.input)

			if tc.valid {
				if err != nil {
					t.Errorf("got error %v, want error nil", err)
					return
				}
				if gc == nil {
					t.Errorf("got gcode nil, want gcode %s", string(tc.input))
					return
				}
				if gc.String() != string(tc.input) {
					t.Errorf("got gcode %s, want gcode %s", gc, string(tc.input))
				}
			} else {
				if err == nil {
					t.Errorf("got error %v, want error nil", err)
				}
				if gc != nil {
					t.Errorf("got gcode %s, want gcode nil", gc.String())
				}
			}
		})
	}
}

func TestGcodeCompare(t *testing.T) {

	gcodeA, err := New('M')
	if err != nil {
		t.Errorf("got %v, want nil error", err)
	}

	gcodeB, err := New('X')
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
		"gca == nil": {gcodeA, nil, false},
		"gcb == gca": {gcodeB, gcodeA, false},
		"gcb == gcb": {gcodeB, gcodeB, true},
		"gcb == nil": {gcodeB, nil, false},
		"gca == agc": {gcodeA, &mockAddressableGcode{}, false},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if tc.inputA.Compare(tc.inputB) != tc.want {
				t.Errorf("got %v, want %v", !tc.want, tc.want)
			}
		})
	}
}

func TestGcodeHasAddress(t *testing.T) {

	gc, err := New('M')
	if err != nil {
		t.Errorf("got %v, want nil error", err)
	}

	cases := map[string]struct {
		input mockHasAddressGcoder
		want  bool
	}{
		"no address":   {gc, false},
		"with address": {&mockAddressableGcode{}, true},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if tc.input.HasAddress() != tc.want {
				t.Errorf("got %v, want %v", tc.input.HasAddress(), tc.want)
			}
		})
	}
}

func TestGcodeString(t *testing.T) {

	gc, err := New('M')
	if err != nil {
		t.Errorf("got %v, want nil error", err)
	}

	cases := map[string]struct {
		input fmt.Stringer
		want  string
	}{
		"M": {gc, "M"},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if tc.input.String() != tc.want {
				t.Errorf("got %v, want %v", tc.input.String(), tc.want)
			}
		})
	}
}

func TestGcodeWord(t *testing.T) {

	gc, err := New('M')
	if err != nil {
		t.Errorf("got %v, want nil error", err)
	}

	cases := map[string]struct {
		input mockWordGcoder
		want  byte
	}{
		"M": {gc, 'M'},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if tc.input.Word() != tc.want {
				t.Errorf("got %v, want %v", tc.input.Word(), tc.want)
			}
		})
	}
}
