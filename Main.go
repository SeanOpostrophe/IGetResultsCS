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
	dieTwo := os.Args[1:][1]
	fmt.Println("You entered a", dieOne, "and a", dieTwo, "for ", programName)
	die1, _ := strconv.Atoi(dieOne)
	die2, _ := strconv.Atoi(dieTwo)

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
	fmt.Println("You rolled a", roll1, "and a", roll2, "for a total of", rollt, "and caused", col, "collateral")

}
