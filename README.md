# naive-elliptic

This is a naive implementation of elliptic curves in affine (xy) coordinates
for curves in Short Weierstrass and Montgomery form. For the first one there are
secp256k1 and P256 defined and for the second one curve25519 is defined.
Montgomery ladder is used for scalar multiplication.

Supported operations:
* BasePointGGet
* ECPointGen
* IsOnCurveCheck
* AddECPoints
* DoubleECPoint
* ScalarMult
* ECPointToString
* StringToECPoint
* PrintECPoint

For real-world applications, one should consider https://github.com/mit-plv/fiat-crypto

## links
https://safecurves.cr.yp.to/

https://hyperelliptic.org/EFD/

https://neuromancer.sk/std/

https://en.wikipedia.org/wiki/Elliptic_curve_point_multiplication

https://web.archive.org/web/20210929025107/http://point-at-infinity.org/ecc/nisttv