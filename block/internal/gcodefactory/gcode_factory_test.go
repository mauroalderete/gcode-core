package gcodefactory

import (
	"fmt"
	"testing"
)

func TestGcodeFactoryNewGcode(t *testing.T) {
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

	gcodeFactory := &GcodeFactory{}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			gc, err := gcodeFactory.NewUnaddressableGcode(tc.input)

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

func TestGcodeAddressableFactoryNewGcodeAddressable(t *testing.T) {
	gcodeFactory := &GcodeFactory{}

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
				gc, err := gcodeFactory.NewAddressableGcodeUint32(tc.word, tc.address)
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
				gc, err := gcodeFactory.NewAddressableGcodeInt32(tc.word, tc.address)
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
				gc, err := gcodeFactory.NewAddressableGcodeFloat32(tc.word, tc.address)
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
				gc, err := gcodeFactory.NewAddressableGcodeString(tc.word, tc.address)
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

func TestParse(t *testing.T) {

	// output:           "N4 G92 E0*67 ;comentario",
	cases := map[string]struct {
		input  string
		valid  bool
		output string
	}{
		"command_0":      {"G92", true, "G92"},
		"command_1":      {"G92.3", true, "G92.3"},
		"command_2":      {"G\"hola\"", true, "G\"hola\""},
		"command_3":      {"N1", true, "N1"},
		"command_4":      {"*1", true, "*1"},
		"command_5":      {"G", true, "G"},
		"command_6":      {"G-92", true, "G-92"},
		"command_7":      {"G-92.0", true, "G-92.0"},
		"command_8":      {"N92.0", false, ""},
		"command_fail_0": {"G 92", false, ""},
		"command_fail_1": {"K92.3", false, ""},
		"command_fail_2": {"G\"\"hola\"", false, ""},
		"command_fail_3": {"N-1", false, ""},
		"command_fail_4": {"*-1", false, ""},
		"command_fail_5": {"", false, ""},
		"command_fail_6": {"K", false, ""},
		"command_fail_7": {"G1.x", false, ""},
		"command_fail_8": {"Nx", false, ""},
	}

	mockGcodeFactory := &GcodeFactory{}

	for name, tc := range cases {
		t.Run(fmt.Sprintf("%s[%s]", name, tc.input), func(t *testing.T) {
			gb, err := mockGcodeFactory.Parse(tc.input)
			if tc.valid {
				if err != nil {
					t.Errorf("got error %v, want error nil", err)
				}

				if gb == nil {
					t.Errorf("got gcodeBlock nil, want gcodeBlock not nil")
					return
				}

				if gb.String() != tc.output {
					t.Errorf("got gcodeBlock (%d)[%s], want gcodeBlock: (%d)[%s]", len(gb.String()), gb.String(), len(tc.output), tc.output)
				}
			} else {
				if err == nil {
					t.Errorf("got error nil, want error not nil")
				}

				if gb != nil {
					t.Errorf("got gcodeBlock not nil, want gcodeBlock nil")
				}
			}
		})
	}
}
