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
	g *big.Int
	x *big.Int
	y *big.Int
}

func New(p *big.Int, q *big.Int) PkSigSchnorr {
	sig := PkSigSchnorr{p: p, q: q, r: new(big.Int).SetInt64(2)}
	return sig
}

func (sig PkSigSchnorr) keyGen() (*big.Int, *big.Int, *big.Int) {
	// WARNING mod operation is performed check it out!
	sig.x, _ = generateRandom(sig.p)
	sig.g = new(big.Int).Mod(sig.randomGem(), sig.p)
	sig.y = new(big.Int).Exp(sig.g, sig.x, sig.p)
	return sig.x, sig.y, sig.g
}

func (sig PkSigSchnorr) randomGem() *big.Int {
	h, _ := generateRandom(sig.p)
	sig.g = new(big.Int).Exp(h, sig.r, sig.p)
	return sig.g
}

func (sig PkSigSchnorr) sign(y *big.Int, g *big.Int, x *big.Int, M *big.Int) (*big.Int, *big.Int) {
	k, _ := generateRandom(sig.q)
	r := new(big.Int).Exp(g, k, sig.p)
	e := sig.hash(M, r)
	diff := k.Sub(k, new(big.Int).Mul(x, e))
	s := diff.Mod(diff, sig.q)
	return e, s
}

func (sig PkSigSchnorr) verify(y *big.Int, g *big.Int, s *big.Int, e *big.Int, M *big.Int) bool {
	first := new(big.Int).Exp(g, s, sig.p)
	second := new(big.Int).Exp(y, e, p)
	mul := first.Mul(first, second)
	r := mul.Mod(mul, p)
	eCreated := sig.hash(M, r)
	if e.Cmp(eCreated) != 0 {
		fmt.Println("Created Sig: ", eCreated)
		fmt.Println("Given Sig: ", e)
		return false
	}
	return true
}

func (sig PkSigSchnorr) hash(M *big.Int, r *big.Int) *big.Int {
	PPlusQ := new(big.Int).Add(sig.p, sig.q)
	sum := new(big.Int).Add(M, r)
	return sum.Add(sum, PPlusQ)
}

func main() {
	PkSigSchnorr
}
