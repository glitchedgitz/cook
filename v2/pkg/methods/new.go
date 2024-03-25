package methods

import (
	"net/url"

	"github.com/ffuf/pencode/pkg/pencode"
)

func New(LeetValues map[string][]string) *Methods {

	m := &Methods{
		LeetValues: LeetValues,
	}
	m.MethodFuncs = make(map[string]func([]string, string, *[]string))
	m.UrlFuncs = make(map[string]func(*url.URL, *[]string))
	m.EncodersFuncs = make(map[string]pencode.Encoder)

	m.SetupMethodFunc()
	m.SetupUrlFunc()
	m.SetupEncodersFunc()

	return m
}
