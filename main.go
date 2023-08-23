package main

import (
	"fmt"
	"os"
)

type command struct {
	name   string
	desc   string
	method func() error
}

func requestExit() error {
	fmt.Println("Exiting")
	os.Exit(0)
	return nil
}

func requestHelp() error {
	fmt.Println("Help")
	return nil
}

func getCommands() map[string]command {
	return map[string]command{
		"help": {
			name:   "Help",
			desc:   "Help message",
			method: requestHelp,
		},
		"exit": {
			name:   "Exit",
			desc:   "Exit Pokedex",
			method: requestExit,
		},
	}
}

func printCommands() {
	validCommands := getCommands()
	for _, cmd := range validCommands {
		fmt.Println(cmd.name)
		fmt.Println(cmd.desc)
	}
}

func handleRequest(request string) {
	validCommands := getCommands()

	command := validCommands[request]
	command.method()
}

func main() {
	for {
		fmt.Print("pokedex > ")

		var input string
		cmd, err := fmt.Scanln(&input)

		cmdStr := fmt.Sprintf("Command: %v", cmd)
		inputStr := fmt.Sprintf("Input: %s", input)

		fmt.Println(cmdStr)
		fmt.Println(inputStr)

		if err == nil {
			handleRequest(input)
		} else {
			print(err)
		}
	}
}
