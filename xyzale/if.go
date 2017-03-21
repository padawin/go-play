package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
	"math/rand"
)

func main() {
	println("This is the three cards game. Enter A, B or C to find the queen and double your bet")
	stack := 100
	threeCards(stack)
}

func threeCards(stack int) {
	if stack == 0 {
		// use exit to avoid else
		println("You lose!")
	} else {
		var position int = random(0, 2)
		var A, B, C, r_answer string
		if (position == 0) {
			A = "Q of hearts"
			B = "7 of spades"
			C = "3 of clubs"
			r_answer = "A"
		} else if (position == 1) {
			A = "7 of spades"
			B = "Q of hearts"
			C = "3 of clubs"
			r_answer = "B"
		} else {
			A = "3 of clubs"
			B = "7 of spades"
			C = "Q of hearts"
			r_answer = "C"
		}
		reader := bufio.NewReader(os.Stdin)
		println("Enter your bet (max: ", stack, "£)")
		text, _ := reader.ReadString('\n')
		bet, _ := strconv.Atoi(strings.TrimSpace(text))
		if (bet > stack) {
			println("You cannot bet more than you have, sell your house first")
			threeCards(stack)
		} else {
			println("You are betting ", bet, "£, find the Queen!")
			println("Enter A, B or C")
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSpace(answer)

			if (strings.ToUpper(answer) == r_answer) {
				stack += bet
				println("Congratulations! You have found the Queen!")
				println("A: ", A, "B: ", B, "C: ", C)
				println("Your current balance is ", stack)
				threeCards(stack)
			} else {
				stack -= bet
				println("----------------------------")
				println("Oh snap! You lost ", bet, "£")
				println("A: ", A, " | B: ", B, " | C: ", C)
				println("Your current balance is ", stack)
				println("----------------------------")
				threeCards(stack)
			}
		}

	}

}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}
