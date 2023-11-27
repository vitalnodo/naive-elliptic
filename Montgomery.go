package naiveelliptic

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
)

// by^2 = x^3 + a*x^2 + b
type MontgomeryCurve struct {
	Curve
}

// G-generator receiving
func (curve MontgomeryCurve) BasePointGGet() (point ECPoint) {
	return *curve.G
}

// ECPoint creation
func (curve MontgomeryCurve) ECPointGen(x, y *big.Int) (point ECPoint) {
	point = ECPoint{x, y}
	if !curve.IsOnCurveCheck(point) {
		log.Panic("ECPointGen")
	}
	return point
}

// DOES P ∈ CURVE?
func (curve MontgomeryCurve) IsOnCurveCheck(a ECPoint) (c bool) {
	lhs := new(big.Int).Mul(a.Y, a.Y)
	lhs = lhs.Mul(curve.b, lhs)
	xxx := new(big.Int).Mul(a.X, a.X)
	xxx = xxx.Mul(xxx, a.X)
	ax := new(big.Int).Mul(curve.a, a.X)
	ax = ax.Mul(big.NewInt(2), ax)
	rhs := new(big.Int).Add(xxx, ax)
	rhs = rhs.Add(rhs, curve.b)
	return lhs.Cmp(rhs) == 0
}

// P + Q
func (curve MontgomeryCurve) AddECPoints(a, b ECPoint) (c ECPoint) {
	zero := big.NewInt(0)
	if a.X.Cmp(zero) == 0 && a.Y.Cmp(zero) == 0 {
		return b
	}
	if b.X.Cmp(zero) == 0 && b.Y.Cmp(zero) == 0 {
		return a
	}
	// x3 = b*(y2-y1)^2/(x2-x1)^2-a-x1-x2
	// y3 = (2*x1+x2+a)*(y2-y1)/(x2-x1)-b*(y2-y1)^3/(x2-x1)^3-y1
	num := new(big.Int).Set(b.Y)
	num.Sub(num, a.Y)
	num.Mul(num, num)
	num.Mul(num, curve.b)

	den := new(big.Int).Set(b.X)
	den.Sub(den, a.X)
	den.Mul(den, den)
	den.ModInverse(den, curve.p)

	c.X = new(big.Int).Set(num)
	c.X.Mul(c.X, den)
	c.X.Sub(c.X, curve.a)
	c.X.Sub(c.X, a.X)
	c.X.Sub(c.X, b.X)
	c.X.Mod(c.X, curve.p)

	c.Y = big.NewInt(2)
	c.Y.Mul(c.Y, a.X)
	c.Y.Add(c.Y, b.X)
	c.Y.Add(c.Y, curve.a)

	num.Sub(b.Y, a.Y)
	den.Sub(b.X, a.X)
	den.ModInverse(den, curve.p)
	c.Y.Mul(c.Y, num)
	c.Y.Mul(c.Y, den)

	num.Sub(b.Y, a.Y)
	tmp1 := new(big.Int).Mul(num, num)
	tmp1.Mul(tmp1, num)
	num.Mul(tmp1, curve.b)

	den.Sub(b.X, a.X)
	tmp1 = new(big.Int).Mul(den, den)
	den.Mul(tmp1, den)
	den.ModInverse(den, curve.p)

	tmp1.Mul(num, den)
	c.Y.Sub(c.Y, tmp1)
	c.Y.Sub(c.Y, a.Y)
	c.Y.Mod(c.Y, curve.p)
	return c
}

