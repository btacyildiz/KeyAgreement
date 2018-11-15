package main

import (
	"fmt"
	"math"
)

func main() {
	/*s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)*/
	fmt.Println("RandNum: ", math.Mod(-1, 2))
}
