package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type locations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type commandConfig struct {
	next string
	prev string
}

type command struct {
	name   string
	desc   string
	method func(config *commandConfig) error
}

func requestExit(config *commandConfig) error {
	fmt.Println("Exiting")
	os.Exit(0)
	return nil
}

func requestHelp(config *commandConfig) error {
	fmt.Println("Help")
	printCommands()
	return nil
}

func setUrlConfig(locationData locations, config *commandConfig) {
	config.prev = *locationData.Previous
	config.next = *locationData.Next
}

func locationRequest(url string, config *commandConfig) (locations, error) {
	var locationData locations

	if len(url) == 0 {
		url = "https://pokeapi.co/api/v2/location/"
	}
	result, err := http.Get(url)

	if err != nil {
		return locationData, errors.New("failed to get location data")
	}

	resultBody, err := io.ReadAll(result.Body)
	result.Body.Close()

	if err != nil {
		return locationData, errors.New("failed to read location data")
	}

	err = json.Unmarshal(resultBody, &locationData)

	if err != nil {
		return locationData, errors.New("failed to unmarshal location data")
	}

	return locationData, nil
}

func requestMapBack(config *commandConfig) error {
	var url string = config.prev

	if len(url) == 0 {
		return errors.New("can't go back: you are at the starting location")
	}

	// Go backward
	fmt.Println("Going backward")
	
	locationData, err := locationRequest(url, config)
	if err == nil {
		setUrlConfig(locationData, config)
	}

	return nil
}

func requestMap(config *commandConfig) error {
	nextUrl := config.next
	backUrl := config.prev

	if len(nextUrl) == 0 && len(backUrl) > 0 {
		return errors.New("can't go forward: you are at the final location")
	}

	// Go forward
	fmt.Println("Going forward")
	locationData, err := locationRequest(nextUrl, config)
	if err == nil {
		setUrlConfig(locationData, config)
	}

	return nil
}

func getCommands() map[string]command {
	return map[string]command{
		"map": {
			name:   "Map",
			desc:   "Display next 20 locations",
			method: requestMap,
		},
		"mapb": {
			name:   "Map Back",
			desc:   "Display previous 20 locations",
			method: requestMapBack,
		},
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

func handleRequest(request string, config *commandConfig) {
	validCommands := getCommands()

	command, ok := validCommands[request]
	if ok {
		err := command.method(config)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(config.next)
			fmt.Println(config.prev)
		}
	} else {
		fmt.Println("Invalid command")
	}
}

func main() {
	var config commandConfig

	for {
		fmt.Print("pokedex > ")

		var input string
		cmd, err := fmt.Scanln(&input)

		cmdStr := fmt.Sprintf("Command: %v", cmd)
		inputStr := fmt.Sprintf("Input: %s", input)

		fmt.Println(cmdStr)
		fmt.Println(inputStr)

		if err == nil {
			handleRequest(input, &config)
		} else {
			print(err)
		}
	}
}
