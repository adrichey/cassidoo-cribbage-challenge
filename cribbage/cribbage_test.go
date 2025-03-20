package cribbage

import (
	"testing"
)

func TestGetRankValue(t *testing.T) {
	value, err := getRankValue("Z")
	if value != 0 || err == nil {
		t.Errorf(`getRankValue("Z") = %v, %v, want 0, error`, value, err)
	}

	value, err = getRankValue("A")
	if value != 1 || err != nil {
		t.Errorf(`getRankValue("A") = %v, %v, want 1, nil`, value, err)
	}

	value, err = getRankValue("J")
	if value != 10 || err != nil {
		t.Errorf(`getRankValue("J") = %v, %v, want 10, nil`, value, err)
	}

	value, err = getRankValue("Q")
	if value != 10 || err != nil {
		t.Errorf(`getRankValue("Q") = %v, %v, want 10, nil`, value, err)
	}

	value, err = getRankValue("K")
	if value != 10 || err != nil {
		t.Errorf(`getRankValue("K") = %v, %v, want 10, nil`, value, err)
	}

	value, err = getRankValue("6")
	if value != 6 || err != nil {
		t.Errorf(`getRankValue("6") = %v, %v, want 6, nil`, value, err)
	}
}

func TestScoreFifteens(t *testing.T) {
	score, err := scoreFifteens([]int{10, 5, 5, 5, 10}, 2)
	if score != 12 || err != nil {
		t.Errorf(`scoreFifteens([]int{10, 5, 5, 5, 10}, 2) = %v, %v, want 12, nil`, score, err)
	}

	score, err = scoreFifteens([]int{9, 6, 8, 5, 7}, 2)
	if score != 4 || err != nil {
		t.Errorf(`scoreFifteens([]int{9, 6, 8, 5, 7}, 2) = %v, %v, want 4, nil`, score, err)
	}

	score, err = scoreFifteens([]int{10, 5, 5, 5, 10}, 3)
	if score != 2 || err != nil {
		t.Errorf(`scoreFifteens([]int{10, 5, 5, 5, 10}, 3) = %v, %v, want 2, nil`, score, err)
	}

	score, err = scoreFifteens([]int{9, 6, 8, 5, 7}, 3)
	if score != 0 || err != nil {
		t.Errorf(`scoreFifteens([]int{9, 6, 8, 5, 7}, 3) = %v, %v, want 0, nil`, score, err)
	}

	score, err = scoreFifteens([]int{1, 4, 5, 10, 5}, 3)
	if score != 2 || err != nil {
		t.Errorf(`scoreFifteens([]int{1, 4, 5, 10, 5}, 3) = %v, %v, want 2, nil`, score, err)
	}

	score, err = scoreFifteens([]int{1, 4, 5, 5, 10}, 4)
	if score != 2 || err != nil {
		t.Errorf(`scoreFifteens([]int{1, 4, 5, 5, 10}, 4) = %v, %v, want 2, nil`, score, err)
	}

	score, err = scoreFifteens([]int{1, 4, 5, 3, 2}, 5)
	if score != 2 || err != nil {
		t.Errorf(`scoreFifteens([]int{1, 4, 5, 3, 2}, 5) = %v, %v, want 2, nil`, score, err)
	}

	score, err = scoreFifteens([]int{1, 4, 5}, 5)
	if score != 0 || err.Error() != errNumOfCardsOutOfBounds {
		t.Errorf(`scoreFifteens([]int{1, 4, 5}, 5) = %v, %v, want 0, %v`, score, err, errNumOfCardsOutOfBounds)
	}

	score, err = scoreFifteens([]int{1, 4, 5, 3, 2}, 1)
	if score != 0 || err.Error() != errNumOfCardsTooSmall {
		t.Errorf(`scoreFifteens([]int{1, 4, 5}, 5) = %v, %v, want 0, %v`, score, err, errNumOfCardsTooSmall)
	}
}

