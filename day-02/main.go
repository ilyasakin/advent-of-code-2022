package main

import (
	"fmt"
	"strconv"
	"strings"
)

const PLAYER_ONE = 1
const PLAYER_TWO = 0

const DRAW_SCORE = 3
const LOST_SCORE = 0
const VICTORY_SCORE = 6

var rock = [2]string{"A", "X"}
var paper = [2]string{"B", "Y"}
var scissors = [2]string{"C", "Z"}

func parseStrategy(rawStrategy string) [][2]string {
	strategySplit := strings.Split(rawStrategy, "\n")
	var strategy [][2]string

	for _, rawRound := range strategySplit {
		rawRoundWoSpace := strings.Replace(rawRound, " ", "", -1)
		round := [2]string{string(rawRoundWoSpace[0]), string(rawRoundWoSpace[1])}
		strategy = append(strategy, round)
	}

	return strategy
}

// Only return score for player 1, since
// task only wants us to calculate player 1's score
func getScoreForHand(hand string) int {
	if hand == rock[PLAYER_ONE] {
		return 1
	}

	if hand == paper[PLAYER_ONE] {
		return 2
	}

	if hand == scissors[PLAYER_ONE] {
		return 3
	}

	return 0
}

// This could be done via a map instead of a bunch of if statements
func evaluateRound(round [2]string) int {
	initialScoreForHand := getScoreForHand(round[PLAYER_ONE])

	if round[PLAYER_ONE] == rock[PLAYER_ONE] && round[PLAYER_TWO] == scissors[PLAYER_TWO] {
		return initialScoreForHand + VICTORY_SCORE
	}

	if round[PLAYER_ONE] == rock[PLAYER_ONE] && round[PLAYER_TWO] == paper[PLAYER_TWO] {
		return initialScoreForHand + LOST_SCORE
	}

	if round[PLAYER_ONE] == rock[PLAYER_ONE] && round[PLAYER_TWO] == rock[PLAYER_TWO] {
		return initialScoreForHand + DRAW_SCORE
	}

	if round[PLAYER_ONE] == paper[PLAYER_ONE] && round[PLAYER_TWO] == rock[PLAYER_TWO] {
		return initialScoreForHand + VICTORY_SCORE
	}

	if round[PLAYER_ONE] == paper[PLAYER_ONE] && round[PLAYER_TWO] == scissors[PLAYER_TWO] {
		return initialScoreForHand + LOST_SCORE
	}

	if round[PLAYER_ONE] == paper[PLAYER_ONE] && round[PLAYER_TWO] == paper[PLAYER_TWO] {
		return initialScoreForHand + DRAW_SCORE
	}

	if round[PLAYER_ONE] == scissors[PLAYER_ONE] && round[PLAYER_TWO] == rock[PLAYER_TWO] {
		return initialScoreForHand + LOST_SCORE
	}

	if round[PLAYER_ONE] == scissors[PLAYER_ONE] && round[PLAYER_TWO] == paper[PLAYER_TWO] {
		return initialScoreForHand + VICTORY_SCORE
	}

	if round[PLAYER_ONE] == scissors[PLAYER_ONE] && round[PLAYER_TWO] == scissors[PLAYER_TWO] {
		return initialScoreForHand + DRAW_SCORE
	}

	return 0
}

func main() {
	rawStrategy := `A Y
B X
C Z`

	strategy := parseStrategy(rawStrategy)
	totalScore := 0

	for _, round := range strategy {
		totalScore = totalScore + evaluateRound(round)
	}

	resultText := "Total score for strategy guide is :score"

	fmt.Print(strings.Replace(resultText, ":score", strconv.Itoa(totalScore), -1))
}
