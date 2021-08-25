package address

import (
	"crypto/rand"
	"filippo.io/edwards25519"
	"golang.org/x/crypto/sha3"
	"gonero/scalar"
	"gonero/utils"
)

type KeyPair struct {
	spendKey, viewKey *edwards25519.Scalar
}

func (k *KeyPair) GetSpendKey() []byte {
	return k.spendKey.Bytes()
}

func (k *KeyPair) GetViewKey() []byte {
	return k.viewKey.Bytes()
}

func (k *KeyPair) GetViewKeyAddr() *edwards25519.Scalar {
	return k.viewKey
}

// create new keypair or load existing one from private seed
// NOTE seed is the private spend key
func NewKeyPair(seed []byte) (*KeyPair, error) {
	r := new([64]byte)
	s := new([32]byte)
	k := new(KeyPair)
	k.spendKey = edwards25519.NewScalar()
	k.viewKey = edwards25519.NewScalar()
	if seed == nil {
		rand.Read(r[:])
		scalar.Reduce(s, r)
	} else {
		copy(s[:], seed)
	}
	if _, err := k.spendKey.SetCanonicalBytes(s[:]); err != nil {
		return nil, err
	}

	h := sha3.NewLegacyKeccak256()
	h.Write(k.GetSpendKey())
	copy(s[:], h.Sum(nil))
	scalar.Reduce32(s)
	if _, err := k.viewKey.SetCanonicalBytes(s[:]); err != nil {
		return nil, err
	}

	return k, nil
}

func (k *KeyPair) GetPublicKeyPair() *PublicKeyPair {
	p := new(PublicKeyPair)
	p.spendKey = edwards25519.NewIdentityPoint()
	p.viewKey = edwards25519.NewIdentityPoint()
	p.spendKey.ScalarMult(k.spendKey, G)
	p.viewKey.ScalarMult(k.viewKey, G)
	return p
}

// using the seed/initial keypair get/generate subaddress keypair for index i>=1
func (k *KeyPair) GetSubKeyPair(acc, index uint32) (*KeyPair, error) {
	S := new([32]byte)
	if acc+index == 0 {
		return k, nil
	}
	s := new(KeyPair)
	s.spendKey = edwards25519.NewScalar()
	s.viewKey = edwards25519.NewScalar()
	a := []byte("SubAddr")
	a = append(a, byte(0x00))
	a = append(a, k.GetViewKey()...)
	a = append(a, utils.LittleEndian(acc)...)
	a = append(a, utils.LittleEndian(index)...)
	h := sha3.NewLegacyKeccak256()
	h.Write(a)
	copy(S[:], h.Sum(nil))
	scalar.Reduce32(S)
	if _, err := s.spendKey.SetCanonicalBytes(S[:]); err != nil {
		return nil, err
	}
	s.spendKey.Add(s.spendKey, k.spendKey)
	s.viewKey.Multiply(k.viewKey, s.spendKey)

	return s, nil
}
