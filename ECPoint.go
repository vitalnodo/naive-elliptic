package naiveelliptic

import (
	"math/big"
)

type ECPoint struct {
	X *big.Int
	Y *big.Int
}

type ECPointOperations interface {
	BasePointGGet() (point ECPoint)
	ECPointGen(x, y *big.Int) (point ECPoint)
	IsOnCurveCheck(a ECPoint) (c bool)
	AddECPoints(a, b ECPoint) (c ECPoint)
	DoubleECPoints(a ECPoint) (c ECPoint)
	ScalarMult(k big.Int, a ECPoint) (c ECPoint)
	ECPointToString(point ECPoint) (s string)
	StringToECPoint(s string) (point ECPoint)
	PrintECPoint(point ECPoint)
}
