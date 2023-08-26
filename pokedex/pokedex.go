package pokedex

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"pokecache"
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
		fmt.Printf("Setting backwards nav: %s \n", *locationData.Previous)
		config.Prev = *locationData.Previous
	}

	if locationData.Next != nil {
		fmt.Printf("Setting forwards nav: %s \n", *locationData.Next)
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

func locationRequest(url string, config *CommandConfig, cache *pokecache.Cache) (locations, error) {
	var locationData locations
	var result []byte

	if len(url) == 0 {
		url = "https://pokeapi.co/api/v2/location/"
	}

	fmt.Printf("Attempting to get cache from %s \n", url)
	result, ok := cache.Get(url)
	// TODO: Maybe turn this into its own function to return []byte
	if !ok {
		fmt.Println("Requested page was not cached - getting from web...")
		httpResponse, err := http.Get(url)

		if err != nil {
			return locationData, errors.New("failed to get location data")
		}

		httpResult, err := io.ReadAll(httpResponse.Body)
		httpResponse.Body.Close()

		if err != nil {
			return locationData, errors.New("failed to read location data")
		}

		result = httpResult
		cache.Add(url, result)
	}

	err := json.Unmarshal(result, &locationData)

	if err != nil {
		return locationData, errors.New("failed to unmarshal location data")
	}

	return locationData, nil
}

func RequestMap(direction string, config *CommandConfig, cache *pokecache.Cache) error {
	var url string

	if direction == "back" {
		if len(config.Prev) == 0 {
			return errors.New("can't go back: you are at the starting location")
		}
		url = config.Prev
	}

	if direction == "next" {
		if len(config.Next) == 0 && len(config.Prev) > 0 {
			return errors.New("can't go forward: you are at the final location")
		}
		url = config.Next
	}

	if len(url) == 0 {
		url = "https://pokeapi.co/api/v2/location/"
	}

	locationData, err := locationRequest(url, config, cache)

	if err == nil {
		printLocations(locationData)
		setUrlConfig(locationData, config)
	} else {
		return err
	}

	return nil
}
