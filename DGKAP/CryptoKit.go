package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

func getGenerator(p *big.Int) (*big.Int, error) {
	rnd, err := generateRandom(p)
	if err == nil {
		rnd.Mul(rnd, rnd)
		fmt.Println("Mull ", rnd)
		final := rnd.Mod(rnd, p)
		return final, nil
	}
	return nil, err
}

func generateRandom(p *big.Int) (*big.Int, error) {
	if p.Cmp(new(big.Int).SetInt64(3)) != 1 {
		return nil, errors.New("p cannot be smaller than 4")
	}
	pVal := new(big.Int).Set(p)
	subVal := new(big.Int).SetInt64(3)
	randInt, err := rand.Int(rand.Reader, pVal.Sub(pVal, subVal))
	if err == nil {
		n := randInt.Add(randInt, new(big.Int).SetInt64(2))
		return n, nil
	}
	return nil, err
}

func main() {
	var p = new(big.Int).SetInt64(123123)
	a, b := generateRandom(p)
	if b == nil {
		fmt.Println("Rnd: ", a)
		fmt.Println("P: ", p)
		g, err := getGenerator(p)
		fmt.Println("Generator: ", g, " err: ", err)
	}
}
