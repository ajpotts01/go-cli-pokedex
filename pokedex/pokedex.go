package pokedex

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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

func setUrlConfig(locationData locations, config *CommandConfig) {
	if locationData.Previous != nil {
		config.Prev = *locationData.Previous
	}

	if locationData.Next != nil {
		config.Next = *locationData.Next
	}
}

func printLocations(locationData locations) {
	fmt.Println("Moving map")
	if locationData.Count > 0 {
		for _, nextLoc := range locationData.Results {
			fmt.Println(nextLoc.Name)
		}
	} else {
		fmt.Println("No locations")
	}
}

func locationRequest(url string, config *CommandConfig) (locations, error) {
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

func RequestMap(direction string, config *CommandConfig) error {
	var url string

	if direction == "back" {
		if len(config.Prev) == 0 {
			return errors.New("can't go back: you are at the starting location")
		}
		url = config.Prev
	}

	if direction == "forward" {
		if len(config.Next) == 0 && len(config.Prev) > 0 {
			return errors.New("can't go forward: you are at the final location")
		}
		url = config.Next
	}

	if len(url) == 0 {
		url = "https://pokeapi.co/api/v2/location/"
	}

	locationData, err := locationRequest(url, config)

	if err == nil {
		printLocations(locationData)
		setUrlConfig(locationData, config)
	} else {
		return err
	}

	return nil
}
