// checksum package implement the checksum algorithm used by Marlin and RepRap firmwares.
// to validate a single block line at gcode file.
//
// For more information about checksum algorithm visit [Checksum algorithm]
//
// It is compatible with [hash.Hash] interface
//
// [Checksum algorithm]: https://reprap.org/wiki/G-code#.2A:_Checksum
// [hash.Hash]: https://pkg.go.dev/hash@go1.18.3
package checksum

import "hash"

//region hash implementation

// digest represents the partial evaluation of a checksum
type digest struct {
	checksum uint8
}

// Sum appends the current hash to b and returns the resulting slice.
// It does not change the underlying hash state.
func (d *digest) Sum(in []byte) []byte {
	return append(in, byte(d.checksum))
}

// Reset resets the Hash to its initial state.
func (d *digest) Reset() {
	d.checksum = 0
}

// Size returns the number of bytes Sum will return.
func (d *digest) Size() int {
	return 1
}

// BlockSize returns the hash's underlying block size.
// The Write method must be able to accept any amount
// of data, but it may operate more efficiently if all writes
// are a multiple of the block size.
func (d *digest) BlockSize() int {
	return 1
}

// Write (via the embedded io.Writer interface) adds more data to the running hash.
// It never returns an error.
func (d *digest) Write(p []byte) (n int, err error) {
	d.checksum = simpleUpdate(d.checksum, p)
	return len(p), nil
}

//#endregion
//#region constructors

// New creates a new hash.Hash computing checksum using the Marlin and RepRap algorithm.
func New() hash.Hash {
	return &digest{}
}

//#endregion
//#region private functions

// simpleUpdate uses a simple algorithm to update the checksum, given a checksum previously computed
func simpleUpdate(checksum uint8, p []byte) uint8 {
	for _, v := range p {
		checksum ^= uint8(v)
	}
	return checksum
}

//#endregion
