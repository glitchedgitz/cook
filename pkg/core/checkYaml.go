package core

import (
	"path"
	"path/filepath"
	"strings"
)

func CheckYaml(p string, array *[]string) bool {
	if val, exists := M["charSet"][p]; exists {
		*array = append(*array, strings.Split(val[0], "")...)
		return true
	}

	if files, exists := M["files"][p]; exists {
		for _, file := range files {
			if strings.HasPrefix(file, "http://") || strings.HasPrefix(file, "https://") {
				CheckFileCache(file)
				FileValues(path.Join(home, ".cache", "cook", filepath.Base(file)), array)
			} else {
				FileValues(file, array)
			}
		}
		return true
	}

	if val, exists := M["lists"][p]; exists {
		*array = append(*array, val...)
		return true
	}

	if val, exists := M["extensions"][p]; exists {
		for _, ext := range val {
			*array = append(*array, "."+ext)
		}
		return true
	}
	return false
}
