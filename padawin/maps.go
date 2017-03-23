package main

import ("fmt"; "bufio"; "os"; "strconv"; "strings")

func main() {
	fmt.Println("Please type every player's names and the scores of their matches")
	reader := bufio.NewReader(os.Stdin)
	players := make(map[string][]int)
	for {
		var name string
		for {
			var nameError error
			fmt.Print("Player name (type 'done' when you are done): ")
			name, nameError = reader.ReadString('\n')
			name = strings.TrimSpace(name)
			if name != "" {
				break
			} else if nameError != nil {
				fmt.Println(nameError)
			}
		}
		if name == "done" {
			break
		}

		for {
			fmt.Print("Player score: ")
			text, _ := reader.ReadString('\n')
			score, err := strconv.Atoi(strings.TrimSpace(text))
			if err == nil && score >= 0 {
				players[name] = append(players[name], score)
				break
			}
			fmt.Println("Positive integer value expected")
		}
	}

	fmt.Println("Summary:")
	var winner string = ""
	var winnerScore = 0
	for name, scores := range players {
		fmt.Println("Player: ", name)
		fmt.Println("Scores: ")
		nbScores := len(scores)
		var sumScores int = 0
		for s := 0; s < nbScores; s++ {
			if s > 0 {
				fmt.Print(", ")
			}
			fmt.Print(scores[s])
			sumScores += scores[s]
		}

		fmt.Println()
		fmt.Println("Total score: ", sumScores)
		fmt.Println()

		if winner == "" || sumScores > winnerScore {
			winner = name
			winnerScore = sumScores
		}
	}

	fmt.Println("The winner is: ", winner)
}
