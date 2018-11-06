package main

import (
	"fmt"
	"hash/crc32"
	"math"
	"strconv"
	"time"
)

// participants perform key agreement protocol
// initialization parameters are defined before

const p = 11
const q = 5
const g = 3

const k1 = 2
const v1 = 3
const x1 = 4
const y1 = 6

const k2 = 5
const v2 = 6
const x2 = 7

func getHashWithTimeStamp(input string) uint32 {
	concat := strconv.FormatInt(time.Now().Unix(), 10) + input
	message := []byte(concat)
	return crc32.ChecksumIEEE(message)
}

func calcTempPubParams(k float64, v float64, x float64) (float64, float64, float64) {
	var w = math.Mod(math.Pow(g, k), p)

	var A = math.Mod(math.Pow(g, v), p)
	// TODO did not use timestamp and hash insecure against replay attack
	fmt.Println("W: ", w, " A: ", A, " X: ", x, " V: ", v, " Q: ", q, " ::: ", math.Mod(-4, 5))
	var partialOfB = (int(getHashWithTimeStamp(strconv.FormatFloat(w))-A*x) / int(v))
	fmt.Println("Partial of B: ", partialOfB)
	var B = math.Mod(float64(partialOfB), q)
	return w, A, B
}

func verifyPubVariables() {

}

func main() {
	// STEP 1 CALCULATE
	a, b, c := calcTempPubParams(k1, v1, x1)

	fmt.Println("Res: ", a, b, c)
}
