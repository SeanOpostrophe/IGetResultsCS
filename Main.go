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
	var col int = collateral(roll1, roll2, dieOneSides, dieTwoSides)
	var col2, heatTotal int = calculateHeat(heatLvl)

	fmt.Println("You rolled a", roll1, "and a", roll2, "and the heat generated", heatTotal, "for a total of", (rollt + heatTotal), "and caused", col+col2, "collateral")
}
func collateral(roll1 int, roll2 int, dieOneSides int, dieTwoSides int) int {
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
	return col
}

func calculateHeat(heatLvl int) (col2 int, heatTotal int) {
	if heatLvl == 0 {
		fmt.Println("and there was no heat")
	} else {
		heatTotal = 0
		col2 = 0
		heatLvlDie := [8]int{10, 12, 10, 8, 6, 4, 1}
		for h := 1; h <= heatLvl; h++ {

			var heatDie int = heatLvlDie[h-1]
			var heatGen int = rand.Intn(heatDie) + 1
			heatTotal = heatTotal + heatGen
			fmt.Println("heat gen is", heatGen, "out of", heatDie)
			if heatGen == heatDie {
				col2 = col2 + 1
			}
			if heatGen == 1 {
				col2 = col2 + 1
			}
		}
	}
	return col2, heatTotal
}
