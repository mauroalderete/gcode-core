package gcodeblock

import (
	"fmt"
	"hash"
	"testing"

	"github.com/mauroalderete/gcode-core/block"
	"github.com/mauroalderete/gcode-core/block/internal/gcodefactory"
	"github.com/mauroalderete/gcode-core/checksum"
	"github.com/mauroalderete/gcode-core/gcode"
	"github.com/mauroalderete/gcode-core/gcode/addressablegcode"
)

func TestNew(t *testing.T) {

	// N4 G92 E0*67 ;comentario
	mockLineNumber, err := addressablegcode.New[uint32]('N', 4)
	if err != nil {
		t.Errorf("got error not nil, want error nil: %v", err)
	}

	mockCommand, err := addressablegcode.New[int32]('G', 92)
	if err != nil {
		t.Errorf("got error not nil, want error nil: %v", err)
	}

	mockParam1, err := addressablegcode.New[int32]('E', 0)
	if err != nil {
		t.Errorf("got error not nil, want error nil: %v", err)
	}

	mockParameters := []gcode.Gcoder{mockParam1}

	mockChecksum, err := addressablegcode.New[uint32]('*', 67)
	if err != nil {
		t.Errorf("got error not nil, want error nil: %v", err)
	}

	mockChecksumFailed, err := addressablegcode.New[uint32]('*', 10)
	if err != nil {
		t.Errorf("got error not nil, want error nil: %v", err)
	}

	mockComment := ";comentario"

	mockHash := checksum.New()

	mockGcodeFactory := &gcodefactory.GcodeFactory{}

	cases := map[string]struct {
		lineNumber         gcode.AddressableGcoder[uint32]
		command            gcode.Gcoder
		parameters         []gcode.Gcoder
		checksum           gcode.AddressableGcoder[uint32]
		comment            string
		hash               hash.Hash
		gcodeFactory       gcode.GcoderFactory
		configLineNumber   bool
		configParameters   bool
		configChecksum     bool
		configComment      bool
		configHash         bool
		configGcodeFactory bool
		valid              bool
		output             string
	}{
		"Single Word command": {
			command: mockCommand,
			output:  "G92",
			valid:   true,
		},
		"no command": {
			command: nil,
			valid:   false,
		},
		"lineNumber nil": {
			command:          mockCommand,
			configLineNumber: true,
			valid:            false,
			output:           "",
		},
		"parameters nil": {
			command:          mockCommand,
			configParameters: true,
			valid:            false,
			output:           "",
		},
		"checksum nil": {
			command:        mockCommand,
			configChecksum: true,
			valid:          false,
			output:         "",
		},
		"hash nil": {
			command:    mockCommand,
			configHash: true,
			valid:      false,
			output:     "",
		},
		"gcodeFactory nil": {
			command:            mockCommand,
			configGcodeFactory: true,
			valid:              false,
			output:             "",
		},
		"+lineNumber": {
			command:          mockCommand,
			configLineNumber: true,
			lineNumber:       mockLineNumber,
			valid:            true,
			output:           "N4 G92",
		},
		"+linenumber+parameters": {
			command:          mockCommand,
			configLineNumber: true,
			lineNumber:       mockLineNumber,
			configParameters: true,
			parameters:       mockParameters,
			valid:            true,
			output:           "N4 G92 E0",
		},
		"+linenumber+parameters+checksum": {
			command:          mockCommand,
			configLineNumber: true,
			lineNumber:       mockLineNumber,
			configParameters: true,
			parameters:       mockParameters,
			configChecksum:   true,
			checksum:         mockChecksum,
			valid:            true,
			output:           "N4 G92 E0*67",
		},
		"+checksum not verified": {
			command:          mockCommand,
			configLineNumber: true,
			lineNumber:       mockLineNumber,
			configParameters: true,
			parameters:       mockParameters,
			configChecksum:   true,
			checksum:         mockChecksumFailed,
			valid:            false,
		},
		"+linenumber+parameters+checksum+comment": {
			command:          mockCommand,
			configLineNumber: true,
			lineNumber:       mockLineNumber,
			configParameters: true,
			parameters:       mockParameters,
			configChecksum:   true,
			checksum:         mockChecksum,
			configComment:    true,
			comment:          mockComment,
			valid:            true,
			output:           "N4 G92 E0*67 ;comentario",
		},
		"+linenumber+parameters+checksum+comment+hash": {
			command:          mockCommand,
			configLineNumber: true,
			lineNumber:       mockLineNumber,
			configParameters: true,
			parameters:       mockParameters,
			configChecksum:   true,
			checksum:         mockChecksum,
			configComment:    true,
			comment:          mockComment,
			configHash:       true,
			hash:             mockHash,
			valid:            true,
			output:           "N4 G92 E0*67 ;comentario",
		},
		"+linenumber+parameters+checksum+comment+hash+gcodeFactory": {
			command:            mockCommand,
			configLineNumber:   true,
			lineNumber:         mockLineNumber,
			configParameters:   true,
			parameters:         mockParameters,
			configChecksum:     true,
			checksum:           mockChecksum,
			configComment:      true,
			comment:            mockComment,
			configHash:         true,
			hash:               mockHash,
			configGcodeFactory: true,
			gcodeFactory:       mockGcodeFactory,
			valid:              true,
			output:             "N4 G92 E0*67 ;comentario",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			gb, err := New(tc.command, func(config block.BlockConstructorConfigurer) error {

				if tc.configLineNumber {
					err := config.SetLineNumber(tc.lineNumber)
					if err != nil {
						return err
					}
				}

				if tc.configParameters {
					err := config.SetParameters(tc.parameters)
					if err != nil {
						return err
					}
				}

				if tc.configChecksum {
					err := config.SetChecksum(tc.checksum)
					if err != nil {
						return err
					}
				}

				if tc.configComment {
					err := config.SetComment(tc.comment)
					if err != nil {
						return err
					}
				}

				if tc.configGcodeFactory {
					err := config.SetGcodeFactory(tc.gcodeFactory)
					if err != nil {
						return err
					}
				}

				if tc.configHash {
					err := config.SetHash(tc.hash)
					if err != nil {
						return err
					}
				}

				return nil
			})

			if tc.valid {
				if err != nil {
					t.Errorf("got error %v, want error nil", err)
				}

				if gb == nil {
					t.Errorf("got gcodeBlock nil, want gcodeBlock not nil")
					return
				}

				if gb.ToLine("%l %c %p%k %m") != tc.output {
					t.Errorf("got gcodeBlock (%d)[%s], want gcodeBlock: (%d)[%s]", len(gb.ToLine("%l %c %p%k %m")), gb.ToLine("%l %c %p%k %m"), len(tc.output), tc.output)
				}
			} else {
				if err == nil {
					t.Errorf("got error nil, want error not nil")
				}
			}
		})
	}
}

