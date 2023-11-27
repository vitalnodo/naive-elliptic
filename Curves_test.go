package naiveelliptic_test

import (
	"math/big"
	"math/rand"
	"testing"

	naiveelliptic "github.com/vitalnodo/naive-elliptic"
)

var curves = []naiveelliptic.ECPointOperations{naiveelliptic.Secp256k1(), naiveelliptic.P256(), naiveelliptic.Curve25519()}

func TestCurve(t *testing.T) {
	for i := range curves {
		curve := curves[i]
		G := curve.BasePointGGet()
		k_bytes := make([]byte, 32)
		rand.Read(k_bytes)
		d_bytes := make([]byte, 32)
		rand.Read(d_bytes)
		k := new(big.Int).SetBytes(k_bytes)
		d := new(big.Int).SetBytes(d_bytes)

		H1 := curve.ScalarMult(*d, G)
		H2 := curve.ScalarMult(*k, H1)

		H3 := curve.ScalarMult(*k, G)
		H4 := curve.ScalarMult(*d, H3)

		if H2.X.Cmp(H4.X) != 0 && H2.Y.Cmp(H4.Y) != 0 {
			t.Fatal()
		}
	}
}

func TestSerializeDeserialize(t *testing.T) {
	for i := range curves {
		curve := curves[i]
		G := curve.BasePointGGet()
		serialized := curve.ECPointToString(G)
		deserialized := curve.StringToECPoint(serialized)
		if deserialized.X.Cmp(G.X) != 0 || deserialized.Y.Cmp(G.Y) != 0 {
			t.Fatal()
		}
	}
}
