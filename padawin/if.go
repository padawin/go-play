package main

import ("bufio"; "fmt"; "os"; "math/rand"; "time"; "strconv"; "strings")

func main() {
	const turns = 10
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println("Find my number between 0 and 100, you have 10 turns!")
	var solution int = r1.Intn(100)
	fmt.Println(solution)
	reader := bufio.NewReader(os.Stdin)
	for turn := 0; turn < turns; turn += 1 {
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		textI, _ := strconv.Atoi(strings.TrimSpace(text))
		if textI < solution {
			fmt.Println("Too low")
		} else if textI > solution {
			fmt.Println("Too high")
		} else {
			fmt.Println("Too correct!")
			return
		}
	}
	fmt.Print("Too bad")
}
