package pokedex

import (
	"pokecache"
)

type CommandConfig struct {
	Next string
	Prev string
}

type Command struct {
	Name   string
	Desc   string
	Method func(extraParam string, config *CommandConfig, cache *pokecache.Cache) error
}
