package naiveelliptic

import (
	"math/big"
)

type Curve struct {
	Name string
	a    *big.Int
	b    *big.Int
	p    *big.Int
	n    *big.Int
	G    *ECPoint
}
