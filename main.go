package main

import (
	"fmt"

	"github.com/adrichey/cassidoo-cribbage-challenge/cribbage"
)

func main() {
	score, err := cribbage.ScoreHand([]string{"7H", "8C", "9D", "JH", "KS"})

	if err != nil {
		panic(err)
	}

	fmt.Println(`Score for the cribbage hand "7H", "8C", "9D", "JH", "KS":`, score)

	score, err = cribbage.ScoreHand([]string{"AH", "2C", "3D", "4S", "5H"})

	if err != nil {
		panic(err)
	}

	fmt.Println(`Score for the cribbage hand "AH", "2C", "3D", "4S", "5H":`, score)
}
