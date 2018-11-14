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

/**
p and q are large prime numbers
g = (Zp)
x = (Zq)
y = g^x mod p
k,v = (Zq) (are randomly selected)
**/
const q = 5
const p = 2*q + 1
const g = 3

const k1 = 2
const v1 = 3
const x1 = 4

var y1 = math.Mod(math.Pow(g, x1), p)

const k2 = 1
const v2 = 2
const x2 = 3

var y2 = math.Mod(math.Pow(g, x2), p)

func getHashWithTimeStamp(input string) uint32 {
	concat := strconv.FormatInt(time.Now().Unix(), 10) + input
	message := []byte(concat)
	return crc32.ChecksumIEEE(message)
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func calcTempPubParams(k float64, v float64, x float64) (float64, float64, float64) {
	var w = math.Mod(math.Pow(g, k), p)

	var A = math.Mod(math.Pow(g, v), p)
	fmt.Println("W: ", w, " A: ", A, " X: ", x, " V: ", v, " Q: ", q, " ::: ", math.Mod(-4, 5))
	var up = math.Abs(w - A*x) // float64(getHashWithTimeStamp(FloatToString(w))) - A*x
	var partialOfB = (int(up) / int(v))
	fmt.Println("Partial of B: ", partialOfB)
	var B = math.Mod(float64(partialOfB), q)
	return w, A, B
}

func verifyPubVariables(w, A, B, y float64) bool {
	// check W
	if 2 > w || w > p {
		return false
	}
	fmt.Println("First check is succeeded")

	if math.Mod(math.Pow(w, q), p) != 1 {
		return false
	}

	fmt.Println("Second check is succeeded")

	//firstPart := math.Pow(g, float64(getHashWithTimeStamp(FloatToString(w))))
	firstPart := math.Mod(math.Pow(g, w), p)
	secondPart := math.Mod((math.Pow(y, A) * math.Pow(A, B)), p)
	fmt.Println("First: ", firstPart, " Second: ", secondPart)
	if firstPart != secondPart {
		return false
	}

	return true
}

func main() {
	/*
		// STEP 1 CALCULATE
		w, A, B := calcTempPubParams(k1, v1, x1)

		verifyRes := verifyPubVariables(w, A, B, y1)
	*/
	// STEP 2

	w2, A2, B2 := calcTempPubParams(k2, v2, x2)

	verifyRes2 := verifyPubVariables(w2, A2, B2, y2)

	fmt.Println("Res: ", w2, A2, B2, " Verify Res2: ", verifyRes2)
}
