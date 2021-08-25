package address

import (
	"filippo.io/edwards25519"
	"golang.org/x/crypto/sha3"
	"gonero/scalar"
	"gonero/utils"
)

// create new stealth address using the loaded public key pair
func (k *PublicKeyPair) NewStealthAddress(r *edwards25519.Scalar, t uint64) ([]byte, error) {
	// P = H(8aR||i)G + B
	p := edwards25519.NewIdentityPoint()
	s := edwards25519.NewScalar()
	p.MultByCofactor(k.viewKey)
	p.ScalarMult(r, p)
	hs := new([32]byte)
	var a []byte
	a = append(a, p.Bytes()...)
	a = append(a, utils.EncodeVarint(t)...)
	h := sha3.NewLegacyKeccak256()
	h.Write(a)
	copy(hs[:], h.Sum(nil))
	scalar.Reduce32(hs)
	if _, err := s.SetCanonicalBytes(hs[:]); err != nil {
		return nil, err
	}
	p.ScalarMult(s, G)
	p.Add(p, k.spendKey)

	return p.Bytes(), nil
}

// used to scan the blockhain for incomming outputs
// r here is the public key for the transaction
// v is the secret view key when dealing with subaddress
func (k *KeyPair) ProcessStealthAddress(v *edwards25519.Scalar, r *edwards25519.Point, t uint64) (map[string][]byte, error) {
	// P = H(8aR||i) + B
	p := edwards25519.NewIdentityPoint()
	s := edwards25519.NewScalar()
	p.MultByCofactor(r)
	if v != nil {
		p.ScalarMult(v, p)
	} else {
		p.ScalarMult(k.viewKey, p)
	}
	hs := new([32]byte)
	var a []byte
	a = append(a, p.Bytes()...)
	a = append(a, utils.EncodeVarint(t)...)
	h := sha3.NewLegacyKeccak256()
	h.Write(a)
	copy(hs[:], h.Sum(nil))
	scalar.Reduce32(hs)
	if _, err := s.SetCanonicalBytes(hs[:]); err != nil {
		return nil, err
	}
	// this step generates x that is needed for key image later
	s.Add(s, k.spendKey)
	// this step calculates P' that should be the same as Destination Key
	// in Tx Output
	p.ScalarMult(s, G)
	res := make(map[string][]byte)
	res["x"] = s.Bytes()
	res["P'"] = p.Bytes()

	return res, nil
}

// create main address transaction key
func GetTranKey(r *edwards25519.Scalar) *edwards25519.Point {
	R := edwards25519.NewIdentityPoint()
	R.ScalarMult(r, G)
	return R
}

// create subaddress transaction key
func (k *PublicKeyPair) GetSubTranKey(r *edwards25519.Scalar) *edwards25519.Point {
	R := edwards25519.NewIdentityPoint()
	R.ScalarMult(r, k.spendKey)
	return R
}
