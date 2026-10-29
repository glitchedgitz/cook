package config

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestParseFilePreservesPaths(t *testing.T) {
	directory := t.TempDir()
	workingDirectory, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(directory); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(workingDirectory); err != nil {
			t.Error(err)
		}
	})

	for _, name := range []string{"file.txt", "index.txt", "list.txt", "example.txt", "candidate.txt", "词表 with spaces.txt"} {
		if err := os.WriteFile(name, []byte("alpha\nbeta\n"), 0600); err != nil {
			t.Fatal(err)
		}
	}
	absolutePath := filepath.Join(directory, "file.txt")
	fileURI := "file:///" + strings.TrimPrefix(filepath.ToSlash(absolutePath), "/")
	cases := []struct {
		name  string
		input string
	}{
		{"absolute", absolutePath},
		{"absolute_spaces_unicode", filepath.Join(directory, "词表 with spaces.txt")},
		{"relative_f", "file.txt"},
		{"relative_i", "index.txt"},
		{"relative_l", "list.txt"},
		{"relative_e", "example.txt"},
		{"relative_control", "candidate.txt"},
		{"relative_dot", "./file.txt"},
		{"relative_spaces_unicode", "词表 with spaces.txt"},
		{"file_uri", fileURI},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			conf := &Config{InputFile: map[string]bool{"words": true}}
			var words []string
			if !conf.ParseFile("words", tc.input, &words) {
				t.Fatal("file parameter was not handled")
			}
			if want := []string{"alpha", "beta"}; !reflect.DeepEqual(words, want) {
				t.Fatalf("ParseFile(%q) = %q, want %q", tc.input, words, want)
			}
		})
	}
}
