package pokedex

import (
	"fmt"
	"os"
	"pokecache"
	"strings"
)

type CommandConfig struct {
	Next string
	Prev string
}

type Command struct {
	Name   string
	Desc   string
	Method func(extraParam string, config *CommandConfig, cache *pokecache.Cache, currentPokedex *Pokedex) error
}

func getCommands() map[string]Command {
	return map[string]Command{
		"catch": {
			Name:   "Catch",
			Desc:   "Attempt to catch any Pokemon and add it to your Pokedex",
			Method: requestCatchAttempt,
		},
		"inspect": {
			Name:   "Inspect",
			Desc:   "Inspect any Pokemon in your Pokedex",
			Method: requestInspect,
		},
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

func requestExit(extraParam string, config *CommandConfig, cache *pokecache.Cache, currentPokedex *Pokedex) error {
	fmt.Println("Exiting")
	os.Exit(0)
	return nil
}

func requestHelp(extraParam string, config *CommandConfig, cache *pokecache.Cache, currentPokedex *Pokedex) error {
	fmt.Println("Help")
	printCommands()
	return nil
}

func requestMapForward(extraParam string, config *CommandConfig, cache *pokecache.Cache, currentPokedex *Pokedex) error {
	return RequestMap("next", config, cache)
}

func requestMapBackward(extraParam string, config *CommandConfig, cache *pokecache.Cache, currentPokedex *Pokedex) error {
	return RequestMap("back", config, cache)
}

func requestExplore(extraParam string, config *CommandConfig, cache *pokecache.Cache, currentPokedex *Pokedex) error {
	return RequestExplore(extraParam, config, cache)
}

func requestCatchAttempt(extraParam string, config *CommandConfig, cache *pokecache.Cache, currentPokedex *Pokedex) error {
	return CatchAttempt(extraParam, config, cache, currentPokedex)
}

func requestInspect(extraParam string, config *CommandConfig, cache *pokecache.Cache, currentPokedex *Pokedex) error {
	return InspectPokedex(extraParam, currentPokedex)
}

func printCommands() {
	validCommands := getCommands()
	for _, cmd := range validCommands {
		fmt.Println(cmd.Name)
		fmt.Println(cmd.Desc)
	}
}

func HandleRequest(request string, config *CommandConfig, cache *pokecache.Cache, currentPokedex *Pokedex) {
	var extraParam string
	validCommands := getCommands()

	requests := strings.Fields(request)

	mainRequest := requests[0]

	if len(requests) > 1 {
		extraParam = requests[1]
	}

	command, ok := validCommands[mainRequest]
	if ok {
		err := command.Method(extraParam, config, cache, currentPokedex)

		if err != nil {
			fmt.Println(err.Error())
		}

	} else {
		fmt.Println("Invalid command")
	}
}
