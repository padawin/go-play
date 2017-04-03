package main

import (
	"fmt"
	"strconv"
	"strings"
	"bufio"
	"os"
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)
	deck := make(map[string]int)
	stack := 100

	suits := make([]string, 4)
	suits[0] = "♣"
	suits[1] = "♠"
	suits[2] = "♥"
	suits[3] = "♦"

	values := make([]string, 0)
	values = append(values, "A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K")

	fmt.Println(suits)

	for i := 0; i < len(suits); i++ {
		for j := 0; j < len(values); j++ {
			index := values[j] + suits[i]
			value := 0

			if (values[j] == "A") {
				value = 11
			} else if (values[j] == "J" || values[j] == "Q" || values[j] == "K") {
				value = 10
			} else {
				value, _ = strconv.Atoi(values[j])
			}

			deck[index] = value
		}
	}

	fmt.Println("Hello gambler, ready to waste your money?")
	fmt.Println("This is Black Jack Game. You have £100. Every time you win you double your bet, else you lose it.")

	for true {
		fmt.Println("---------------------")
		fmt.Println("Place your bet, max £", stack, ":")
		text, _ := reader.ReadString('\n')
		bet, _ := strconv.Atoi(strings.TrimSpace(text))

		if (bet > stack) {
			fmt.Println("You cannot bet more than what you have. Sell your house first.")
			continue
		}

		// pick one random card
		randomCard := GetRand(deck)
		userAmount := deck[randomCard]
		playerCards := make([]string, 0)
		playerCards = append(playerCards, randomCard)
		println("Here your first card. Ask for one more card (hit) or stand:")
		fmt.Println(playerCards)
		bust := false

		// ask user cards until he gives up or busts
		for (true) {
			println("Type hit or stand:")
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSpace(answer)
			answer = strings.ToLower(answer)
			if (answer != "hit" && answer != "stand") {
				continue
			}
			if (answer == "hit") {
				randomCard := GetRand(deck)
				playerCards = append(playerCards, randomCard)
				fmt.Println(playerCards)
				userAmount = userAmount + deck[randomCard];
				if userAmount > 21 {
					bust = true
					break;
				}
			}
			if (answer == "stand") {
				break
			}
		}

		if (bust) {
			println("You busted!! Try again!!")
			stack = stack - bet
			if (stack == 0) {
				fmt.Println(suits)
				fmt.Println("| YOU LOSE -- GAME OVER |")
				fmt.Println(suits)
				break
			}
			continue
		}
		println()

		// if user does not bust, so draw cards for the dealer until the value is < of user's value
		dealerAmount := 0;
		dealerCards := make([]string, 0)

		// dealer first card
		randomCard = GetRand(deck)
		dealerAmount = dealerAmount + deck[randomCard];
		dealerCards = append(dealerCards, randomCard)

		fmt.Println("-- Dealer starts with: --")
		fmt.Println(dealerCards);
		// dealer hits until he's below the player's amount
		for dealerAmount <= userAmount {
			time.Sleep(1500 * time.Millisecond)
			randomCard = GetRand(deck)
			dealerAmount = dealerAmount + deck[randomCard];
			dealerCards = append(dealerCards, randomCard)
			fmt.Println("-Dealer hits:", randomCard)
		}

		fmt.Println("-- Dealer cards: --")
		fmt.Println(dealerCards);

		if (dealerAmount < userAmount || dealerAmount > 21) {
			// player wins
			fmt.Println("★★ YOU WON £", bet, " ★★")
			stack += bet
		} else {
			// dealer wins
			fmt.Println("^^^ Oh, Snap!! You lost £", bet)
			stack -= bet
		}

		// if stack == 0 break
		if (stack == 0) {
			fmt.Println(suits)
			fmt.Println("| YOU LOSE -- GAME OVER |")
			fmt.Println(suits)
			break
		}
	}
}

func GetRand(a map[string]int) string {
	// produce a pseudo-random number between 0 and len(a)-1
	i := int(float32(len(a)) * rand.Float32())
	for x := range a {
		if i == 0 {
			return x
		} else {
			i--
		}
	}
	panic("impossible")
}
