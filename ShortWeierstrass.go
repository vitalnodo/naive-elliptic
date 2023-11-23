package naiveelliptic

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
)

// y^2 = x^3 + a*x + b
type ShortWeierstrassCurve struct {
	Curve
}

// G-generator receiving
func (curve ShortWeierstrassCurve) BasePointGGet() (point ECPoint) {
	return *curve.G
}

// ECPoint creation
func (curve ShortWeierstrassCurve) ECPointGen(x, y *big.Int) (point ECPoint) {
	point = ECPoint{x, y}
	if !curve.IsOnCurveCheck(point) {
		log.Panic("ECPointGen")
	}
	return point
}

// DOES P ∈ CURVE?
func (curve ShortWeierstrassCurve) IsOnCurveCheck(a ECPoint) (c bool) {
	lhs := new(big.Int).Mul(a.Y, a.Y)
	xxx := new(big.Int).Mul(a.X, a.X)
	xxx = xxx.Mul(xxx, a.X)
	ax := new(big.Int).Mul(curve.a, a.X)
	rhs := new(big.Int).Add(xxx, ax)
	rhs = rhs.Add(rhs, curve.b)
	return lhs.Cmp(rhs) == 0
}

// P + Q
func (curve ShortWeierstrassCurve) AddECPoints(a, b ECPoint) (c ECPoint) {
	zero := big.NewInt(0)
	if a.X.Cmp(zero) == 0 && a.Y.Cmp(zero) == 0 {
		return b
	}
	if b.X.Cmp(zero) == 0 && b.Y.Cmp(zero) == 0 {
		return a
	}
	m_num := new(big.Int).Sub(b.Y, a.Y)
	m_den := new(big.Int).Sub(b.X, a.X)
	m_den = m_den.ModInverse(m_den, curve.p)
	m := new(big.Int).Mul(m_num, m_den)
	m = m.Mod(m, curve.p)

	c.X = new(big.Int).Mul(m, m)
	c.X = c.X.Sub(c.X, a.X)
	c.X = c.X.Sub(c.X, b.X)
	c.X = c.X.Mod(c.X, curve.p)

	c.Y = new(big.Int).Sub(a.X, c.X)
	c.Y = c.Y.Mul(m, c.Y)
	c.Y = c.Y.Sub(c.Y, a.Y)
	c.Y = c.Y.Mod(c.Y, curve.p)
	return c
}

// 2P
func (curve ShortWeierstrassCurve) DoubleECPoints(a ECPoint) (c ECPoint) {
	zero := big.NewInt(0)
	if a.X.Cmp(zero) == 0 && a.Y.Cmp(zero) == 0 {
		return a
	}
	lambda := new(big.Int).Mul(a.X, a.X)
	lambda = lambda.Mul(big.NewInt(3), lambda)
	lambda = lambda.Add(lambda, curve.a)
	lambda_den := new(big.Int).Mul(big.NewInt(2), a.Y)
	lambda_den = lambda_den.ModInverse(lambda_den, curve.p)
	lambda = lambda.Mul(lambda, lambda_den)
	c.X = new(big.Int).Mul(lambda, lambda)
	c.X = c.X.Sub(c.X, new(big.Int).Mul(big.NewInt(2), a.X))
	c.X = c.X.Mod(c.X, curve.p)

	c.Y = new(big.Int).Sub(a.X, c.X)
	c.Y = c.Y.Mul(lambda, c.Y)
	c.Y = c.Y.Sub(c.Y, a.Y)
	c.Y = c.Y.Mod(c.Y, curve.p)
	return c
}

// k * P
func (curve ShortWeierstrassCurve) ScalarMult(k big.Int, a ECPoint) (c ECPoint) {
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
func (curve ShortWeierstrassCurve) ECPointToString(point ECPoint) (s string) {
	result := make([]byte, 1+32+32)
	result[0] = 0x04
	x := result[1:33]
	y := result[33:]
	point.X.FillBytes(x)
	point.Y.FillBytes(y)
	return fmt.Sprintf("%x", result)
}

// Deserialize point
func (curve ShortWeierstrassCurve) StringToECPoint(s string) (point ECPoint) {
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
func (curve ShortWeierstrassCurve) PrintECPoint(point ECPoint) {
	fmt.Printf("(%d,%d)\n", point.X, point.Y)
}
