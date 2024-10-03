package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	fmt.Print("Pokedex > ")

	for scanner.Scan() {
		input := strings.ToLower(scanner.Text())

		if command, exists := commands[input]; exists {
			err := command.callback()
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}

		fmt.Print("Pokedex > ")
	}
}
