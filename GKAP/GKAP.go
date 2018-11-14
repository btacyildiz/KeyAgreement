package main

import (
	"fmt"
	"hash/crc32"
	"math"
	"math/rand"
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
r = (Zq) (randomly selected)
**/
const q = 5
const p = 2*q + 1
const g = 3

const k1 = 2
const v1 = 3
const x1 = 4

var y1 = math.Mod(math.Pow(g, x1), p)

const k2 = 3
const v2 = 4
const x2 = 3

var y2 = math.Mod(math.Pow(g, x2), p)

type participant struct {
	// inited private values
	x int
	// inited public values
	k int
	v int
	r int
	y int
	// calculated public variables
	A int
	B int
	w int
	z int
	a int
	b int
	g int
}

var participants [2]participant

const participantCount = 2

func printParticipant(node participant) {
	fmt.Println("x:", node.x, " k:", node.k, " v:", node.v,
		" r:", node.r, " y:", node.y, " A:", node.A, " B:", node.B, " w:", node.w,
		" z:", node.z, " a:", node.a, " b:", node.b, " g:", node.g)
}

func initParticipant() participant {
	r := rnd(q)
	k := rnd(q)
	v := rnd(q)
	x := rnd(q)
	y := int(math.Mod(math.Pow(g, float64(x)), p))
	return participant{k: k, v: v, x: x, y: y, r: r}
}

func rnd(max int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return (r1.Int()%(max-1) + 1)
}

func getHashWithTimeStamp(input string) uint32 {
	concat := strconv.FormatInt(time.Now().Unix(), 10) + input
	message := []byte(concat)
	return crc32.ChecksumIEEE(message)
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func calcTempPubParams(node participant) participant {
	node.w = int(math.Mod(math.Pow(g, float64(node.k)), p))
	node.A = int(math.Mod(math.Pow(g, float64(node.v)), p))

	fmt.Println("W: ", node.w, " A: ", node.A, " X: ", node.x, " V: ", node.v, " Q: ", q)
	var up = math.Abs(float64(node.w - node.A*node.x))
	var partialOfB = (int(up) / node.v)
	fmt.Println("Partial of B: ", partialOfB)
	node.B = int(math.Mod(float64(partialOfB), q))
	return node
}

func calcTempSecretKeys(node participant, index int) participant {
	before := (index - 1) % participantCount
	after := (index + 1) % participantCount
	div := (float64)(participants[after].w / participants[before].w)
	powK := math.Pow(div, (float64)(node.k))
	node.z = (int)(math.Mod(powK, p))
	node.a = (int)(math.Mod(math.Pow(g, (float64)(node.r)), p))
	powR := math.Pow(div, (float64)(node.r))
	node.b = (int)(math.Mod(powR, p))
	node.g = node.r + (int)(math.Mod((float64)(node.z*node.a*node.b*node.k), q))
	return node
}

func verifyTempSecretKeys(node participant, index int) bool {
	// first check
	v1 := math.Pow(g, (float64)(node.g))
	v2 := node.a * (int)(math.Pow((float64)(node.w), (float64)(node.z*node.a*node.b*node.g)))
	// second check
	if math.Mod(v1, p) != math.Mod((float64)(v2), p) {
		return false
	}

	before := (index - 1) % participantCount
	after := (index + 1) % participantCount

	div := (float64)(participants[after].w / participants[before].w)
	v1 = math.Pow(div, (float64)(node.g))
	mul := node.z * node.a * node.b * node.g
	v2 = node.b * (int)(math.Pow((float64)(node.z), (float64)(mul)))
	if (int)(v1) != v2 {
		return false
	}
	return true
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
	node1 := initParticipant()
	printParticipant(node1)
	// printout public variables
	fmt.Println("Public Variables")
	fmt.Println("q:", q, " p:", p, " g:", g)
	/*
		// STEP 1 CALCULATE
		w, A, B := calcTempPubParams(k1, v1, x1)
		verifyRes := verifyPubVariables(w, A, B, y1)
	*/
	// STEP 2
	fmt.Println("X2: ", x2, " Y2: ", y2)

	node1 = calcTempPubParams(node1)

	verifyRes1 := verifyPubVariables(float64(node1.w), float64(node1.A),
		float64(node1.B), float64(node1.y))

	fmt.Println("Node1 verify res: ", verifyRes1)
	printParticipant(node1)
}
