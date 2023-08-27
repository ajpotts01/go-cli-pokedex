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

	cacheTtl := time.Duration(5 * time.Second)
	cache = pokecache.NewCache(cacheTtl)

	for {
		fmt.Print("pokedex > ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		if scanner.Err() == nil {
			pokedex.HandleRequest(input, &config, &cache)
		} else {
			fmt.Println(scanner.Err().Error())
		}
	}
}
