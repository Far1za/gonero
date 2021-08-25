package main

import (
	//"encoding/hex"
	"fmt"
	"gonero/address"
)

func main() {
	//a := "c595161ea20ccd8c692947c2d3ced471e9b13a18b150c881232794e8042bf107"
	//b := "43tXwm6UNNvSyMdHU4Jfeg4GRgU7KEVAfHo3B5RrXYMjZMRaowr68y12HSo14wv2qcYqqpG1U5AHrJtBdFHKPDEA9UxK6Hy"
	//decoded, err := hex.DecodeString(a)
	//if err != nil {
	//	fmt.Println(err)
	//}
	k, _ := address.NewKeyPair(nil)
	//k, err := address.NewKeyPair(decoded)
	//if err != nil {
	//	fmt.Println(err)
	//}
	p := k.GetPublicKeyPair()
	// public keypair
	//p, err := address.LoadPublicKeyPair(b)
	//if err != nil {
	//	fmt.Println(err)
	//}
	fmt.Printf("====================\nPrivateSpendKey: %x\nPrivateViewKey: %x\n", k.GetSpendKey(), k.GetViewKey())
	fmt.Printf("====================\nPublicSpendKey: %x\nPublicViewKey: %x\n", p.GetSpendKey(), p.GetViewKey())
	fmt.Println("PublicAddress:", p.GetPublicAddress(address.MainAddr))
	ia, _ := p.NewIntAddress(address.IntAddr)
	fmt.Println("IntegratedAddress:", ia)
	s, _ := k.GetSubKeyPair(3, 1000)
	ps := s.GetPublicKeyPair()
	fmt.Printf("====================\nSubPrivateSpendKey: %x\nSubPrivateViewKey: %x\n", s.GetSpendKey(), s.GetViewKey())
	fmt.Printf("====================\nSubPublicSpendKey: %x\nSubPublicViewKey: %x\n", ps.GetSpendKey(), ps.GetViewKey())
	fmt.Println("SubAddress:", ps.GetPublicAddress(address.SubAddr))
	// create stealth address
	r, _ := address.NewRandom()
	fmt.Printf("r : %x\n", r.Bytes())
	R := ps.GetSubTranKey(r)
	fmt.Printf("R : %x\n", R.Bytes())
	sa, _ := ps.NewStealthAddress(r, 1)
	fmt.Printf("NewStealthAddress: \nP: %x\n", sa)
	// verify stealth address
	SA, _ := s.ProcessStealthAddress(k.GetViewKeyAddr(), R, 1)
	fmt.Printf("ProcessStealthAddress: \nx: %x\nP': %x\n", SA["x"], SA["P'"])
}
