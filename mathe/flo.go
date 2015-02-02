package mathe

import (
	"fmt"
	"math"
	"math/big"
)

func Me(f float64) {
	f = 435.32345
	b := big.NewInt(0).SetUint64(math.Float64bits(f))
	fmt.Println()
	for i := 0; i < 64; i++ {
		fmt.Print(b.Bit(i))
	}
	fmt.Println()
}