func TestScoreRuns(t *testing.T) {
	score := scoreRuns([]string{"A", "4", "5", "3", "2"})
	if score != 5 {
		t.Errorf(`scoreRuns([]string{"A", "4", "5", "3", "2"}) = %v, want 5`, score)
	}

	score = scoreRuns([]string{"A", "3", "5", "7", "9"})
	if score != 0 {
		t.Errorf(`scoreRuns([]string{"A", "3", "5", "7", "9"}) = %v, want 5`, score)
	}

	score = scoreRuns([]string{"2", "A", "5", "9", "6"})
	if score != 0 {
		t.Errorf(`scoreRuns([]string{"2", "A", "5", "9", "6"}) = %v, want 0`, score)
	}

	score = scoreRuns([]string{"A", "5", "8", "3", "2"})
	if score != 3 {
		t.Errorf(`scoreRuns([]string{"A", "5", "8", "3", "2"}) = %v, want 3`, score)
	}

	score = scoreRuns([]string{"8", "6", "7", "5", "A"})
	if score != 4 {
		t.Errorf(`scoreRuns([]string{"8", "6", "7", "5", "A"}) = %v, want 4`, score)
	}

	score = scoreRuns([]string{"2", "A"})
	if score != 0 {
		t.Errorf(`scoreRuns([]string{"2", "A"}) = %v, want 0`, score)
	}

	score = scoreRuns([]string{"8", "6", "7", "5", "A", "3", "2"})
	if score != 4 {
		t.Errorf(`scoreRuns([]string{"8", "6", "7", "5", "A", "3", "2"}) = %v, want 4`, score)
	}

	score = scoreRuns([]string{"10", "J", "Q", "K", "9"})
	if score != 5 {
		t.Errorf(`scoreRuns([]string{"10", "J", "Q", "K", "9"}) = %v, want 5`, score)
	}
}

func TestScoreHand(t *testing.T) {
	score, err := ScoreHand([]string{"7H", "8C", "9D"})
	if score != 0 || err.Error() != errInvalidHandLength {
		t.Errorf(`ScoreHand([]string{"7H", "8C", "9D"}) = %v, %v, want 0, %v`, score, err, errInvalidHandLength)
	}

	score, err = ScoreHand([]string{"7Y", "8C", "9D", "JH", "KS"})
	if score != 0 || err.Error() != errInvalidCardFound {
		t.Errorf(`ScoreHand([]string{"7Y", "8C", "9D", "JH", "KS"}) = %v, %v, want 0, %v`, score, err, errInvalidCardFound)
	}

	score, err = ScoreHand([]string{"7H", "7C", "9D", "KH", "KS"})
	if score != 4 || err != nil {
		t.Errorf(`ScoreHand([]string{"7H", "7C", "9D", "KH", "KS"}) = %v, %v, want 4, nil`, score, err)
	}

	score, err = ScoreHand([]string{"7H", "7C", "9D", "7H", "KS"})
	if score != 6 || err != nil {
		t.Errorf(`ScoreHand([]string{"7H", "7C", "9D", "7H", "KS"}) = %v, %v, want 6, nil`, score, err)
	}

	score, err = ScoreHand([]string{"7H", "7C", "7D", "7S", "KS"})
	if score != 12 || err != nil {
		t.Errorf(`ScoreHand([]string{"7H", "7C", "7D", "7S", "KS"}) = %v, %v, want 12, nil`, score, err)
	}

	score, err = ScoreHand([]string{"7H", "8C", "9D", "JH", "KS"})
	if score != 5 || err != nil {
		t.Errorf(`ScoreHand([]string{"7H", "8C", "9D", "JH", "KS"}) = %v, %v, want 5, nil`, score, err)
	}

	score, err = ScoreHand([]string{"AH", "2C", "3D", "4S", "5H"})
	if score != 7 || err != nil {
		t.Errorf(`ScoreHand([]string{"AH", "2C", "3D", "4S", "5H"}) = %v, %v, want 7, nil`, score, err)
	}
}
