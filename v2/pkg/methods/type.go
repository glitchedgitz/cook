package methods

import (
	"net/url"

	"github.com/ffuf/pencode/pkg/pencode"
)

type Methods struct {
	LeetValues    map[string][]string
	MethodFuncs   map[string]func([]string, string, *[]string)
	UrlFuncs      map[string]func(*url.URL, *[]string)
	EncodersFuncs map[string]pencode.Encoder
}
