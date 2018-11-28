package main

import "math/big"

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

}

func (sig PkSigSchnorr) hash(M big.Int, r big.Int) big.Int {

}

func (sig PkSigSchnorr) verify() {

}
