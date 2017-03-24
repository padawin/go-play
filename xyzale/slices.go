package main

import (
	"math/rand"
	"fmt"
	"bufio"
	"os"
	"strings"
	"time"
)

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	reader := bufio.NewReader(os.Stdin)
	println("In this game i will take a word and I'll shuffle its letters")
	println("Find the word in max 3 attempts")
	println("--------------------------------")
	s := make([]string, 0)
	s = append(s, "a", "w", "k", "w", "a", "r", "d")
	c := make([]string, len(s))
	copy(c, s)
	for i := range s {
		j := r1.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}

	println("SHUFFLED WORD:")
	fmt.Println(s)
	for i := 0; i < 3; i++ {
		println("Attempt ", i + 1, ":")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)
		answer = strings.ToLower(answer)

		if (answer == "awkward") {
			println("YAY! Answer correct!")
			fmt.Println("Answer:", c)
			return
		} else {
			println("Oh snap! Wrong answer, try again!")
			println("--------------------------------")
		}
	}

	println("Shame on you. The correct answer was:")
	fmt.Println(c)
	println("just like this moment for you...")
}
