package main

import (
	"fmt"
	"os"
	"pokedex"
)

func requestExit(config *pokedex.CommandConfig) error {
	fmt.Println("Exiting")
	os.Exit(0)
	return nil
}

func requestHelp(config *pokedex.CommandConfig) error {
	fmt.Println("Help")
	printCommands()
	return nil
}

func requestMapForward(config *pokedex.CommandConfig) error {
	return pokedex.RequestMap("forward", config)
}

func requestMapBackward(config *pokedex.CommandConfig) error {
	return pokedex.RequestMap("backward", config)
}

func getCommands() map[string]pokedex.Command {
	return map[string]pokedex.Command{
		"map": {
			Name:   "Map",
			Desc:   "Display next 20 locations",
			Method: requestMapForward,
		},
		"mapb": {
			Name:   "Map Back",
			Desc:   "Display previous 20 locations",
			Method: requestMapBackward,
		},
		"help": {
			Name:   "Help",
			Desc:   "Help message",
			Method: requestHelp,
		},
		"exit": {
			Name:   "Exit",
			Desc:   "Exit Pokedex",
			Method: requestExit,
		},
	}
}

func printCommands() {
	validCommands := getCommands()
	for _, cmd := range validCommands {
		fmt.Println(cmd.Name)
		fmt.Println(cmd.Desc)
	}
}

func handleRequest(request string, config *pokedex.CommandConfig) {
	validCommands := getCommands()

	command, ok := validCommands[request]
	if ok {
		err := command.Method(config)

		if err != nil {
			fmt.Println(err)
		}

	} else {
		fmt.Println("Invalid command")
	}
}

func main() {
	var config pokedex.CommandConfig

	for {
		fmt.Print("pokedex > ")

		var input string
		_, err := fmt.Scanln(&input)

		// cmdStr := fmt.Sprintf("Command: %v", cmd)
		// inputStr := fmt.Sprintf("Input: %s", input)

		// fmt.Println(cmdStr)
		// fmt.Println(inputStr)

		if err == nil {
			handleRequest(input, &config)
		} else {
			print(err)
		}
	}
}
