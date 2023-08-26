package main

import (
	"fmt"
	"os"
	"pokecache"
	"pokedex"
	"strings"
	"time"
)

func requestExit(extraParam string, config *pokedex.CommandConfig, cache *pokecache.Cache) error {
	fmt.Println("Exiting")
	os.Exit(0)
	return nil
}

func requestHelp(extraParam string, config *pokedex.CommandConfig, cache *pokecache.Cache) error {
	fmt.Println("Help")
	printCommands()
	return nil
}

func requestMapForward(extraParam string, config *pokedex.CommandConfig, cache *pokecache.Cache) error {
	return pokedex.RequestMap("next", config, cache)
}

func requestMapBackward(extraParam string, config *pokedex.CommandConfig, cache *pokecache.Cache) error {
	return pokedex.RequestMap("back", config, cache)
}

func requestExplore(extraParam string, config *pokedex.CommandConfig, cache *pokecache.Cache) error {
	return pokedex.RequestExplore(extraParam, config, cache)
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
		"explore": {
			Name:   "Explore",
			Desc:   "Explore an area for Pokemon",
			Method: requestExplore,
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
	var extraParam string
	validCommands := getCommands()

	requests := strings.Split(request, " ")

	if len(requests) > 1 {
		request = requests[0]
		extraParam = requests[1]
	}
	command, ok := validCommands[request]
	if ok {
		err := command.Method(extraParam, config, cache)

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
