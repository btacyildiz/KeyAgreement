package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
	var p,q,a = 7,3,2
	
	var res = int(math.Pow(float64(a),float64(q)))

	fmt.Printf("p=%d q=%d a=%d\n",p,q,a)
	fmt.Printf("%d\n",int(res))
	
	if res % p == 1{
		fmt.Printf("YES\n")
	}else{
		fmt.Printf("No\n")
	}*/

	var x = "asd"
	var y = "bcd"
	var t int = 5

	fmt.Println("%s", x+y+strconv.Itoa(t))

}