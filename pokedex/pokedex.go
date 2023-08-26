package pokedex

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"pokecache"
)

/**************
* STRUCTS
***************/
type locations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type locationDetails struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int `json:"chance"`
				ConditionValues []struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"condition_values"`
				MaxLevel int `json:"max_level"`
				Method   struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
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

func printPokemon(locationData locationDetails) {
	fmt.Printf("Exploring %s...\n", locationData.Location.Name)

	if len(locationData.PokemonEncounters) > 0 {
		fmt.Println("Found Pokemon:")
		for _, nextEncounter := range locationData.PokemonEncounters {
			fmt.Printf("\t - %s\n", nextEncounter.Pokemon.Name)
		}
	}
}

func retrieveData(url string, cache *pokecache.Cache) ([]byte, error) {
	var result []byte

	if len(url) == 0 {
		return result, errors.New("must supply a specific URL")
	}

	fmt.Printf("Attempting to get %s from cache \n", url)
	result, ok := cache.Get(url)

	if !ok {
		fmt.Println("Requested page was not cached - getting from web...")
		httpResponse, err := http.Get(url)

		if err != nil {
			return result, errors.New("failed to get location data")
		}

		result, err := io.ReadAll(httpResponse.Body)
		httpResponse.Body.Close()

		if err != nil {
			return result, errors.New("failed to read location data")
		}

		cache.Add(url, result)
	}

	return result, nil
}

func exploreRequest(url string, config *CommandConfig, cache *pokecache.Cache) (locationDetails, error) {
	var locationData locationDetails
	var rawData []byte

	rawData, err := retrieveData(url, cache)
	if err != nil {
		return locationData, errors.New("failed to retrieve explore data")
	}

	err = json.Unmarshal(rawData, &locationData)

	if err != nil {
		return locationData, errors.New("failed to unmarshal location data")
	}

	return locationData, nil
}

func locationRequest(url string, config *CommandConfig, cache *pokecache.Cache) (locations, error) {
	var locationData locations
	var rawData []byte

	// Can do a stock standard pg. 1 version of this if no URL provided
	if len(url) == 0 {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	rawData, err := retrieveData(url, cache)

	if err != nil {
		return locationData, errors.New("failed to retrieve location data")
	}

	err = json.Unmarshal(rawData, &locationData)

	if err != nil {
		return locationData, errors.New("failed to unmarshal location data")
	}

	return locationData, nil
}

func RequestExplore(location string, config *CommandConfig, cache *pokecache.Cache) error {
	const baseUrl = "https://pokeapi.co/api/v2/location-area/"

	url := baseUrl + location

	locationData, err := exploreRequest(url, config, cache)

	if err != nil {
		return err
	}

	printPokemon(locationData)

	return nil
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

	if err != nil {
		return err
	}

	printLocations(locationData)
	setUrlConfig(locationData, config)

	return nil
}
