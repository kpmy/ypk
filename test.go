package main

import (
	"fmt"
	"math"
	"ypk/mathe"
)

func main() {
	fmt.Println(mathe.Me(0))
	fmt.Println(mathe.Me(math.Inf(1)))
	fmt.Println(mathe.Me(math.Inf(-1)))
	fmt.Println(mathe.Me(math.NaN()))
	fmt.Println(mathe.Me(548))
}
