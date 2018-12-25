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
	p          *big.Int
	q          *big.Int
	r          *big.Int
	privateKey PkSigSchnorrPrivate
	publicKey  PkSigSchorrPublic
}

type PkSigSchnorrPrivate *big.Int

type PkSigSchorrPublic struct {
	y *big.Int
	g *big.Int
}

type PkSigSchorrSign struct {
	s *big.Int
	e *big.Int
}

func newPkSigSchnorr(p *big.Int, q *big.Int) PkSigSchnorr {
	sig := PkSigSchnorr{p: p, q: q, r: new(big.Int).SetInt64(2)}
	return sig
}

func (sig PkSigSchnorr) keyGen() (PkSigSchnorrPrivate, PkSigSchorrPublic) {
	sig.privateKey, _ = generateRandom(sig.p)
	sig.publicKey.g = new(big.Int).Mod(sig.randomGem(), sig.p)
	sig.publicKey.y = new(big.Int).Exp(sig.publicKey.g, sig.privateKey, sig.p)
	return sig.privateKey, sig.publicKey
}

func (sig PkSigSchnorr) randomGem() *big.Int {
	h, _ := generateRandom(sig.p)
	return new(big.Int).Exp(h, sig.r, sig.p)
}

func (sig PkSigSchnorr) sign(public PkSigSchorrPublic, private PkSigSchnorrPrivate, M *big.Int) PkSigSchorrSign {
	k, _ := generateRandom(sig.q)
	r := new(big.Int).Exp(g, k, sig.p)
	e := sig.hash(M, r)
	diff := k.Sub(k, new(big.Int).Mul(private, e))
	s := diff.Mod(diff, sig.q)
	return PkSigSchorrSign{e: e, s: s}
}

func (sig PkSigSchnorr) verify(public PkSigSchorrPublic, signature PkSigSchorrSign, M *big.Int) bool {
	first := new(big.Int).Exp(g, signature.s, sig.p)
	second := new(big.Int).Exp(public.y, signature.e, sig.p)
	mul := first.Mul(first, second)
	r := mul.Mod(mul, sig.p)
	eCreated := sig.hash(M, r)
	if signature.e.Cmp(eCreated) != 0 {
		fmt.Println("Created Sig: ", eCreated)
		fmt.Println("Given Sig:   ", signature.e)
		return false
	}
	return true
}

func (sig PkSigSchnorr) hash(M *big.Int, r *big.Int) *big.Int {
	PPlusQ := new(big.Int).Add(sig.p, sig.q)
	sum := new(big.Int).Add(M, r)
	return sum.Add(sum, PPlusQ)
}
