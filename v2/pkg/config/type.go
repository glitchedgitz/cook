package config

import "sync"

type Config struct {
	wg              sync.WaitGroup
	ConfigInfo      string
	M               map[string]map[string][]string
	CheckM          map[string][]string
	HomeFolder      string
	ReConfigure     bool
	ConfigPath      string
	CachePath       string
	IngredientsPath string
	InputFile       map[string]bool
	Min             int
	Verbose         bool
}
