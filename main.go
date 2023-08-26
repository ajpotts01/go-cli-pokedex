package main

import (
	"fmt"
	"os"
	"pokecache"
	"pokedex"
	"time"
)

func requestExit(config *pokedex.CommandConfig, cache *pokecache.Cache) error {
	fmt.Println("Exiting")
	os.Exit(0)
	return nil
}

func requestHelp(config *pokedex.CommandConfig, cache *pokecache.Cache) error {
	fmt.Println("Help")
	printCommands()
	return nil
}

func requestMapForward(config *pokedex.CommandConfig, cache *pokecache.Cache) error {
	return pokedex.RequestMap("next", config, cache)
}

func requestMapBackward(config *pokedex.CommandConfig, cache *pokecache.Cache) error {
	return pokedex.RequestMap("back", config, cache)
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

func handleRequest(request string, config *pokedex.CommandConfig, cache *pokecache.Cache) {
	validCommands := getCommands()

	command, ok := validCommands[request]
	if ok {
		err := command.Method(config, cache)

		if err != nil {
			fmt.Println(err)
		}

	} else {
		fmt.Println("Invalid command")
	}
}

func main() {
	var config pokedex.CommandConfig
	var cache pokecache.Cache

	cacheTtl := time.Duration(10 * time.Second)
	cache = pokecache.NewCache(cacheTtl)

	for {
		fmt.Print("pokedex > ")

		var input string
		_, err := fmt.Scanln(&input)

		if err == nil {
			handleRequest(input, &config, &cache)
		} else {
			print(err)
		}
	}
}
