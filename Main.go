package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//  a pool of two dice comprises a stat
//  collateral increases when minimum or maximum rolled
func main() {
	argLength := len(os.Args[1:])
	fmt.Printf("Arg length is %d\n", argLength)

	for i, a := range os.Args[1:] {
		fmt.Printf("Arg %d is %s\n", i+1, a)
	}
	programName := os.Args[0]
	dieOne := os.Args[1]
	dieTwo := os.Args[2]
	var heatInput = os.Args[3]
	fmt.Println("You entered a", dieOne, "and a", dieTwo, " and set the heat level to", heatInput, " for ", programName)
	die1, _ := strconv.Atoi(dieOne)
	die2, _ := strconv.Atoi(dieTwo)
	heatLvl, _ := strconv.Atoi(heatInput)

	rand.Seed(time.Now().Unix())

	var dieOneSides int = die1
	var dieTwoSides int = die2
	var roll1 int = rand.Intn(dieOneSides) + 1
	var roll2 int = rand.Intn(dieTwoSides) + 1
	var rollt int = roll1 + roll2
	var col int = 0
	if roll1 == 1 {
		col = col + 1

	}
	if roll1 == dieOneSides {
		col = col + 1
	}
	if roll2 == 1 {
		col = col + 1
	}
	if roll2 == dieTwoSides {
		col = col + 1
	}
	var heatT int = 0
	heat := 8 //[8]int{0, 10, 12, 10, 8, 6, 4, 1}
	for h := 0; h <= heatLvl; h++ {
		var heatGen int = rand.Intn(heat)
		heatT = heatT + heatGen
		if heatGen == heat /*[h]*/ {
			col = col + 1
		}
		if heatGen == 1 {
			col = col + 1
		}
	}
	fmt.Println("You rolled a", roll1, "and a", roll2, "and the heat generated", heatT, "for a total of", (rollt + heatT), "and caused", col, "collateral")
}
