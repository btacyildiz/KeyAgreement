package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b = new(big.Int).SetString("11", 10)
	var c, d = new(big.Int).SetString("6", 10)
	var e = new(big.Int).SetUint64(2)
	fmt.Println(b, d, c)
	fmt.Println(a.Sub(a, e))
}
