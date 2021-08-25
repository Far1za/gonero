package base58

import b58 "github.com/mr-tron/base58"

// decode monero address
func Decode(addr string) []byte {
	var a []byte
	for x := 0; x < len(addr); x += 11 {
		b, _ := b58.Decode(addr[x:min(x+11, len(addr))])
		a = append(a, b...)
	}
	return a
}

// encode monero address
func Encode(data ...[]byte) string {
	var a string
	var d []byte
	for _, i := range data {
		d = append(d, i...)
	}
	s := len(d)
	for x := 0; x < s; x += 8 {
		e := b58.Encode(d[x:min(x+8, len(d))])
		if x+8 < s && len(e) < 11 {
			e = "1" + e
		}
		a += e
	}
	return a
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
