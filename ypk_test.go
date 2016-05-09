package ypk

import (
	"fmt"
	"math"
	"testing"

	"github.com/kpmy/ypk/mathe"
)

func TestMathe(*testing.T) {
	fmt.Println(mathe.Me(0))
	fmt.Println(mathe.Me(math.Inf(1)))
	fmt.Println(mathe.Me(math.Inf(-1)))
	fmt.Println(mathe.Me(math.NaN()))
	fmt.Println(mathe.Me(548))
}
