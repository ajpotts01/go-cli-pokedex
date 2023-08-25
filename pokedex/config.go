package pokedex

type CommandConfig struct {
	Next string
	Prev string
}

type Command struct {
	Name   string
	Desc   string
	Method func(config *CommandConfig) error
}
