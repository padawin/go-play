package main

import (
	"strconv"
	"strings"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	word := [7]string{"c", "o", "c", "o", "n", "u", "t"}
	shown := [7]bool{false, false, false, false, false, false, false}

	println("--- Find the secret word! ---")
	println("Instructions: The secret word contains ", len(word), " characters")
	println("choose the character position, then enter the character to guess it")
	println("Enter only lowercase letters, uppercase letters will not be considered")
	for true {
		// Print the word
		str := "WORD: "
		for i := 0; i < len(word); i++ {
			if (shown[i]) {
				str = str + word[i] + " "
			} else {
				str = str + "_ "
			}
		}
		println(str)

		// choose letter position
		println("Enter the char position (1-7)")
		text, _ := reader.ReadString('\n')
		position, _ := strconv.Atoi(strings.TrimSpace(text))
		position--

		if (position > 6 || position < 0) {
			println("Position must be between 1 and 7")
			println("--------------------")
			continue
		}


		// guess the letter
		println("Now guess the letter in position ", position + 1)
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		if (answer == word[position]) {
			println("Yay!! This is the right letter!")
			shown[position] = true
		} else {
			println("Oh snap! Try again...")
		}

		// check all letters are guessed
		guessed := true
		for i := 0; i < len(shown); i++ {
			if (!shown[i]) {
				guessed = false
			}
		}

		// if so, congrats and exit
		if (guessed) {
			println("!! CONGRATULATIONS !!")
			println("The secret word was COCONUT!!")
			break;
		}
	}
}
