package main

import(
	"bufio"
	"os"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Contact Manager v1.0")
	fmt.Println("Commands: add, list, search, update, delete, help, exit")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Error reading input: ", err)
			continue
		}

		command := strings.TrimSpace(input)

		switch command {
		case "search":	
			fmt.Print("Enter name to search: ")
			needle, err := reader.ReadString('\n')

			if err != nil {
				fmt.Println("Error reading input: ", err)
				continue
			}

			search(needle)
		case "list":
			list()
		case "help":
			help()
		case "exit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Printf("Command %s is not recognized\n", command)
		}
	}
}

func help() {
	fmt.Println("Usage:")
	fmt.Println("\tadd\t- Add new contact")
	fmt.Println("\tlist\t- Display list of saved contacts")
	fmt.Println("\tsearch\t- Search contacts within given string")
	fmt.Println("\tupdate\t- Update an existing contact")
	fmt.Println("\tdelete\t- Delete a contact")
	fmt.Println("\thelp\t- Display this list")
	fmt.Println("\texit\t- Exit this program")
}
