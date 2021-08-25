package address

import (
	"crypto/rand"
	"filippo.io/edwards25519"
	"golang.org/x/crypto/sha3"
	"gonero/scalar"
)

// generate new random value used for stealth addresses generation
func NewRandom() (*edwards25519.Scalar, error) {
	s := edwards25519.NewScalar()
	r := new([32]byte)
	rand.Read(r[:])
	h := sha3.NewLegacyKeccak256()
	h.Write(r[:])
	copy(r[:], h.Sum(nil))
	scalar.Reduce32(r)
	return s.SetCanonicalBytes(r[:])
}
