package main

import ("fmt"; "bufio"; "os"; "strings"; "time")

func main() {
    fmt.Println("At each turn, type a letter or a character to execute an action")
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("What would you like to do? ")
        action, _ := reader.ReadString('\n')
        action = strings.TrimSpace(action)
        switch (action) {
            default:
                fmt.Println("Can you speak english please and give me proper instructions?")
                break
            case "k":
                fmt.Println("You go up")
                break
            case "j":
                fmt.Println("You go down")
                break
            case "h":
                fmt.Println("You go left")
                break
            case "l":
                fmt.Println("You go right")
                break
            case "e":
                fmt.Println("You eat a biscuit you found in your pocket")
                break
            case "q":
                fmt.Println("Bye bye!")
                return
            case "x":
                fmt.Println("Remember, 'X' marks the spot! ... you start to dig like a crazy")
                break
        }
    }
}
