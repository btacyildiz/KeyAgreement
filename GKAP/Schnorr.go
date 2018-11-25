package main

import (
	"fmt"

	"github.com/dedis/kyber"
	"github.com/dedis/kyber/group/edwards25519"
	"github.com/dedis/kyber/sign/schnorr"
)

func main() {

	suite := edwards25519.NewBlakeSHA256Ed25519()
	rand := suite.XOF([]byte("example"))
	fmt.Println("Rand: ", rand)

	dataToBeSigned := "test1234"
	out, err := schnorr.Sign(kyber.Random, rand, []byte(dataToBeSigned))
}
