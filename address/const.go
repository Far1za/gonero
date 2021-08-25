package address

import "filippo.io/edwards25519"

const (
	MaxUint64 = 1<<64 - 1
)

var (
	MainAddr = byte(0x12)
	IntAddr  = byte(0x13)
	SubAddr  = byte(0x2A)
	G        = edwards25519.NewGeneratorPoint()
)
