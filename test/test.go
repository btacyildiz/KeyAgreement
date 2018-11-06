package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var concatenated = strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println("Time: ", concatenated)
}
