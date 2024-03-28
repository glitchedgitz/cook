package config

import "sync"

type Config struct {
	wg               sync.WaitGroup
	ConfigInfo       string
	Ingredients      map[string]map[string][]string
	CheckIngredients map[string][]string
	InputFile        map[string]bool
	HomeFolder       string
	ReConfigure      bool
	ConfigPath       string
	CachePath        string
	IngredientsPath  string
	Verbose          bool
}
