package main

import (
	"strconv"
	"strings"
	"bufio"
	"os"
)

func main() {
	correct_answers := 0;
	println("This is a quiz. Give as more correct answers as you can.")
	spacer()
	if (question1()) {
		correct_answers++
	}
	if (question2()) {
		correct_answers++
	}
	if (question3()) {
		correct_answers++
	}
	spacer()
	println("You gave ", correct_answers, " correct answers out of 3")
}

func spacer() {
	println("-----------\n")
}

func question1() bool {
	reader := bufio.NewReader(os.Stdin)
	println("QUESTION 1:")
	println("Which came first? The egg or the chicken?")
	println("Type 0 for Egg, 1 for Chicken")
	text, _ := reader.ReadString('\n')
	answer, _ := strconv.Atoi(strings.TrimSpace(text))

	switch answer {
	case 0:
		println("Wrong! Go back to school!")
		spacer()
		return false
		break
	case 1:
		println("Correct! \n Researchers found that the formation of egg shells relies on a protein found only in a chicken's ovaries.")
		println("Learn more here: http://www.dailymail.co.uk/sciencetech/article-1294341/Chicken-really-DID-come-egg-say-scientists.html")
		spacer()
		return true
		break
	default:
		println("You can enter 0 or 1 only")
		return question1()
		break
	}

	return false
}

func question2() bool {
	reader := bufio.NewReader(os.Stdin)
	println("QUESTION 2:")
	println("What's the surname of Emma, the Spice Girls singer?")
	println("Type 0 for Chisholm, 1 for Bunton, 2 for Brown")
	text, _ := reader.ReadString('\n')
	answer, _ := strconv.Atoi(strings.TrimSpace(text))

	switch answer {
	case 0, 2:
		println("Wrong! You are a goat!")
		spacer()
		return false
		break
	case 1:
		println("Correct! \n I was pretty sure you were a Spice Girls fan")
		println("Not a heavy metal fan: http://i1082.photobucket.com/albums/j365/stormdogy117/Random%20stuff/HugeMetalFan.jpg")
		spacer()
		return true
		break
	default:
		println("You can enter 0, 1 or 2 only")
		return question2()
		break
	}

	return false
}

func question3() bool {
	reader := bufio.NewReader(os.Stdin)
	println("QUESTION 3:")
	println("What flavour is Cointreau?")
	println("Type 0 for Apple, 1 for Lemon, 2 for Orange")
	text, _ := reader.ReadString('\n')
	answer, _ := strconv.Atoi(strings.TrimSpace(text))

	switch answer {
	case 0, 1:
		println("Wrong! You have to drink more!")
		spacer()
		return false
		break
	case 2:
		println("Correct! \n To know that either you are french or you are a drunkard, or maybe both")
		spacer()
		return true
		break
	default:
		println("You can enter 0, 1 or 2 only")
		return question3()
		break
	}

	return false
}