package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var stat1 int = 1
	var stat2 int = 1
	var chk1 int = rand.Intn(stat1) + 1
	var chk2 int = rand.Intn(stat2) + 1
	var chkt int = chk1 + chk2
	var col int = 0
	if chk1 == 1 {
		col = col + 1
	}
	if chk1 == stat1 {
		col = col + 1
	}
	if chk2 == 1 {
		col = col + 1
	}
	if chk2 == stat2 {
		col = col + 1
	}
	fmt.Println("You rolled a", chk1, "and a", chk2, "for a total of", chkt, "and caused", col, "collateral")

}
