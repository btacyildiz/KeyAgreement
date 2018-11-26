package main

import (
	"encoding/hex"
	"fmt"

	"github.com/dedis/kyber/group/edwards25519"
	"github.com/dedis/kyber/sign/schnorr"
	"github.com/dedis/kyber/util/key"
)

func main() {
	suite := edwards25519.NewBlakeSHA256Ed25519()
	rand := suite.XOF([]byte("example"))
	fmt.Println("Rand: ", rand.Read)

	kp := key.NewKeyPair(suite)

	dataToBeSigned := "test1234"
	out, err := schnorr.Sign(suite, kp.Private, []byte(dataToBeSigned))
	if err != nil {
		fmt.Println("Sign failed")
		return
	}
	fmt.Println("Signature: ", hex.Dump(out))

	err = schnorr.Verify(suite, kp.Public, []byte(dataToBeSigned), out)
	if err != nil {
		fmt.Println("Verify failed")
		return
	}
	fmt.Println("Verify succeeded")
	return
}
