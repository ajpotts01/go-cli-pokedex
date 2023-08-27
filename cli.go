package main

import (
	"bufio"
	"fmt"
	"os"
	"pokecache"
	"pokedex"
	"time"
)

func startCli() {
	var config pokedex.CommandConfig
	var cache pokecache.Cache
	var userPokedex pokedex.Pokedex

	cacheTtl := time.Duration(5 * time.Second)
	cache = pokecache.NewCache(cacheTtl)

	userPokedex = pokedex.NewPokedex()

	for {
		fmt.Print("pokedex > ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if scanner.Err() == nil {
			pokedex.HandleRequest(input, &config, &cache, &userPokedex)
		} else {
			fmt.Println(scanner.Err().Error())
		}
	}
}
