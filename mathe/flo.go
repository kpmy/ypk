package mathe

import (
	"math"
	"math/big"
)

// http://sr.wikipedia.org/wiki/%D0%94%D0%B2%D0%BE%D1%81%D1%82%D1%80%D1%83%D0%BA%D0%B0_%D1%82%D0%B0%D1%87%D0%BD%D0%BE%D1%81%D1%82#mediaviewer/File:IEEE_754_Double_Floating_Point_Format(sr).svg
func Me(f float64) (m float64, e float64) {
	switch f {
	case 0:
		return 0, 0
	case math.Inf(1), math.Inf(-1):
		return 1, math.MaxInt32
	case math.NaN():
		return 2, math.MaxInt32
	}
	b := big.NewInt(0).SetUint64(math.Float64bits(f))
	sign := 1.0
	if b.Bit(63) == 1 {
		sign = -1
	}
	mant := big.NewInt(0)
	for i := 0; i < 52; i++ {
		mant.SetBit(mant, i, b.Bit(i))
	}
	exp := big.NewInt(0)
	j := 0
	for i := 52; i < 63; i++ {
		exp.SetBit(exp, j, b.Bit(i))
		j++
	}
	mult := big.NewInt(0)
	mult.SetBit(mult, 52, 1)
	m = float64(mant.Int64())
	x := float64(mult.Int64())
	m = sign + sign*(m/x)
	e = float64(exp.Int64()) - 1023
	return
}
