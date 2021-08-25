package utils

// NOTE you need this function to flip the bytes in the right order
// https://cs.opensource.google/go/go/+/refs/tags/go1.17:src/encoding/binary/binary.go;l=68
func LittleEndian(v uint32) []byte {
	var b [4]byte
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	return b[:]
}

// per byte comparison on given slices
func TestBytes(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// https://developpaper.com/explain-the-principle-of-varint-coding-in-detail/
// encode integers to bytes using varint algorithm
func EncodeVarint(x uint64) []byte {
	var buf [10]byte
	var n int
	for n = 0; x > 127; n++ {
		//X & 0x7F means to take out the lower 7bit data, and 0x80 means to add 1 to the highest bit
		buf[n] = 0x80 | uint8(x&0x7F)
		//Move 7 bits to the right to continue the following data processing
		x >>= 7
	}
	//Last byte data
	buf[n] = uint8(x)
	n++
	return buf[0:n]
}

// decode varint encoded integers
func DecodeVarint(buf []byte) (x uint64, n int) {
	for shift := uint(0); shift < 64; shift += 7 {
		if n >= len(buf) {
			return 0, 0
		}
		b := uint64(buf[n])
		n++
		//This is divided into three steps:
		//1: B & 0x7F get the lower 7bits valid data
		//2: (B & 0x7F) < shift because it is a small end sequence, each time a byte data is processed, it needs to move 7bits to the high order
		//3: put the data X together with the current byte data
		x |= (b & 0x7F) << shift
		if (b & 0x80) == 0 {
			return x, n
		}
	}

	// The number is too large to represent in a 64-bit value.
	return 0, 0
}
