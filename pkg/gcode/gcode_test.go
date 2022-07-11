package gcode_test

import (
	"fmt"
	"testing"

	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/address"
	"github.com/mauroalderete/gcode-skew-transform-cli/pkg/gcode"
)

func TestNewGcode(t *testing.T) {
	t.Run("valids", func(t *testing.T) {

		cases := []struct {
			word byte
		}{
			{'W'},
			{'X'},
			{'N'},
		}

		for i, c := range cases {
			t.Run(fmt.Sprintf("(%v)", i), func(t *testing.T) {
				gc, err := gcode.NewGcode(c.word)
				if err != nil {
					t.Errorf("got %v, want X12", err)
					return
				}
				if gc == nil {
					t.Errorf("got nil gcode, want %v", c.word)
					return
				}
				if gc.String() != string(c.word) {
					t.Errorf("got %s, want %v", gc, c.word)
				}
			})
		}
	})

	t.Run("invalids", func(t *testing.T) {

		t.Run("word", func(t *testing.T) {
			cases := []struct {
				word byte
			}{
				{'+'},
				{'\t'},
				{'"'},
			}

			for i, c := range cases {
				t.Run(fmt.Sprintf("(%v)", i), func(t *testing.T) {
					gc, err := gcode.NewGcode(c.word)
					if err == nil {
						t.Errorf("got nil error, want not nil error")
					}
					if gc != nil {
						t.Errorf("got %v gcode, want nil gcode", gc)
					}
				})
			}
		})
	})
}

func TestNewGcodeAddressable(t *testing.T) {
	t.Run("valids", func(t *testing.T) {

		t.Run("address integer", func(t *testing.T) {
			gc, err := gcode.NewGcodeAddressable[int32]('X', 12)
			if err != nil {
				t.Errorf("got %v, want X12", err)
				return
			}
			if gc.String() != "X12" {
				t.Errorf("got %s, want X12", gc)
			}
		})

		t.Run("address float", func(t *testing.T) {
			gc, err := gcode.NewGcodeAddressable[float32]('X', 12.3)
			if err != nil {
				t.Errorf("got %v, want X12.3", err)
				return
			}
			if gc.String() != "X12.3" {
				t.Errorf("got %s, want X12.3", gc)
			}
		})

		t.Run("address string", func(t *testing.T) {
			gc, err := gcode.NewGcodeAddressable('X', "\"lorem ipsu\"")
			if err != nil {
				t.Errorf("got %v, want X\"lorem ipsu\"", err)
				return
			}
			if gc.String() != "X\"lorem ipsu\"" {
				t.Errorf("got %s, want X\"lorem ipsu\"", gc)
			}
		})
	})

	t.Run("invalids", func(t *testing.T) {

		t.Run("word", func(t *testing.T) {
			cases := []struct {
				word    byte
				address int32
			}{
				{'+', 2},
				{'\t', 2},
				{'"', 2},
			}

			for i, c := range cases {
				t.Run(fmt.Sprintf("(%v)", i), func(t *testing.T) {
					gc, err := gcode.NewGcodeAddressable(c.word, c.address)
					if err == nil {
						t.Errorf("got nil error, want not nil error")
					}
					if gc != nil {
						t.Errorf("got %v gcode, want nil gcode", gc)
					}
				})
			}
		})

		t.Run("address string", func(t *testing.T) {
			cases := []struct {
				word    byte
				address string
			}{
				{'X', ""},
				{'X', "\"\t\""},
				{'X', "\"\"\""},
			}

			for i, c := range cases {
				t.Run(fmt.Sprintf("(%v)", i), func(t *testing.T) {
					gc, err := gcode.NewGcodeAddressable(c.word, c.address)
					if err == nil {
						t.Errorf("got nil error, want not nil error")
					}
					if gc != nil {
						t.Errorf("got %v gcode, want nil gcode", gc)
					}
				})
			}
		})
	})
}

func TestAddressSetValuePersistenceOnGcodeAddressable(t *testing.T) {
	t.Run("caso 1", func(t *testing.T) {
		gc, err := gcode.NewGcodeAddressable[int32]('X', 99)
		if err != nil {
			t.Errorf("got error: %v, want error: nil", err)
		}
		if gc == nil {
			t.Errorf("got gcode: nil, want gcode: not nil")
		}

		var add *address.Address[int32]
		add = gc.Address()
		err = add.SetValue(12)
		if err != nil {
			t.Errorf("got error: %v, want error: nil", err)
		}

		add = gc.Address()
		err = add.SetValue(120)
		if err != nil {
			t.Errorf("got error: %v, want error: nil", err)
		}

		add = gc.Address()
		if add.Value() != 120 {
			t.Errorf("got address: %v, want address: 120", add.Value())
		}

	})
}

func TestAddressInmutableOnGcodeAddressable(t *testing.T) {
	t.Run("caso 1", func(t *testing.T) {
		gc, err := gcode.NewGcodeAddressable[int32]('X', 100)
		if err != nil {
			t.Errorf("got error: %v, want error: nil", err)
		}
		if gc == nil {
			t.Errorf("got gcode: nil, want gcode: not nil")
		}

		add, err := address.NewAddress[int32](101)
		if err != nil {
			t.Errorf("got error: %v, want error: nil", err)
		}
		if add == nil {
			t.Errorf("got address: nil, want address: not nil")
		}

		addFromGcode := gc.Address()
		addFromGcode = add
		addFromGcode = gc.Address()

		if addFromGcode.Value() == 101 {
			t.Errorf("got address: %v, want address: 100", addFromGcode.Value())
		}
	})
}

func TestGcodeCompare(t *testing.T) {

	gcodeA, err := gcode.NewGcode('M')
	if err != nil {
		t.Errorf("got %v, want nil error", err)
	}

	gcodeB, err := gcode.NewGcode('X')
	if err != nil {
		t.Errorf("got %v, want nil error", err)
	}

	gcodeAddressableA, err := gcode.NewGcodeAddressable[int32]('N', 11)
	if err != nil {
		t.Errorf("got %v, want nil error", err)
	}

	gcodeAddressableB, err := gcode.NewGcodeAddressable[int32]('N', 22)
	if err != nil {
		t.Errorf("got %v, want nil error", err)
	}

	cases := []struct {
		gcodeBase   gcode.Gcoder
		gcodeTarget gcode.Gcoder
		result      bool
	}{
		{gcodeA, gcodeA, true},
		{gcodeA, gcodeB, false},
		{gcodeA, gcodeAddressableA, false},
		{gcodeA, gcodeAddressableB, false},
		{gcodeA, nil, false},
		{gcodeB, gcodeA, false},
		{gcodeB, gcodeB, true},
		{gcodeB, gcodeAddressableA, false},
		{gcodeB, gcodeAddressableB, false},
		{gcodeB, nil, false},
		{gcodeAddressableA, gcodeA, false},
		{gcodeAddressableA, gcodeB, false},
		{gcodeAddressableA, gcodeAddressableA, true},
		{gcodeAddressableA, gcodeAddressableB, false},
		{gcodeAddressableA, nil, false},
		{gcodeAddressableB, gcodeA, false},
		{gcodeAddressableB, gcodeB, false},
		{gcodeAddressableB, gcodeAddressableA, false},
		{gcodeAddressableB, gcodeAddressableB, true},
		{gcodeAddressableB, nil, false},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("(%d)", i), func(t *testing.T) {
			if c.gcodeBase.Compare(c.gcodeTarget) != c.result {
				t.Errorf("got %v, want %v", !c.result, c.result)
			}
		})
	}
}
