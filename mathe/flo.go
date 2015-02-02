package mathe

import (
	"fmt"
	"math"
	"math/big"
)

// http://sr.wikipedia.org/wiki/%D0%94%D0%B2%D0%BE%D1%81%D1%82%D1%80%D1%83%D0%BA%D0%B0_%D1%82%D0%B0%D1%87%D0%BD%D0%BE%D1%81%D1%82#mediaviewer/File:IEEE_754_Double_Floating_Point_Format(sr).svg
func Me(f float64) {
	f = 435.32345
	b := big.NewInt(0).SetUint64(math.Float64bits(f))
	for i := 63; i >= 0; i-- {
		fmt.Print(b.Bit(i))
		switch i {
		case 63, 52:
			fmt.Print(" ")
		}
	}
	fmt.Println()
	sign := 1.0
	if b.Bit(63) == 1 {
		sign = -1
	}
	fmt.Println(sign)
	mant := big.NewInt(0)
	for i := 0; i < 52; i++ {
		fmt.Print(b.Bit(i))
		mant.SetBit(mant, i, b.Bit(i))
	}
	fmt.Println()
	fmt.Println(mant)
	exp := big.NewInt(0)
	j := 0
	for i := 52; i < 63; i++ {
		fmt.Print(b.Bit(i))
		exp.SetBit(exp, j, b.Bit(i))
		j++
	}
	fmt.Println()
	fmt.Println(exp)
	fmt.Println(mant)
	mult := big.NewInt(0)
	mult.SetBit(mult, 52, 1)
	fmt.Println(mant, mult)
	fmt.Println()
	m := float64(mant.Int64())
	x := float64(mult.Int64())
	fmt.Println(sign + sign*(m/x))
	fmt.Println(float64(exp.Int64()) - 1023)

}
