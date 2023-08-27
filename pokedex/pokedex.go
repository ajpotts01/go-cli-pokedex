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
* TYPES
***************/
type Pokedex struct {
	pokedex map[string]pokemon
}

type pokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Forms          []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	Height    int `json:"height"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	ID                     int    `json:"id"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			VersionGroup struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name      string `json:"name"`
	Order     int    `json:"order"`
	PastTypes []any  `json:"past_types"`
	Species   struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      string `json:"back_default"`
		BackFemale       string `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  string `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      string `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale string `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault string `json:"front_default"`
				FrontFemale  any    `json:"front_female"`
			} `json:"dream_world"`
			Home struct {
				FrontDefault     string `json:"front_default"`
				FrontFemale      string `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale string `json:"front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault      string `json:"back_default"`
					BackGray         string `json:"back_gray"`
					BackTransparent  string `json:"back_transparent"`
					FrontDefault     string `json:"front_default"`
					FrontGray        string `json:"front_gray"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault      string `json:"back_default"`
					BackGray         string `json:"back_gray"`
					BackTransparent  string `json:"back_transparent"`
					FrontDefault     string `json:"front_default"`
					FrontGray        string `json:"front_gray"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"yellow"`
			} `json:"generation-i"`
			GenerationIi struct {
				Crystal struct {
					BackDefault           string `json:"back_default"`
					BackShiny             string `json:"back_shiny"`
					BackShinyTransparent  string `json:"back_shiny_transparent"`
					BackTransparent       string `json:"back_transparent"`
					FrontDefault          string `json:"front_default"`
					FrontShiny            string `json:"front_shiny"`
					FrontShinyTransparent string `json:"front_shiny_transparent"`
					FrontTransparent      string `json:"front_transparent"`
				} `json:"crystal"`
				Gold struct {
					BackDefault      string `json:"back_default"`
					BackShiny        string `json:"back_shiny"`
					FrontDefault     string `json:"front_default"`
					FrontShiny       string `json:"front_shiny"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"gold"`
				Silver struct {
					BackDefault      string `json:"back_default"`
					BackShiny        string `json:"back_shiny"`
					FrontDefault     string `json:"front_default"`
					FrontShiny       string `json:"front_shiny"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"silver"`
			} `json:"generation-ii"`
			GenerationIii struct {
				Emerald struct {
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"emerald"`
				FireredLeafgreen struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"firered-leafgreen"`
				RubySapphire struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"ruby-sapphire"`
			} `json:"generation-iii"`
			GenerationIv struct {
				DiamondPearl struct {
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackDefault      string `json:"back_default"`
						BackFemale       string `json:"back_female"`
						BackShiny        string `json:"back_shiny"`
						BackShinyFemale  string `json:"back_shiny_female"`
						FrontDefault     string `json:"front_default"`
						FrontFemale      string `json:"front_female"`
						FrontShiny       string `json:"front_shiny"`
						FrontShinyFemale string `json:"front_shiny_female"`
					} `json:"animated"`
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationVi struct {
				OmegarubyAlphasapphire struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			GenerationVii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  any    `json:"front_female"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationViii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  string `json:"front_female"`
				} `json:"icons"`
			} `json:"generation-viii"`
		} `json:"versions"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

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

func NewPokedex() Pokedex {
	newPokedex := Pokedex{
		pokedex: make(map[string]pokemon),
	}

	return newPokedex
}

func PrintPokedex(currentPokedex *Pokedex) {
	fmt.Println("Current Pokedex:")
	for _, val := range currentPokedex.pokedex {
		fmt.Printf("\t - %s\n", val.Name)
	}
}

func InspectPokedex(targetPokemon string, currentPokedex *Pokedex) error {
	retrievedPokemon, ok := currentPokedex.pokedex[targetPokemon]

	if !ok {
		fmt.Println("You have not caught this Pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", retrievedPokemon.Name)
	fmt.Printf("Height: %v\n", retrievedPokemon.Height)
	fmt.Printf("Weight: %v\n", retrievedPokemon.Weight)

	if len(retrievedPokemon.Stats) > 0 {
		fmt.Printf("Stats:\n")
		for _, stat := range retrievedPokemon.Stats {
			fmt.Printf("\t -%s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
	}

	if len(retrievedPokemon.Types) > 0 {
		fmt.Printf("Types:\n")
		for _, pokeType := range retrievedPokemon.Types {
			fmt.Printf("\t -%s\n", pokeType.Type.Name)
		}
	}

	return nil
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

		httpResult, err := io.ReadAll(httpResponse.Body)
		httpResponse.Body.Close()

		if err != nil {
			return result, errors.New("failed to read location data")
		}

		result = httpResult
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
		fmt.Println(err.Error())
		return locationData, errors.New("failed to unmarshal explore data")
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
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	locationData, err := locationRequest(url, config, cache)

	if err != nil {
		return err
	}

	printLocations(locationData)
	setUrlConfig(locationData, config)

	return nil
}
