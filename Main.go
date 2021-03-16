package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var die1 int = 1
	var die2 int = 1
	var result1 int = rand.Intn(die1) + 1
	var result2 int = rand.Intn(die2) + 1
	var resultT int = result1 + result2
	var col int = 0
	if result1 == 1 {
		col = col + 1
	}
	if result1 == die1 {
		col = col + 1
	}
	if result2 == 1 {
		col = col + 1
	}
	if result2 == die2 {
		col = col + 1
	}
	fmt.Println("You rolled a", resultT)
	var heatlvl int = 0
	heat := [8]int{0, 10, 12, 10, 8, 6, 4, 1}
	for h := 0; h <= heatlvl; h++ {
		var heatT int = rand.Intn(heat)
		if heatT == heat[h] {
			col = col + 1
		}
		if heatT == 1 {
			col = col + 1
			fmt.Println("and the heat generated", heatT, "for a total of", (resultT + heatT), col, "collateral was generated")
		}
	}
}