// 2P
func (curve MontgomeryCurve) DoubleECPoints(a ECPoint) (c ECPoint) {
	zero := big.NewInt(0)
	if a.X.Cmp(zero) == 0 && a.Y.Cmp(zero) == 0 {
		return a
	}
	// x3 = b*(3*x1^2+2*a*x1+1)^2/(2*b*y1)^2-a-x1-x1
	// y3 = (2*x1+x1+a)*(3*x1^2+2*a*x1+1)/(2*b*y1)-b*(3*x1^2+2*a*x1+1)^3/(2*b*y1)^3-y1
	num := big.NewInt(3)
	num.Mul(num, a.X)
	num.Mul(num, a.X)
	tmp1 := big.NewInt(2)
	tmp1.Mul(tmp1, curve.a)
	tmp1.Mul(tmp1, a.X)
	num.Add(num, tmp1)
	num.Add(num, big.NewInt(1))
	num.Mul(num, num)
	num.Mul(num, curve.b)
	den := big.NewInt(2)
	den.Mul(den, curve.b)
	den.Mul(den, a.Y)
	den.Mul(den, den)
	den.ModInverse(den, curve.p)
	c.X = new(big.Int).Mul(num, den)
	c.X.Sub(c.X, curve.a)
	c.X.Sub(c.X, a.X)
	c.X.Sub(c.X, a.X)
	c.X.Mod(c.X, curve.p)

	c.Y = big.NewInt(2)
	c.Y.Mul(c.Y, a.X)
	c.Y.Add(c.Y, a.X)
	c.Y.Add(c.Y, curve.a)

	num = big.NewInt(3)
	num.Mul(num, a.X)
	num.Mul(num, a.X)

	tmp1 = big.NewInt(2)
	tmp1.Mul(tmp1, curve.a)
	tmp1.Mul(tmp1, a.X)
	num.Add(num, tmp1)
	num.Add(num, big.NewInt(1))

	den = big.NewInt(2)
	den.Mul(den, curve.b)
	den.Mul(den, a.Y)
	den.ModInverse(den, curve.p)
	c.Y.Mul(c.Y, num)
	c.Y.Mul(c.Y, den)

	num = big.NewInt(3)
	num.Mul(num, a.X)
	num.Mul(num, a.X)

	tmp1 = big.NewInt(2)
	tmp1.Mul(tmp1, curve.a)
	tmp1.Mul(tmp1, a.X)

	num.Add(num, tmp1)
	num.Add(num, big.NewInt(1))

	tmp2 := new(big.Int).Mul(num, num)
	num.Mul(tmp2, num)
	num.Mul(num, curve.b)

	den = big.NewInt(2)
	den.Mul(den, curve.b)
	den.Mul(den, a.Y)

	tmp2 = new(big.Int).Mul(den, den)
	den.Mul(tmp2, den)
	den.ModInverse(den, curve.p)
	tmp1.Mul(num, den)
	c.Y.Sub(c.Y, tmp1)
	c.Y.Sub(c.Y, a.Y)
	c.Y.Mod(c.Y, curve.p)
	return c
}

// k * P
func (curve MontgomeryCurve) ScalarMult(k big.Int, a ECPoint) (c ECPoint) {
	R0 := ECPoint{big.NewInt(0), big.NewInt(0)}
	R1 := a
	bit := k.BitLen() - 1
	for bit >= 0 {
		if k.Bit(bit) == 0 {
			R1 = curve.AddECPoints(R0, R1)
			R0 = curve.DoubleECPoints(R0)
		} else {
			R0 = curve.AddECPoints(R0, R1)
			R1 = curve.DoubleECPoints(R1)
		}
		bit -= 1
	}
	R0.X.Mod(R0.X, curve.p)
	R0.Y.Mod(R0.Y, curve.p)
	return R0
}

// Serialize point
func (curve MontgomeryCurve) ECPointToString(point ECPoint) (s string) {
	result := make([]byte, 1+32+32)
	result[0] = 0x04
	x := result[1:33]
	y := result[33:]
	point.X.FillBytes(x)
	point.Y.FillBytes(y)
	return fmt.Sprintf("%x", result)
}

// Deserialize point
func (curve MontgomeryCurve) StringToECPoint(s string) (point ECPoint) {
	data, err := hex.DecodeString(s)
	if err != nil {
		log.Panic(err)
	}
	if data[0] != 0x04 || len(data) != 65 {
		log.Panic("StringToECPoint")
	}
	point.X = new(big.Int).SetBytes(data[1:33])
	point.Y = new(big.Int).SetBytes(data[33:])
	return point
}

// PrintECPoint
func (curve MontgomeryCurve) PrintECPoint(point ECPoint) {
	fmt.Printf("(%d,%d)\n", point.X, point.Y)
}
