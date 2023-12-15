package naiveelliptic

import (
	"math/big"
)

type Curve struct {
	Name string
	A    *big.Int
	B    *big.Int
	P    *big.Int
	N    *big.Int
	G    *ECPoint
}
