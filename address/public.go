package address

import (
	"crypto/rand"
	"filippo.io/edwards25519"
	"fmt"
	"golang.org/x/crypto/sha3"
	"gonero/base58"
	"gonero/utils"
	"math/big"
)

type PublicKeyPair struct {
	spendKey, viewKey *edwards25519.Point
}

func (k *PublicKeyPair) GetSpendKey() []byte {
	return k.spendKey.Bytes()
}

func (k *PublicKeyPair) GetViewKey() []byte {
	return k.viewKey.Bytes()
}

// get the public keypair for a given address
func LoadPublicKeyPair(addr string) (*PublicKeyPair, error) {
	p := new(PublicKeyPair)
	p.spendKey = edwards25519.NewIdentityPoint()
	p.viewKey = edwards25519.NewIdentityPoint()
	a := base58.Decode(addr)
	h := sha3.NewLegacyKeccak256()
	h.Write(a[:len(a)-4])
	s := h.Sum(nil)
	if !utils.TestBytes(a[len(a)-4:], s[:4]) {
		return nil, fmt.Errorf("address checksum fail")
	}
	if _, err := p.spendKey.SetBytes(a[1:33]); err != nil {
		return nil, err
	}
	if _, err := p.viewKey.SetBytes(a[33:65]); err != nil {
		return nil, err
	}
	return p, nil
}

func (k *PublicKeyPair) GetPublicAddress(network byte) string {
	var a []byte
	a = append([]byte{network}, k.spendKey.Bytes()...)
	a = append(a, k.viewKey.Bytes()...)
	h := sha3.NewLegacyKeccak256()
	h.Write(a)
	s := h.Sum(nil)
	a = append(a, s[:4]...)
	return base58.Encode(a)
}

// generates new integrated address using the given address
// network byte + public spend key + public view key + 64-bit payment ID + checksum
// NewIntAddress([]byte{IntAddr}, a.GetSpendKey(), a.GetViewKey())
func (k *PublicKeyPair) NewIntAddress(network byte) (string, error) {
	//var addr string
	var a []byte
	a = append([]byte{network}, k.spendKey.Bytes()...)
	a = append(a, k.viewKey.Bytes()...)
	i := new(big.Int)
	r, err := rand.Int(rand.Reader, i.SetUint64(MaxUint64))
	if err != nil {
		return "", err
	}
	a = append(a, r.Bytes()...)
	h := sha3.NewLegacyKeccak256()
	h.Write(a)
	s := h.Sum(nil)
	a = append(a, s[:4]...)
	return base58.Encode(a), nil
}
