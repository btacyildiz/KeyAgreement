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

type PkSigSchorrSign struct {
	s *big.Int
	e *big.Int
}

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

func (sig PkSigSchnorr) sign(private *big.Int, g *big.Int, M *big.Int) *PkSigSchorrSign {
	k, _ := generateRandom(sig.q)
	r := new(big.Int).Exp(g, k, sig.p)
	e := sig.hash(M, r)
	diff := k.Sub(k, new(big.Int).Mul(private, e))
	s := diff.Mod(diff, sig.q)
	return &PkSigSchorrSign{e: e, s: s}
}

func (sig PkSigSchnorr) verify(public *big.Int, g *big.Int, sign *PkSigSchorrSign, M *big.Int) bool {
	first := new(big.Int).Exp(g, sign.s, sig.p)
	second := new(big.Int).Exp(public, sign.e, sig.p)
	mul := first.Mul(first, second)
	r := mul.Mod(mul, sig.p)
	eCreated := sig.hash(M, r)
	if sign.e.Cmp(eCreated) != 0 {
		fmt.Println("Created Sig: ", eCreated)
		fmt.Println("Given Sig:   ", sign.e)
		return false
	}
	return true
}

func (sig PkSigSchnorr) hash(M *big.Int, r *big.Int) *big.Int {
	PPlusQ := new(big.Int).Add(sig.p, sig.q)
	sum := new(big.Int).Add(M, r)
	return sum.Add(sum, PPlusQ)
}

func test() {
	sig := newPkSigSchnorr(p, q)
	privateKey, publicKey, g := sig.keyGen()
	sigOmega, _ := new(big.Int).SetString("123213123213", 10)

	signature := sig.sign(privateKey, g, sigOmega)
	res := sig.verify(publicKey, g, signature, sigOmega)
	fmt.Println("Res: ", res)
}
