package main

import (
	"fmt"
	"math/big"
)

var participants [3]*big.Int
var p, _ = new(big.Int).SetString("156816585111264668689583680968857341596876961491501655859473581156994765485015490912709775771877391134974110808285244016265856659644360836326566918061490651852930016078015163968109160397122004869749553669499102243382571334855815358562585736488447912605222780091120196023676916968821094827532746274593222577067", 10)
var q, _ = new(big.Int).SetString("178408292555632334344791840484428670798438480745750827929736790578497382742507745456354887885938695567487055404142622008132928329822180418163283459030745325926465008039007581984054580198561002434874776834749551121691285667427907679281292868244223956302611390045560098011838458484410547413766373137296611288533", 10)

func calculateKey(node *big.Int, index int) *big.Int {
	resulting_key := new(big.Int).Mod(node, p)
	for i := 0; i < len(participants); i++ {
		if i != index {
			fmt.Println("Calculating key in of ", i)
			mulled := new(big.Int).Mul(resulting_key, participants[i])
			resulting_key = mulled.Mod(mulled, p)
		}
	}
	return resulting_key
}
func main() {
	/*
		var ck1, _ = new(big.Int).SetString("5", 10)
		var ck2, _ = new(big.Int).SetString("5", 10)
		var ck3, _ = new(big.Int).SetString("16", 10)

		fmt.Println(new(big.Int).Sub(ck1, ck3))
		fmt.Println(new(big.Int).Mod(new(big.Int).Sub(ck1, ck3), ck2))
	*/
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(new(big.Int).SetBytes(b))

	/*participants[0] = ck1
	participants[1] = ck2
	participants[2] = ck3

	for i := 0; i < len(participants); i++ {
		fmt.Println("Node ", i, " Key: ", calculateKey(participants[i], i))
	}

	/*var e = new(big.Int).SetUint64(2)
	fmt.Println(b, d, c)
	fmt.Println(a.Sub(a, e))*/
}
