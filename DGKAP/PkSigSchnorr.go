package main

import (
	"fmt"
	"math/big"
)

/*
PkSigSchnorr
Partial implementation of pksig_schnorr91
package in charm crypto
https://jhuisi.github.io/charm/_modules/pksig_schnorr91.html#SchnorrSig
*/
type PkSigSchnorr struct {
	p *big.Int
	q *big.Int
	r *big.Int
}

//type PkSigSchnorrPrivate *big.Int

/*type PkSigSchorrPublic struct {
	y *big.Int
	g *big.Int
}*/
/*
type PkSigSchorrSign struct {
	s *big.Int
	e *big.Int
}
*/
func newPkSigSchnorr(p *big.Int, q *big.Int) PkSigSchnorr {
	sig := PkSigSchnorr{p: p, q: q, r: new(big.Int).SetInt64(2)}
	sig.print()
	return sig
}

func (sig PkSigSchnorr) print() {
	fmt.Println("p: ", sig.p)
	fmt.Println("q: ", sig.q)
	fmt.Println("r: ", sig.r)
}

func (sig PkSigSchnorr) keyGen() (*big.Int, *big.Int, *big.Int) {
	privateKey, _ := generateRandom(sig.p)
	g := new(big.Int).Mod(sig.randomGem(), sig.p)
	fmt.Println("Berfore Sig.g: ", g)
	publicKey := new(big.Int).Exp(g, privateKey, p)
	sig.print()
	return privateKey, publicKey, g
}

func (sig PkSigSchnorr) randomGem() *big.Int {
	h, _ := generateRandom(sig.p)
	return new(big.Int).Exp(h, sig.r, sig.p)
}

func (sig PkSigSchnorr) sign(private *big.Int, g *big.Int, M *big.Int) (*big.Int, *big.Int) {
	k, _ := generateRandom(sig.q)
	fmt.Println("Sig g: Hello")
	fmt.Println("Sig g: ", g)
	r := new(big.Int).Exp(g, k, sig.p)
	e := sig.hash(M, r)
	diff := k.Sub(k, new(big.Int).Mul(private, e))
	s := diff.Mod(diff, sig.q)
	return e, s
}

func (sig PkSigSchnorr) verify(public *big.Int, g *big.Int, s *big.Int, e *big.Int, M *big.Int) bool {
	first := new(big.Int).Exp(g, s, sig.p)
	second := new(big.Int).Exp(public, e, sig.p)
	mul := first.Mul(first, second)
	r := mul.Mod(mul, sig.p)
	eCreated := sig.hash(M, r)
	if e.Cmp(eCreated) != 0 {
		fmt.Println("Created Sig: ", eCreated)
		fmt.Println("Given Sig:   ", e)
		return false
	}
	return true
}

func (sig PkSigSchnorr) hash(M *big.Int, r *big.Int) *big.Int {
	PPlusQ := new(big.Int).Add(sig.p, sig.q)
	sum := new(big.Int).Add(M, r)
	return sum.Add(sum, PPlusQ)
}
