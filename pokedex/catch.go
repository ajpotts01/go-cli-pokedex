package pokedex

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"pokecache"
)

func catchRequest(url string, config *CommandConfig, cache *pokecache.Cache) (pokemon, error) {
	var pokemonData pokemon
	var rawData []byte

	rawData, err := retrieveData(url, cache)
	if err != nil {
		return pokemonData, errors.New("failed to retrieve explore data")
	}

	err = json.Unmarshal(rawData, &pokemonData)

	if err != nil {
		fmt.Println(err.Error())
		return pokemonData, errors.New("failed to unmarshal explore data")
	}

	return pokemonData, nil

}

func catch(newPokemon pokemon) bool {
	baseXp := newPokemon.BaseExperience
	successRate := int(float64(baseXp) * 0.5)

	attempt := rand.Intn(baseXp)

	fmt.Printf("Base XP: %v\n", baseXp)
	fmt.Printf("Success at: %v\n", successRate)
	fmt.Printf("You rolled: %v\n", attempt)

	if attempt >= successRate {
		fmt.Println("Pokemon caught!")
		return true
	}

	return false
}

func CatchAttempt(pokemon string, config *CommandConfig, cache *pokecache.Cache, userPokedex *Pokedex) error {
	const baseUrl = "https://pokeapi.co/api/v2/pokemon/"

	url := baseUrl + pokemon

	pokemonData, err := catchRequest(url, config, cache)

	if err != nil {
		return err
	}

	// Catch attempt
	success := catch(pokemonData)

	if success {
		userPokedex.pokedex[pokemon] = pokemonData
	}

	fmt.Println("Current Pokedex:")
	for _, val := range userPokedex.pokedex {
		fmt.Printf("\t - %s\n", val.Name)
	}

	return nil
}
