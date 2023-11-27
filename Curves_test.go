package naiveelliptic_test

import (
	"fmt"
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

type TestP256Struct struct {
	scalar string
	x      string
	y      string
}

func TestP256(t *testing.T) {
	tests := []TestP256Struct{
		{"1", "6B17D1F2E12C4247F8BCE6E563A440F277037D812DEB33A0F4A13945D898C296", "4FE342E2FE1A7F9B8EE7EB4A7C0F9E162BCE33576B315ECECBB6406837BF51F5"},
		{"2", "7CF27B188D034F7E8A52380304B51AC3C08969E277F21B35A60B48FC47669978", "07775510DB8ED040293D9AC69F7430DBBA7DADE63CE982299E04B79D227873D1"},
		{"3", "5ECBE4D1A6330A44C8F7EF951D4BF165E6C6B721EFADA985FB41661BC6E7FD6C", "8734640C4998FF7E374B06CE1A64A2ECD82AB036384FB83D9A79B127A27D5032"},
		{"4", "E2534A3532D08FBBA02DDE659EE62BD0031FE2DB785596EF509302446B030852", "E0F1575A4C633CC719DFEE5FDA862D764EFC96C3F30EE0055C42C23F184ED8C6"},
		{"5", "51590B7A515140D2D784C85608668FDFEF8C82FD1F5BE52421554A0DC3D033ED", "E0C17DA8904A727D8AE1BF36BF8A79260D012F00D4D80888D1D0BB44FDA16DA4"},
		{"112233445566778899", "339150844EC15234807FE862A86BE77977DBFB3AE3D96F4C22795513AEAAB82F", "B1C14DDFDC8EC1B2583F51E85A5EB3A155840F2034730E9B5ADA38B674336A21"},
		{"115792089210356248762697446949407573529996955224135760342422259061068512044367", "7CF27B188D034F7E8A52380304B51AC3C08969E277F21B35A60B48FC47669978", "F888AAEE24712FC0D6C26539608BCF244582521AC3167DD661FB4862DD878C2E"},
		{"115792089210356248762697446949407573529996955224135760342422259061068512044368", "6B17D1F2E12C4247F8BCE6E563A440F277037D812DEB33A0F4A13945D898C296", "B01CBD1C01E58065711814B583F061E9D431CCA994CEA1313449BF97C840AE0A"},
	}
	for i := range tests {
		scalar, _ := new(big.Int).SetString(tests[i].scalar, 10)
		x, _ := new(big.Int).SetString(tests[i].x, 16)
		y, _ := new(big.Int).SetString(tests[i].y, 16)
		curve := naiveelliptic.P256()
		actual_res := curve.ScalarMult(*scalar, curve.BasePointGGet())
		if actual_res.X.Cmp(x) != 0 || actual_res.Y.Cmp(y) != 0 {
			fmt.Println(i)
			t.Fatal()
		}

	}

}
