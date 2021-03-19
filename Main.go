package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	argLength := len(os.Args[1:])
	// use fmt.Printf() to format string
	fmt.Printf("Arg length is %d", argLength) 
	for i, a := range os.Args[1:] {
        fmt.Printf("Arg %d is %s\n", i+1, a) 
	}
	
	rand.Seed(time.Now().Unix())
	var die1 int = 6
	var die2 int = 12
	var heatlvl int = 1
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

	var heatT int = 0
	heat := 8 //[8]int{0, 10, 12, 10, 8, 6, 4, 1}
	for h := 0; h <= heatlvl; h++ {
		var heatGen int = rand.Intn(heat)
		heatT = heatT + heatGen
		if heatGen == heat /*[h]*/ {
			col = col + 1
		}
		if heatGen == 1 {
			col = col + 1

		}
	}
	fmt.Println("You rolled a", resultT, "and the heat generated", heatT, "for a total of", (resultT + heatT), "and", col, "collateral was caused")
}
