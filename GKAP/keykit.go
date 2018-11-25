package keykit

import (
	"fmt"
	"math/big"
	"crypto/rand"
)

func getGenerator(p big.Int) big.Int {
	rnd, err := generateRandom(p) 
	if err != nil{
		rnd.Mul(rnd, rnd).Div
		fmt.Println("Mull ", rnd)
		final := rnd.Mod(final, p)
		return final
	}
	return err
}

func generateRandom(p big.Int) big.Int, error {
	n, err := rand.Int(rand.Reader, (p - 2) ) -1 
	if err != nill{
		return n, nil
	}
	return nil, err
}
