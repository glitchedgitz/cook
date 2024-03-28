package config

import (
	"path"
	"path/filepath"
	"strings"
)

func (conf *Config) checkFileSet(p string, array *[]string) bool {
	if files, exists := conf.Ingredients["files"][p]; exists {

		conf.CheckFileCache(p, files)
		FileValues(path.Join(conf.CachePath, p), array)
		return true

	} else if files, exists := conf.Ingredients["raw-files"][p]; exists {

		tmp := make(map[string]bool)
		for _, file := range files {
			if strings.HasPrefix(file, "http://") || strings.HasPrefix(file, "https://") {
				RawFileValues(path.Join(conf.ConfigPath, filepath.Base(file)), tmp)
			} else {
				RawFileValues(file, tmp)
			}
		}

		for k := range tmp {
			*array = append(*array, k)
		}
		return true

	}

	return false
}

func (conf *Config) CheckYaml(p string, array *[]string) bool {

	if val, exists := conf.Ingredients["lists"][p]; exists {
		*array = append(*array, val...)
		return true
	}

	if val, exists := conf.Ingredients["ports"][p]; exists {
		ParsePorts(val, array)
		return true
	}

	if conf.checkFileSet(p, array) {
		return true
	}

	return false
}
