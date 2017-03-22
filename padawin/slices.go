package main

import ("fmt"; "bufio"; "os"; "strings"; "strconv")

func main() {
	var staff []string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Pop101 - The next gen Personel manager")
	fmt.Println("You can do the following commands:")
	fmt.Println("- a/A %name%: Add a member of staff")
	fmt.Println("- l/L: List all the staff")
	fmt.Println("- e/E %id%: Edit a member of staff")
	fmt.Println("- d/D %id%: Delete a member of staff")
	fmt.Println("- v/V %id%: View a member of staff")

	for {
		fmt.Println("")
		fmt.Print("Command: ")
		command, _ := reader.ReadString('\n')
		action := string(command[0])
		argument := strings.TrimSpace(command[2:])
		if action != "l" && action != "L" && argument == "" {
			fmt.Println("Argument missing")
			continue
		}

		switch (action) {
			case "a", "A":
				name := argument
				staff = append(staff, name)
				fmt.Println("Staff ", name, " added. ID is ", len(staff) - 1)
				break
			case "l", "L":
				var length int = len(staff)
				for i := 0; i < length; i++ {
					fmt.Println(i, ": ", staff[i])
				}
				break
			case "e", "E":
				staffId, _ := strconv.Atoi(argument)
				if staffId >= len(staff) {
					fmt.Println("Unknown staff id")
					continue
				}

				fmt.Print("Staff new name: ")
				newName, _ := reader.ReadString('\n')
				staff[staffId] = newName
				break
			case "d", "D":
				staffId, _ := strconv.Atoi(argument)
				if staffId >= len(staff) {
					fmt.Println("Unknown staff id")
					continue
				}

				staff = append(
					staff[:staffId],
					staff[staffId + 1:]...
				)
				break
			case "v", "V":
				staffId, _ := strconv.Atoi(argument)
				if staffId >= len(staff) {
					fmt.Println("Unknown staff id")
					continue
				}

				fmt.Println("Staff member info:")
				fmt.Println("ID: ", staffId)
				fmt.Println("Name: ", staff[staffId])
				break
		}
	}
}
