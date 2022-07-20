package checksum_test

import (
	"fmt"

	"github.com/mauroalderete/gcode-core/checksum"
)

func Example() {
	h := checksum.New()

	const line = "N4 G92 E0" //N4 G92 E0*67

	h.Reset()
	_, err := h.Write([]byte(line))
	if err != nil {
		_ = fmt.Errorf("failed to calculate hash of the line %s: %w", line, err)
	}

	var result = int32(h.Sum(nil)[0])

	fmt.Printf("Hash is: %d", result)

	// Output:
	// Hash is: 67
}

func Example_second() {
	h := checksum.New()

	const line = "N4 G92 E0" //N4 G92 E0*67

	h.Reset()
	// iteration for each item of the string
	// the checksum to be constructed gradually
	for _, b := range []byte(line) {
		_, err := h.Write([]byte{b})
		if err != nil {
			_ = fmt.Errorf("failed to calculate hash of the line %s: %w", line, err)
		}
	}

	var result = int32(h.Sum(nil)[0])

	fmt.Printf("Hash is: %d", result)

	// Output:
	// Hash is: 67
}