func TestParse(t *testing.T) {

	mockHash := checksum.New()
	mockGcodeFactory := &gcodefactory.GcodeFactory{}

	// output:           "N4 G92 E0*67 ;comentario",
	cases := map[string]struct {
		input  string
		valid  bool
		output string
	}{
		"command_0":               {"G92", true, "G92"},
		"command_1":               {"G92 1", false, ""},
		"command_2":               {"*G92", false, ""},
		"command_3":               {" G92", true, "G92"},
		"command_4":               {"G92 ", true, "G92"},
		"command_5":               {"G 92", false, ""},
		"command_6":               {"\"G92", false, ""},
		"command_7":               {"\"G92\"", false, ""},
		"command_8":               {"G\"92\"", true, "G\"92\""},
		"command_9":               {"G\" 92\"", true, "G\" 92\""},
		"command_10":              {"G\"92 \"", true, "G\"92 \""},
		"command_11":              {"G\" 92 \"", true, "G\" 92 \""},
		"command_12":              {"G\"\"\"92\"\"\"", true, "G\"\"\"92\"\"\""},
		"command_13":              {"G\"\"\"\"92\"\"\"", false, ""},
		"command_14":              {"G\"\"\"92\"\"\"\"", false, ""},
		"command_15":              {" G\"\"\"92\"\"\"", true, "G\"\"\"92\"\"\""},
		"command_16":              {" G\"\"\"\"92\"\"\"", false, ""},
		"command_17":              {" G\"\"\"92\"\"\"\"", false, ""},
		"command_18":              {"G\"\"\"92\"\"\" ", true, "G\"\"\"92\"\"\""},
		"command_19":              {"G\"\"\"\"92\"\"\" ", false, ""},
		"command_20":              {"G\"\"\"92\"\"\"\" ", false, ""},
		"command_21":              {"G\"\" \"92\" \"\"", false, ""},
		"command_22":              {"G\" \"\"92\"\" \"", true, "G\" \"\"92\"\" \""},
		"command_23":              {"G\" \"\"92 \"\" \"", true, "G\" \"\"92 \"\" \""},
		"command_24":              {"G\" \"\" 92 \"\" \"", true, "G\" \"\" 92 \"\" \""},
		"linenumber_0":            {"N100 G92", true, "N100 G92"},
		"linenumber_1":            {"N100 G92 1", false, ""},
		"linenumber_2":            {"N100 *G92", false, ""},
		"linenumber_3":            {"N100  G92", true, "N100 G92"},
		"linenumber_4":            {"N100 G92 ", true, "N100 G92"},
		"linenumber_5":            {"N100 G 92", false, ""},
		"linenumber_6":            {"N100 \"G92", false, ""},
		"linenumber_7":            {"N100  \"G92\"", false, ""},
		"linenumber_8":            {"N100 G\"92\"", true, "N100 G\"92\""},
		"linenumber_9":            {"N100 G\" 92\"", true, "N100 G\" 92\""},
		"linenumber_10":           {"N100 G\"92 \"", true, "N100 G\"92 \""},
		"linenumber_11":           {"N100 G\" 92 \"", true, "N100 G\" 92 \""},
		"linenumber_12":           {"N100 G\"\"\"92\"\"\"", true, "N100 G\"\"\"92\"\"\""},
		"linenumber_13":           {"N100 G\"\"\"\"92\"\"\"", false, ""},
		"linenumber_14":           {"N100 G\"\"\"92\"\"\"\"", false, ""},
		"linenumber_15":           {"N100  G\"\"\"92\"\"\"", true, "N100 G\"\"\"92\"\"\""},
		"linenumber_16":           {"N100  G\"\"\"\"92\"\"\"", false, ""},
		"linenumber_17":           {"N100  G\"\"\"92\"\"\"\"", false, ""},
		"linenumber_18":           {"N100 G\"\"\"92\"\"\" ", true, "N100 G\"\"\"92\"\"\""},
		"linenumber_19":           {"N100 G\"\"\"\"92\"\"\" ", false, ""},
		"linenumber_20":           {"N100 G\"\"\"92\"\"\"\" ", false, ""},
		"linenumber_21":           {"N100 G\"\" \"92\" \"\"", false, ""},
		"linenumber_22":           {"N100 G\" \"\"92\"\" \"", true, "N100 G\" \"\"92\"\" \""},
		"linenumber_23":           {"N100 G\" \"\"92 \"\" \"", true, "N100 G\" \"\"92 \"\" \""},
		"linenumber_24":           {"N100 G\" \"\" 92 \"\" \"", true, "N100 G\" \"\" 92 \"\" \""},
		"linenumber_fail_0":       {"N 100 G92", false, ""},
		"linenumber_fail_1":       {"N  G92 1", false, ""},
		"linenumber_fail_2":       {"N 100 *G92", false, ""},
		"linenumber_fail_3":       {"N   G92", false, ""},
		"linenumber_fail_4":       {"N 100 G92 ", false, ""},
		"linenumber_fail_5":       {"N  G 92", false, ""},
		"linenumber_fail_6":       {"N 100 \"G92", false, ""},
		"linenumber_fail_7":       {"N  \"G92\"", false, ""},
		"linenumber_fail_8":       {"N 100 G\"92\"", false, ""},
		"linenumber_fail_9":       {"N  G\" 92\"", false, ""},
		"linenumber_fail_10":      {"N 100 G\"92 \"", false, ""},
		"linenumber_fail_11":      {"N  G\" 92 \"", false, ""},
		"linenumber_fail_12":      {"N 100 G\"\"\"92\"\"\"", false, ""},
		"linenumber_fail_13":      {"N  G\"\"\"\"92\"\"\"", false, ""},
		"linenumber_fail_14":      {"N 100 G\"\"\"92\"\"\"\"", false, ""},
		"linenumber_fail_15":      {"N   G\"\"\"92\"\"\"", false, ""},
		"linenumber_fail_16":      {"N 100  G\"\"\"\"92\"\"\"", false, ""},
		"linenumber_fail_17":      {"N   G\"\"\"92\"\"\"\"", false, ""},
		"linenumber_fail_18":      {"N 100 G\"\"\"92\"\"\" ", false, ""},
		"linenumber_fail_19":      {"N  G\"\"\"\"92\"\"\" ", false, ""},
		"linenumber_fail_20":      {"N 100 G\"\"\"92\"\"\"\" ", false, ""},
		"linenumber_fail_21":      {"N G\"\" \"92\" \"\"", false, ""},
		"linenumber_fail_22":      {"N 100 G\" \"\"92\"\" \"", false, ""},
		"linenumber_fail_23":      {"N G\" \"\"92 \"\" \"", false, ""},
		"linenumber_fail_24":      {"N 100 G\" \"\" 92 \"\" \"", false, ""},
		"parameters_0":            {"N100 G92 X1.0 Y2.0 Z3.0", true, "N100 G92 X1.0 Y2.0 Z3.0"},
		"parameters_1":            {"N100 G92 1X1.0 Y2.0 Z3.0", false, ""},
		"parameters_2":            {"N100 *G92 X1.0 Y2.0 Z3.0", false, ""},
		"parameters_3":            {"N100  G92   X1.0 Y2.0 Z3.0", true, "N100 G92 X1.0 Y2.0 Z3.0"},
		"parameters_4":            {"N100 G92 X1.0 Y2.0 Z3.0", true, "N100 G92 X1.0 Y2.0 Z3.0"},
		"parameters_5":            {"N100 G 92 X1.0 Y2.0 Z3.0", false, ""},
		"parameters_6":            {"N100 \"G92 X1.0 Y2.0 Z3.0", false, ""},
		"parameters_7":            {"N100  \"G92\" X1.0 Y2.0 Z3.0", false, ""},
		"parameters_8":            {"N100 G\"92\" X1.0 Y2.0 Z3.0", true, "N100 G\"92\" X1.0 Y2.0 Z3.0"},
		"parameters_9":            {"N100 G\" 92\" X1.0 Y2.0 Z3.0", true, "N100 G\" 92\" X1.0 Y2.0 Z3.0"},
		"parameters_10":           {"N100 G\"92 \" X1.0 Y2.0 Z3.0", true, "N100 G\"92 \" X1.0 Y2.0 Z3.0"},
		"parameters_11":           {"N100 G\" 92 \" X1.0 Y2.0 Z3.0", true, "N100 G\" 92 \" X1.0 Y2.0 Z3.0"},
		"parameters_12":           {"N100 G\"\"\"92\"\"\" X1.0 Y2.0 Z3.0", true, "N100 G\"\"\"92\"\"\" X1.0 Y2.0 Z3.0"},
		"parameters_13":           {"N100 G\"\"\"\"92\"\"\" X1.0 Y2.0 Z3.0", false, ""},
		"parameters_14":           {"N100 G\"\"\"92\"\"\"\" X1.0 Y2.0 Z3.0", false, ""},
		"parameters_15":           {"N100  G\"\"\"92\"\"\" X1.0 Y2.0 Z3.0", true, "N100 G\"\"\"92\"\"\" X1.0 Y2.0 Z3.0"},
		"parameters_16":           {"N100  G\"\"\"\"92\"\"\" X1.0 Y2.0 Z3.0", false, ""},
		"parameters_17":           {"N100  G\"\"\"92\"\"\"\" X1.0 Y2.0 Z3.0", false, ""},
		"parameters_18":           {"N100 G\"\"\"92\"\"\" X1.0 Y2.0 Z3.0", true, "N100 G\"\"\"92\"\"\" X1.0 Y2.0 Z3.0"},
		"parameters_19":           {"N100 G\"\"\"\"92\"\"\" X1.0 Y2.0 Z3.0", false, ""},
		"parameters_20":           {"N100 G\"\"\"92\"\"\"\" X1.0 Y2.0 Z3.0", false, ""},
		"parameters_21":           {"N100 G\"\" \"92\" \"\" X1.0 Y2.0 Z3.0", false, ""},
		"parameters_22":           {"N100 G\" \"\"92\"\" \" X1.0 Y2.0 Z3.0", true, "N100 G\" \"\"92\"\" \" X1.0 Y2.0 Z3.0"},
		"parameters_23":           {"N100 G\" \"\"92 \"\" \" X1.0 Y2.0 Z3.0", true, "N100 G\" \"\"92 \"\" \" X1.0 Y2.0 Z3.0"},
		"parameters_24":           {"N100 G\" \"\" 92 \"\" \" X1.0 Y2.0 Z3.0", true, "N100 G\" \"\" 92 \"\" \" X1.0 Y2.0 Z3.0"},
		"parameters_negatives_0":  {"N100 G92 X-1.0 Y2.0 Z3.0", true, "N100 G92 X-1.0 Y2.0 Z3.0"},
		"parameters_negatives_1":  {"N100 G92 1X1.0 Y2.0 Z3.0", false, ""},
		"parameters_negatives_2":  {"N100 *G92 X-1.0 Y-2.0 Z3.0", false, ""},
		"parameters_negatives_3":  {"N100  G92   X-1.0 Y-2.0 Z-3.0", true, "N100 G92 X-1.0 Y-2.0 Z-3.0"},
		"parameters_negatives_4":  {"N100 G92 X-1.0 Y2.0 Z3.0", true, "N100 G92 X-1.0 Y2.0 Z3.0"},
		"parameters_negatives_5":  {"N100 G 92 X-1.0 Y-2.0 Z3.0", false, ""},
		"parameters_negatives_6":  {"N100 \"G92 X-1.0 Y-2.0 Z-3.0", false, ""},
		"parameters_negatives_7":  {"N100  \"G92\" X-1.0 Y2.0 Z3.0", false, ""},
		"parameters_negatives_8":  {"N100 G\"92\" X1.0 Y-2.0 Z3.0", true, "N100 G\"92\" X1.0 Y-2.0 Z3.0"},
		"parameters_negatives_9":  {"N100 G\" 92\" X1.0 Y2.0 Z-3.0", true, "N100 G\" 92\" X1.0 Y2.0 Z-3.0"},
		"parameters_negatives_10": {"N100 G\"92 \" X-1.0 Y-2.0 Z3.0", true, "N100 G\"92 \" X-1.0 Y-2.0 Z3.0"},
		"parameters_negatives_11": {"N100 G\" 92 \" X-1.0 Y-2.0 Z-3.0", true, "N100 G\" 92 \" X-1.0 Y-2.0 Z-3.0"},
		"parameters_negatives_12": {"N100 G\"\"\"92\"\"\" X1.0 Y-2.0 Z-3.0", true, "N100 G\"\"\"92\"\"\" X1.0 Y-2.0 Z-3.0"},
		"parameters_negatives_13": {"N100 G\"\"\"\"92\"\"\" X1.0 Y-2.0 Z-3.0", false, ""},
		"parameters_negatives_14": {"N100 G\"\"\"92\"\"\"\" X1.0 Y2.0 Z-3.0", false, ""},
		"parameters_negatives_15": {"N100  G\"\"\"92\"\"\" X1.0 Y2.0 Z-3.0", true, "N100 G\"\"\"92\"\"\" X1.0 Y2.0 Z-3.0"},
		"parameters_negatives_16": {"N100  G\"\"\"\"92\"\"\" X1.0 Y2.0 Z-3.0", false, ""},
		"parameters_negatives_17": {"N100  G\"\"\"92\"\"\"\" X1.0 Y-2.0 Z3.0", false, ""},
		"parameters_negatives_18": {"N100 G\"\"\"92\"\"\" X1.0 Y-2.0 Z3.0", true, "N100 G\"\"\"92\"\"\" X1.0 Y-2.0 Z3.0"},
		"parameters_negatives_19": {"N100 G\"\"\"\"92\"\"\" X1.0 Y-2.0 Z3.0", false, ""},
		"parameters_negatives_20": {"N100 G\"\"\"92\"\"\"\" X1.0 Y2.0 Z-3.0", false, ""},
		"parameters_negatives_21": {"N100 G\"\" \"92\" \"\" X1.0 Y2.0 Z-3.0", false, ""},
		"parameters_negatives_22": {"N100 G\" \"\"92\"\" \" X1.0 Y-2.0 Z3.0", true, "N100 G\" \"\"92\"\" \" X1.0 Y-2.0 Z3.0"},
		"parameters_negatives_23": {"N100 G\" \"\"92 \"\" \" X-1.0 Y2.0 Z-3.0", true, "N100 G\" \"\"92 \"\" \" X-1.0 Y2.0 Z-3.0"},
		"parameters_negatives_24": {"N100 G\" \"\" 92 \"\" \" X1.0 Y-2.0 Z-3.0", true, "N100 G\" \"\" 92 \"\" \" X1.0 Y-2.0 Z-3.0"},
		"parameters_fail_0":       {"N100 G92 X1.0 Y2.0 Z 3.0", false, ""},
		"parameters_fail_1":       {"N100 G92 1X1.0 Y 2.0 Z3.0", false, ""},
		"parameters_fail_2":       {"N100 *G92 X1.0 Y2.0 Z3.0", false, ""},
		"parameters_fail_3":       {"N100  G92   X1.0 Y2.0 ZZ.0", false, ""},
		"parameters_fail_4":       {"N100 G92 X1.0 Y 2.0 Z3.0", false, ""},
		"parameters_fail_5":       {"N100 G 92 X1.0 Y2.0 Z3.0", false, ""},
		"parameters_fail_6":       {"N100 \"G92 X1 .0 Y2.0 Z3.0", false, ""},
		"parameters_fail_7":       {"N100  \"G92\" X1.0 Y2 .0 Z3.0", false, ""},
		"parameters_fail_8":       {"N100 G\"92\" X1.0 Y 2 .0 Z3.0", false, ""},
		"parameters_fail_9":       {"N100 G\" 92\" X1.0 Y 2.0 Z3.0", false, ""},
		"parameters_fail_10":      {"N100 G\"92 \" X1.0 Y2. 0 Z3.0", false, ""},
		"parameters_fail_11":      {"N100 G\" 92 \" X1.0 Y2.0 Z 3.0", false, ""},
		"parameters_fail_12":      {"N100 G\"\"\"92\"\"\" X1.0 Y 2.0 Z3.0", false, ""},
		"parameters_fail_13":      {"N100 G\"\"\"\"92\"\"\" X1 .0 Y2.0 Z3.0", false, ""},
		"parameters_fail_14":      {"N100 G\"\"\"92\"\"\"\" X1 .0 Y2.0 Z3.0", false, ""},
		"parameters_fail_15":      {"N100  G\"\"\"92\"\"\" X1.0 Y2. 0 Z3.0", false, ""},
		"parameters_fail_16":      {"N100  G\"\"\"\"92\"\"\" X1.0 Y2.0 Z3 .0", false, ""},
		"parameters_fail_17":      {"N100  G\"\"\"92\"\"\"\" X1.0 Y 2.0 Z 3.0", false, ""},
		"parameters_fail_18":      {"N100 G\"\"\"92\"\"\" X1.0 Y2 .0 Z 3.0", false, ""},
		"parameters_fail_19":      {"N100 G\"\"\"\"92\"\"\" X1.0 Y 2.0 Z3.0", false, ""},
		"parameters_fail_20":      {"N100 G\"\"\"92\"\"\"\" X1.0 Y2.0 Z 3.0", false, ""},
		"parameters_fail_21":      {"N100 G\"\" \"92\" \"\" X1.0 Y2.0 Z 3.0", false, ""},
		"parameters_fail_22":      {"N100 G\" \"\"92\"\" \" X1.0 Y2 .0 Z3.0", false, ""},
		"parameters_fail_23":      {"N100 G\" \"\"92 \"\" \" X1.0 Y2 .0 Z3.0", false, ""},
		"parameters_fail_24":      {"N100 G\" \"\" 92 \"\" \" X1.0 Y2.0 Z 3.0", false, ""},
		"checksum_0":              {"N100 G92 X1.0 Y2.0 Z3.0*10", true, "N100 G92 X1.0 Y2.0 Z3.0*10"},
		"checksum_1":              {"N100 G92 1X1.0 Y2.0 Z3.0*10", false, ""},
		"checksum_2":              {"N100 *G92 X1.0 Y2.0 Z3.0*10", false, ""},
		"checksum_3":              {"N100  G92   X1.0 Y2.0 Z3.0*10", true, "N100 G92 X1.0 Y2.0 Z3.0*10"},
		"checksum_4":              {"N100 G92 X1.0 Y2.0 Z3.0*10", true, "N100 G92 X1.0 Y2.0 Z3.0*10"},
		"checksum_5":              {"N100 G 92 X1.0 Y2.0 Z3.0*10", false, ""},
		"checksum_6":              {"N100 \"G92 X1.0 Y2.0 Z3.0*10", false, ""},
		"checksum_7":              {"N100  \"G92\" X1.0 Y2.0 Z3.0*10", false, ""},
		"checksum_8":              {"N100 G\"92\" X1.0 Y2.0 Z3.0*10", true, "N100 G\"92\" X1.0 Y2.0 Z3.0*10"},
		"checksum_9":              {"N100 G\" 92\" X1.0 Y2.0 Z3.0*10", true, "N100 G\" 92\" X1.0 Y2.0 Z3.0*10"},
		"checksum_10":             {"N100 G\"92 \" X1.0 Y2.0 Z3.0*10", true, "N100 G\"92 \" X1.0 Y2.0 Z3.0*10"},
		"checksum_11":             {"N100 G\" 92 \" X1.0 Y2.0 Z3.0*10", true, "N100 G\" 92 \" X1.0 Y2.0 Z3.0*10"},
		"checksum_12":             {"N100 G\"\"\"92\"\"\" X1.0 Y2.0 Z3.0*10", true, "N100 G\"\"\"92\"\"\" X1.0 Y2.0 Z3.0*10"},
		"checksum_13":             {"N100 G\"\"\"\"92\"\"\" X1.0 Y2.0 Z3.0*10", false, ""},
		"checksum_14":             {"N100 G\"\"\"92\"\"\"\" X1.0 Y2.0 Z3.0*10", false, ""},
		"checksum_15":             {"N100  G\"\"\"92\"\"\" X1.0 Y2.0 Z3.0*10", true, "N100 G\"\"\"92\"\"\" X1.0 Y2.0 Z3.0*10"},
		"checksum_16":             {"N100  G\"\"\"\"92\"\"\" X1.0 Y2.0 Z3.0*10", false, ""},
		"checksum_17":             {"N100  G\"\"\"92\"\"\"\" X1.0 Y2.0 Z3.0*10", false, ""},
		"checksum_18":             {"N100 G\"\"\"92\"\"\" X1.0 Y2.0 Z3.0*10", true, "N100 G\"\"\"92\"\"\" X1.0 Y2.0 Z3.0*10"},
		"checksum_19":             {"N100 G\"\"\"\"92\"\"\" X1.0 Y2.0 Z3.0*10", false, ""},
		"checksum_20":             {"N100 G\"\"\"92\"\"\"\" X1.0 Y2.0 Z3.0*10", false, ""},
		"checksum_21":             {"N100 G\"\" \"92\" \"\" X1.0 Y2.0 Z3.0*10", false, ""},
		"checksum_22":             {"N100 G\" \"\"92\"\" \" X1.0 Y2.0 Z3.0*10", true, "N100 G\" \"\"92\"\" \" X1.0 Y2.0 Z3.0*10"},
		"checksum_23":             {"N100 G\" \"\"92 \"\" \" X1.0 Y2.0 Z3.0*10", true, "N100 G\" \"\"92 \"\" \" X1.0 Y2.0 Z3.0*10"},
		"checksum_24":             {"N100 G\" \"\" 92 \"\" \" X1.0 Y2.0 Z3.0*10", true, "N100 G\" \"\" 92 \"\" \" X1.0 Y2.0 Z3.0*10"},
		"checksum_25":             {"N100 G92 X-1.0 Y2.0 Z3.0*10", true, "N100 G92 X-1.0 Y2.0 Z3.0*10"},
		"checksum_26":             {"N100 G92 1X1.0 Y2.0 Z3.0*10", false, ""},
		"checksum_27":             {"N100 *G92 X-1.0 Y-2.0 Z3.0*10", false, ""},
		"checksum_28":             {"N100  G92   X-1.0 Y-2.0 Z-3.0*10", true, "N100 G92 X-1.0 Y-2.0 Z-3.0*10"},
		"checksum_29":             {"N100 G92 X-1.0 Y2.0 Z3.0*10", true, "N100 G92 X-1.0 Y2.0 Z3.0*10"},
		"checksum_30":             {"N100 G 92 X-1.0 Y-2.0 Z3.0*10", false, ""},
		"checksum_31":             {"N100 \"G92 X-1.0 Y-2.0 Z-3.0*10", false, ""},
		"checksum_32":             {"N100  \"G92\" X-1.0 Y2.0 Z3.0*10", false, ""},
		"checksum_33":             {"N100 G\"92\" X1.0 Y-2.0 Z3.0*10", true, "N100 G\"92\" X1.0 Y-2.0 Z3.0*10"},
		"checksum_34":             {"N100 G\" 92\" X1.0 Y2.0 Z-3.0*10", true, "N100 G\" 92\" X1.0 Y2.0 Z-3.0*10"},
		"checksum_35":             {"N100 G\"92 \" X-1.0 Y-2.0 Z3.0*10", true, "N100 G\"92 \" X-1.0 Y-2.0 Z3.0*10"},
		"checksum_36":             {"N100 G\" 92 \" X-1.0 Y-2.0 Z-3.0*10", true, "N100 G\" 92 \" X-1.0 Y-2.0 Z-3.0*10"},
		"checksum_37":             {"N100 G\"\"\"92\"\"\" X1.0 Y-2.0 Z-3.0*10", true, "N100 G\"\"\"92\"\"\" X1.0 Y-2.0 Z-3.0*10"},
		"checksum_38":             {"N100 G\"\"\"\"92\"\"\" X1.0 Y-2.0 Z-3.0*10", false, ""},
		"checksum_39":             {"N100 G\"\"\"92\"\"\"\" X1.0 Y2.0 Z-3.0*10", false, ""},
		"checksum_40":             {"N100  G\"\"\"92\"\"\" X1.0 Y2.0 Z-3.0*10", true, "N100 G\"\"\"92\"\"\" X1.0 Y2.0 Z-3.0*10"},
		"checksum_41":             {"N100  G\"\"\"\"92\"\"\" X1.0 Y2.0 Z-3.0*10", false, ""},
		"checksum_42":             {"N100  G\"\"\"92\"\"\"\" X1.0 Y-2.0 Z3.0*10", false, ""},
		"checksum_43":             {"N100 G\"\"\"92\"\"\" X1.0 Y-2.0 Z3.0*10", true, "N100 G\"\"\"92\"\"\" X1.0 Y-2.0 Z3.0*10"},
		"checksum_44":             {"N100 G\"\"\"\"92\"\"\" X1.0 Y-2.0 Z3.0*10", false, ""},
		"checksum_45":             {"N100 G\"\"\"92\"\"\"\" X1.0 Y2.0 Z-3.0*10", false, ""},
		"checksum_46":             {"N100 G\"\" \"92\" \"\" X1.0 Y2.0 Z-3.0*10", false, ""},
		"checksum_47":             {"N100 G\" \"\"92\"\" \" X1.0 Y-2.0 Z3.0*10", true, "N100 G\" \"\"92\"\" \" X1.0 Y-2.0 Z3.0*10"},
		"checksum_48":             {"N100 G\" \"\"92 \"\" \" X-1.0 Y2.0 Z-3.0*10", true, "N100 G\" \"\"92 \"\" \" X-1.0 Y2.0 Z-3.0*10"},
		"checksum_49":             {"N100 G\" \"\" 92 \"\" \" X1.0 Y-2.0 Z-3.0*10", true, "N100 G\" \"\" 92 \"\" \" X1.0 Y-2.0 Z-3.0*10"},
		"checksum_fail_1":         {"N100 G\" \"\"92\"\" \" X1.0 Y-2.0 Z3.0 *10.3", false, ""},
		"checksum_fail_2":         {"N100 G\" \"\"92 \"\" \" X-1.0 Y2.0 Z-3.0 *10.0", false, ""},
		"checksum_fail_3":         {"N100 G\" \"\" 92 \"\" \" X1.0 Y-2.0 Z-3.0 *-10", false, ""},
		"checksum_fail_4":         {"N100 G\" \"\"92\"\" \" X1.0 Y-2.0 Z3.0 *-10.3", false, ""},
		"checksum_fail_5":         {"N100 G\" \"\"92\"\" \" X1.0 Y-2.0 Z3.0 *10", false, ""},
		"checksum_fail_6":         {"N100 G\" \"\"92 \"\" \" X-1.0 Y2.0 Z-3.0 *10", false, ""},
		"checksum_fail_7":         {"N100 G\" \"\" 92 \"\" \" X1.0 Y-2.0 Z-3.0 *10", false, ""},
		"checksum_fail_8":         {"N100 G\" \"\"92\"\" \" X1.0 Y-2.0 Z3.0 *10", false, ""},
		"comments_0":              {"N100 G92 X1.0 Y2.0 Z3.0*10;lorem ipsu", true, "N100 G92 X1.0 Y2.0 Z3.0*10 ;lorem ipsu"},
		"comments_1":              {"N100 G92 1X1.0 Y2.0 Z3.0*10;lorem ipsu", false, ""},
		"comments_2":              {"N100 *G92 X1.0 Y2.0 Z3.0*10;lorem ipsu", false, ""},
		"comments_3":              {"N100  G92   X1.0 Y2.0 Z3.0*10;lorem ipsu", true, "N100 G92 X1.0 Y2.0 Z3.0*10 ;lorem ipsu"},
		"comments_4":              {"N100 G92 X1.0 Y2.0 Z3.0*10;lorem ipsu", true, "N100 G92 X1.0 Y2.0 Z3.0*10 ;lorem ipsu"},
		"comments_5":              {"N100 G 92 X1.0 Y2.0 Z3.0*10;lorem ipsu", false, ""},
		"comments_6":              {"N100 \"G92 X1.0 Y2.0 Z3.0*10;lorem ipsu", false, ""},
		"comments_7":              {"N100  \"G92\" X1.0 Y2.0 Z3.0*10;lorem ipsu", false, ""},
		"comments_8":              {"N100 G\"92\" X1.0 Y2.0 Z3.0*10;lorem ipsu", true, "N100 G\"92\" X1.0 Y2.0 Z3.0*10 ;lorem ipsu"},
		"comments_9":              {"N100 G\" 92\" X1.0 Y2.0 Z3.0*10;lorem ipsu", true, "N100 G\" 92\" X1.0 Y2.0 Z3.0*10 ;lorem ipsu"},
		"comments_10":             {"N100 G\"92 \" X1.0 Y2.0 Z3.0*10;lorem ipsu", true, "N100 G\"92 \" X1.0 Y2.0 Z3.0*10 ;lorem ipsu"},
		"comments_11":             {"N100 G\" 92 \" X1.0 Y2.0 Z3.0*10;lorem ipsu", true, "N100 G\" 92 \" X1.0 Y2.0 Z3.0*10 ;lorem ipsu"},
		"comments_12":             {"N100 G\"\"\"92\"\"\" X1.0 Y2.0 Z3.0*10;lorem ipsu", true, "N100 G\"\"\"92\"\"\" X1.0 Y2.0 Z3.0*10 ;lorem ipsu"},
		"comments_13":             {"N100 G\"\"\"\"92\"\"\" X1.0 Y2.0 Z3.0*10;lorem ipsu", false, ""},
		"comments_14":             {"N100 G\"\"\"92\"\"\"\" X1.0 Y2.0 Z3.0*10;lorem ipsu", false, ""},
		"comments_15":             {"N100  G\"\"\"92\"\"\" X1.0 Y2.0 Z3.0*10;lorem ipsu", true, "N100 G\"\"\"92\"\"\" X1.0 Y2.0 Z3.0*10 ;lorem ipsu"},
		"comments_16":             {"N100  G\"\"\"\"92\"\"\" X1.0 Y2.0 Z3.0*10;lorem ipsu", false, ""},
		"comments_17":             {"N100  G\"\"\"92\"\"\"\" X1.0 Y2.0 Z3.0*10;lorem ipsu", false, ""},
		"comments_18":             {"N100 G\"\"\"92\"\"\" X1.0 Y2.0 Z3.0*10;lorem ipsu", true, "N100 G\"\"\"92\"\"\" X1.0 Y2.0 Z3.0*10 ;lorem ipsu"},
		"comments_19":             {"N100 G\"\"\"\"92\"\"\" X1.0 Y2.0 Z3.0*10;lorem ipsu", false, ""},
		"comments_20":             {"N100 G\"\"\"92\"\"\"\" X1.0 Y2.0 Z3.0*10;lorem ipsu", false, ""},
		"comments_21":             {"N100 G\"\" \"92\" \"\" X1.0 Y2.0 Z3.0*10;lorem ipsu", false, ""},
		"comments_22":             {"N100 G\" \"\"92\"\" \" X1.0 Y2.0 Z3.0*10;lorem ipsu", true, "N100 G\" \"\"92\"\" \" X1.0 Y2.0 Z3.0*10 ;lorem ipsu"},
		"comments_23":             {"N100 G\" \"\"92 \"\" \" X1.0 Y2.0 Z3.0*10;lorem ipsu", true, "N100 G\" \"\"92 \"\" \" X1.0 Y2.0 Z3.0*10 ;lorem ipsu"},
		"comments_24":             {"N100 G\" \"\" 92 \"\" \" X1.0 Y2.0 Z3.0*10;lorem ipsu", true, "N100 G\" \"\" 92 \"\" \" X1.0 Y2.0 Z3.0*10 ;lorem ipsu"},
		"comments_25":             {"N100 G92 X-1.0 Y2.0 Z3.0*10;lorem ipsu", true, "N100 G92 X-1.0 Y2.0 Z3.0*10 ;lorem ipsu"},
		"comments_26":             {"N100 G92 1X1.0 Y2.0 Z3.0*10;lorem ipsu", false, ""},
		"comments_27":             {"N100 *G92 X-1.0 Y-2.0 Z3.0*10;lorem ipsu", false, ""},
		"comments_28":             {"N100  G92   X-1.0 Y-2.0 Z-3.0*10;lorem ipsu", true, "N100 G92 X-1.0 Y-2.0 Z-3.0*10 ;lorem ipsu"},
		"comments_29":             {"N100 G92 X-1.0 Y2.0 Z3.0*10;lorem ipsu", true, "N100 G92 X-1.0 Y2.0 Z3.0*10 ;lorem ipsu"},
		"comments_30":             {"N100 G 92 X-1.0 Y-2.0 Z3.0*10;lorem ipsu", false, ""},
		"comments_31":             {"N100 \"G92 X-1.0 Y-2.0 Z-3.0*10;lorem ipsu", false, ""},
		"comments_32":             {"N100  \"G92\" X-1.0 Y2.0 Z3.0*10;lorem ipsu", false, ""},
		"comments_33":             {"N100 G\"92\" X1.0 Y-2.0 Z3.0*10;lorem ipsu", true, "N100 G\"92\" X1.0 Y-2.0 Z3.0*10 ;lorem ipsu"},
		"comments_34":             {"N100 G\" 92\" X1.0 Y2.0 Z-3.0*10;lorem ipsu", true, "N100 G\" 92\" X1.0 Y2.0 Z-3.0*10 ;lorem ipsu"},
		"comments_35":             {"N100 G\"92 \" X-1.0 Y-2.0 Z3.0*10;lorem ipsu", true, "N100 G\"92 \" X-1.0 Y-2.0 Z3.0*10 ;lorem ipsu"},
		"comments_36":             {"N100 G\" 92 \" X-1.0 Y-2.0 Z-3.0*10;lorem ipsu", true, "N100 G\" 92 \" X-1.0 Y-2.0 Z-3.0*10 ;lorem ipsu"},
		"comments_37":             {"N100 G\"\"\"92\"\"\" X1.0 Y-2.0 Z-3.0*10;lorem ipsu", true, "N100 G\"\"\"92\"\"\" X1.0 Y-2.0 Z-3.0*10 ;lorem ipsu"},
		"comments_38":             {"N100 G\"\"\"\"92\"\"\" X1.0 Y-2.0 Z-3.0*10;lorem ipsu", false, ""},
		"comments_39":             {"N100 G\"\"\"92\"\"\"\" X1.0 Y2.0 Z-3.0*10;lorem ipsu", false, ""},
		"comments_40":             {"N100  G\"\"\"92\"\"\" X1.0 Y2.0 Z-3.0*10;lorem ipsu", true, "N100 G\"\"\"92\"\"\" X1.0 Y2.0 Z-3.0*10 ;lorem ipsu"},
		"comments_41":             {"N100  G\"\"\"\"92\"\"\" X1.0 Y2.0 Z-3.0*10;lorem ipsu", false, ""},
		"comments_42":             {"N100  G\"\"\"92\"\"\"\" X1.0 Y-2.0 Z3.0*10;lorem ipsu", false, ""},
		"comments_43":             {"N100 G\"\"\"92\"\"\" X1.0 Y-2.0 Z3.0*10;lorem ipsu", true, "N100 G\"\"\"92\"\"\" X1.0 Y-2.0 Z3.0*10 ;lorem ipsu"},
		"comments_44":             {"N100 G\"\"\"\"92\"\"\" X1.0 Y-2.0 Z3.0*10;lorem ipsu", false, ""},
		"comments_45":             {"N100 G\"\"\"92\"\"\"\" X1.0 Y2.0 Z-3.0*10;lorem ipsu", false, ""},
		"comments_46":             {"N100 G\"\" \"92\" \"\" X1.0 Y2.0 Z-3.0*10;lorem ipsu", false, ""},
		"comments_47":             {"N100 G\" \"\"92\"\" \" X1.0 Y-2.0 Z3.0*10;lorem ipsu", true, "N100 G\" \"\"92\"\" \" X1.0 Y-2.0 Z3.0*10 ;lorem ipsu"},
		"comments_48":             {"N100 G\" \"\"92 \"\" \" X-1.0 Y2.0 Z-3.0*10;lorem ipsu", true, "N100 G\" \"\"92 \"\" \" X-1.0 Y2.0 Z-3.0*10 ;lorem ipsu"},
		"comments_49":             {"N100 G\" \"\" 92 \"\" \" X1.0 Y-2.0 Z-3.0*10;lorem ipsu", true, "N100 G\" \"\" 92 \"\" \" X1.0 Y-2.0 Z-3.0*10 ;lorem ipsu"},
		"comments_fail_1":         {"N100 G\" \"\"92\"\" \" X1.0 Y-2.0 Z3.0 *10.3 ; lorem ipsu", false, ""},
		"comments_fail_2":         {"N100 G\" \"\"92 \"\" \" X-1.0 Y2.0 Z-3.0 *10.0 ;;lorem ipsulorem ipsu", false, ""},
		"comments_fail_3":         {"N100 G\" \"\" 92 \"\" \" X1.0 Y-2.0 Z-3.0 *-10 ;lorem ipsu", false, ""},
		"comments_fail_4":         {"N100 G\" \"\"92\"\" \" X1.0 Y-2.0 Z3.0 *-10.3 ;lorem ipsu", false, ""},
		"comments_fail_5":         {"N100 G\" \"\"92\"\" \" X1.0 Y-2.0 Z3.0 *10 ;lorem ipsu", false, ""},
		"comments_fail_6":         {"N100 G\" \"\"92 \"\" \" X-1.0 Y2.0 Z-3.0 *10 ;lorem ipsu", false, ""},
		"comments_fail_7":         {"N100 G\" \"\" 92 \"\" \" X1.0 Y-2.0 Z-3.0 *10 ;lorem ipsu", false, ""},
		"comments_fail_8":         {"N100 G\" \"\"92\"\" \" X1.0 Y-2.0 Z3.0 *10", false, ""},
		"special_0":               {"N92", false, ""},
		"special_1":               {"G\"\"\"92\"\"\" X1.0 Y2.0 Z3.0 G\"\"\"92\"\"\"", true, "G\"\"\"92\"\"\" X1.0 Y2.0 Z3.0 G\"\"\"92\"\"\""},
		"special_2":               {"N2.3 G21", false, ""},
		"special_3":               {"N2 K21", false, ""},
	}

	for name, tc := range cases {
		t.Run(fmt.Sprintf("%s[%s]", name, tc.input), func(t *testing.T) {
			gb, err := Parse(tc.input, func(config block.BlockParserConfigurer) error {

				err := config.SetGcodeFactory(mockGcodeFactory)
				if err != nil {
					t.Errorf("cann't save mockGcodeFactory: %v", err)
				}

				err = config.SetHash(mockHash)
				if err != nil {
					t.Errorf("cann't save mockHash: %v", err)
				}

				return nil
			})

			if tc.valid {
				if err != nil {
					t.Errorf("got error %v, want error nil", err)
				}

				if gb == nil {
					t.Errorf("got gcodeBlock nil, want gcodeBlock not nil")
					return
				}

				if gb.ToLine("%l %c %p%k %m") != tc.output {
					t.Errorf("got gcodeBlock (%d)[%s], want gcodeBlock: (%d)[%s]", len(gb.ToLine("%l %c %p%k %m")), gb.ToLine("%l %c %p%k %m"), len(tc.output), tc.output)
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

func TestParse_ConfigurationError_first(t *testing.T) {
	gb, err := Parse("N1 G1", func(config block.BlockParserConfigurer) error {
		return fmt.Errorf("something went wrong")
	})
	if err == nil {
		t.Errorf("got error nil, want error not nil")
	}
	if gb != nil {
		t.Errorf("got gcodeBlock not nil, want gcodeBlock nil")
	}
}

func TestGcodeblogk_Parameters(t *testing.T) {

	var cases = [1]struct {
		source     string
		parameters []string
	}{
		{
			source:     "N7 G1 X2.0 Y2.0 F3000.0",
			parameters: []string{"X2.0", "Y2.0", "F3000.0"},
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("(%v)", i), func(t *testing.T) {
			b, err := Parse(tc.source)
			if err != nil {
				t.Errorf("got %v, want nil error", err)
				return
			}
			if b == nil {
				t.Errorf("got nil block, want %v", tc.source)
				return
			}

			if len(b.Parameters()) != len(tc.parameters) {
				t.Errorf("got parameters size %d, want parameters size %d", len(b.Parameters()), len(tc.parameters))
				return
			}

			match := true
			for i, s := range tc.parameters {
				if b.Parameters()[i].String() != s {
					match = false
					break
				}
			}

			if !match {
				t.Errorf("got %v, want %v", b.ToLine("%p"), tc.parameters)
			}
		})
	}

}

func TestGcodeblock_Calculate(t *testing.T) {

	var cases = [1]struct {
		source        string
		checksumValue uint32
	}{
		{
			source:        "N7 G1 X2.0 Y2.0 F3000.0",
			checksumValue: 85,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("(%v)", i), func(t *testing.T) {
			b, err := Parse(tc.source)
			if err != nil {
				t.Errorf("got %v, want nil error", err)
				return
			}
			if b == nil {
				t.Errorf("got nil block, want %v", tc.source)
				return
			}

			gc, err := b.CalculateChecksum()
			if err != nil {
				t.Errorf("got error %v, want error nil", err)
				return
			}

			if gc.Address() != tc.checksumValue {
				t.Errorf("got checksum value %d, want checksum value %d", gc.Address(), tc.checksumValue)
			}
		})
	}
}

func TestGcodeblock_Verify(t *testing.T) {

	var cases = [1]struct {
		source string
		err    bool
		ok     bool
	}{
		{
			source: "N7 G1 X2.0 Y2.0 F3000.0",
			err:    true,
			ok:     false,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("(%v)", i), func(t *testing.T) {
			b, err := Parse(tc.source)
			if err != nil {
				t.Errorf("got %v, want nil error", err)
				return
			}
			if b == nil {
				t.Errorf("got nil block, want %v", tc.source)
				return
			}

			ok, err := b.VerifyChecksum()
			if ok != tc.ok || (err != nil) != tc.err {
				t.Errorf("got error %v verified %v, want error %v verified %v", err, ok, tc.err, tc.ok)
			}
		})
	}
}

func TestGcodeblock_ToLineError(t *testing.T) {

	mockCommand, err := addressablegcode.New[int32]('G', 92)
	if err != nil {
		t.Errorf("got error not nil, want error nil: %v", err)
	}

	t.Run("params slice empty", func(t *testing.T) {

		b, err := New(mockCommand, func(config block.BlockConstructorConfigurer) error {
			params := []gcode.Gcoder{}
			err := config.SetParameters(params)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			t.Errorf("got %v, want nil error", err)
			return
		}

		a := b.ToLine("%l")
		if a != "" {
			t.Errorf("got (%d)[%s], want (0)[]", len(a), a)
			return
		}
	})
}
