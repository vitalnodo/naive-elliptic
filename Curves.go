package naiveelliptic

import (
	"log"
	"math/big"
)

func Secp256k1() ShortWeierstrassCurve {
	c := ShortWeierstrassCurve{}
	c.Curve.Name = "secp256k1"
	c.Curve.a = big.NewInt(0)
	c.Curve.b = big.NewInt(7)
	p, ok := new(big.Int).SetString("fffffffffffffffffffffffffffffffffffffffffffffffffffffffefffffc2f", 16)
	if !ok {
		log.Panic("secp256k1")
	}
	c.Curve.p = p
	n, ok := new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	if !ok {
		log.Panic("secp256k1")
	}
	c.Curve.n = n
	G_x, ok := new(big.Int).SetString("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798", 16)
	if !ok {
		log.Panic("secp256k1")
	}
	G_y, ok := new(big.Int).SetString("483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8", 16)
	if !ok {
		log.Panic("secp256k1")
	}
	c.G = &ECPoint{G_x, G_y}
	return c
}

func P256() ShortWeierstrassCurve {
	c := ShortWeierstrassCurve{}
	c.Curve.Name = "P256"
	a, ok := new(big.Int).SetString("ffffffff00000001000000000000000000000000fffffffffffffffffffffffc", 16)
	if !ok {
		log.Panic("P256")
	}
	c.Curve.a = a
	b, ok := new(big.Int).SetString("5ac635d8aa3a93e7b3ebbd55769886bc651d06b0cc53b0f63bce3c3e27d2604b", 16)
	if !ok {
		log.Panic("P256")
	}
	c.Curve.b = b
	p, ok := new(big.Int).SetString("ffffffff00000001000000000000000000000000ffffffffffffffffffffffff", 16)
	if !ok {
		log.Panic("P256")
	}
	c.Curve.p = p
	G_x, ok := new(big.Int).SetString("6b17d1f2e12c4247f8bce6e563a440f277037d812deb33a0f4a13945d898c296", 16)
	if !ok {
		log.Panic("P256")
	}
	G_y, ok := new(big.Int).SetString("4fe342e2fe1a7f9b8ee7eb4a7c0f9e162bce33576b315ececbb6406837bf51f5", 16)
	if !ok {
		log.Panic("P256")
	}
	c.G = &ECPoint{G_x, G_y}
	n, ok := new(big.Int).SetString("ffffffff00000000ffffffffffffffffbce6faada7179e84f3b9cac2fc632551", 16)
	if !ok {
		log.Panic("P256")
	}
	c.Curve.n = n
	return c
}
