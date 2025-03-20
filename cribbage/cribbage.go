package cribbage

import (
	"errors"
	"slices"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
)

const errInvalidHandLength = "hand length must equal five"
const errInvalidCardFound = "invalid card found"
const errNumOfCardsOutOfBounds = "numOfCards must be less than the length of cardRankValues"
const errNumOfCardsTooSmall = "numOfCards must be equal to 2 or greater"

var deck []string

func init() {
	deck = make([]string, 0, 52)
	suits := []string{"H", "C", "D", "S"}
	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	for _, s := range suits {
		for _, r := range ranks {
			deck = append(deck, r+s)
		}
	}
}

func ScoreHand(hand []string) (int, error) {
	if len(hand) != 5 {
		return 0, errors.New(errInvalidHandLength)
	}

	score := 0
	cardRanks := make([]string, 0, len(hand))
	cardRankValues := make([]int, 0, len(hand))

	for _, card := range hand {
		if !slices.Contains(deck, card) {
			return 0, errors.New(errInvalidCardFound)
		}

		suit := card[len(card)-1:]
		rank := strings.TrimSuffix(card, suit)
		rankValue, err := getRankValue(rank)

		if err != nil {
			return 0, err
		}

		cardRanks = append(cardRanks, rank)
		cardRankValues = append(cardRankValues, rankValue)
	}

	// Score the pairs, three of a kind, fours first
	matchingRanks := make(map[string]int)
	for _, rank := range cardRanks {
		matchingRanks[rank]++
	}

	for _, numOfCards := range matchingRanks {
		if numOfCards == 2 {
			score += 2
		}

		if numOfCards == 3 {
			score += 6
		}

		if numOfCards == 4 {
			score += 12
		}
	}

	// Score the fifteens for two, three, four, and five cards respectively
	for i := 2; i <= 5; i++ {
		s, err := scoreFifteens(cardRankValues, i)

		if err != nil {
			return 0, err
		}

		score += s
	}

	// Score the runs in the hand
	score += scoreRuns(cardRanks)

	return score, nil
}

func getRankValue(r string) (int, error) {
	if r == "A" {
		return 1, nil
	}

	if r == "J" || r == "Q" || r == "K" {
		return 10, nil
	}

	score, err := strconv.ParseInt(r, 10, 0)

	return int(score), err
}

func scoreFifteens(cardRankValues []int, numOfCards int) (int, error) {
	valuesLength := len(cardRankValues)

	if numOfCards > len(cardRankValues) {
		return 0, errors.New(errNumOfCardsOutOfBounds)
	}

	if numOfCards < 2 {
		return 0, errors.New(errNumOfCardsTooSmall)
	}

	score := 0

	combinations := combin.Combinations(valuesLength, numOfCards)

	for _, v := range combinations {
		comboValue := 0

		for _, i := range v {
			comboValue += cardRankValues[i]
		}

		if comboValue == 15 {
			score += 2
		}
	}

	return score, nil
}

func scoreRuns(cardRanks []string) int {
	if len(cardRanks) <= 2 {
		return 0
	}

	values := make([]int, len(cardRanks))

	for i := 0; i < len(cardRanks); i++ {
		rv, _ := getRankValue(cardRanks[i])

		if cardRanks[i] == "J" {
			rv += 1
		}

		if cardRanks[i] == "Q" {
			rv += 2
		}

		if cardRanks[i] == "K" {
			rv += 3
		}

		values[i] = rv
	}

	slices.Sort(values)

	runCount := 0
	for i := 0; i < len(values); i++ {
		currentRunCount := 1

		for j := i; j < len(values)-1; j++ {
			if values[j]+1 == values[j+1] {
				currentRunCount++
			} else {
				break
			}
		}

		if currentRunCount >= 3 && currentRunCount > runCount {
			runCount = currentRunCount
		}
	}

	return runCount
}
