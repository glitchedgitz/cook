package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"github.com/ffuf/pencode/pkg/pencode"
	"github.com/giteshnxtlvl/cook/pkg/cook"
	"github.com/giteshnxtlvl/cook/pkg/methods"
	"github.com/giteshnxtlvl/cook/pkg/parse"
)

var methodFunc = map[string]func([]string, string, *[]string){
	"u":          methods.Upper,
	"upper":      methods.Upper,
	"l":          methods.Lower,
	"lower":      methods.Lower,
	"t":          methods.Title,
	"title":      methods.Title,
	"fb":         methods.FileBase,
	"filebase":   methods.FileBase,
	"leet":       methods.Leet,
	"json":       methods.GetJsonField,
	"smart":      methods.SmartWords,
	"smartjoin":  methods.SmartWordsJoin,
	"regex":      methods.Regex,
	"split":      methods.Split,
	"sort":       methods.Sort,
	"sortu":      methods.SortUnique,
	"splitindex": methods.SplitIndex,
	"reverse":    methods.Reverse,
	"c":          methods.Charcode,
	"charcode":   methods.Charcode,
}

var urlFuncitons = map[string]func(*url.URL, *[]string){
	"k":         methods.UrlKey,
	"keys":      methods.UrlKey,
	"sub":       methods.UrlSub,
	"subdomain": methods.UrlSub,
	"allsub":    methods.UrlAllSub,
	"tld":       methods.UrlTld,
	"user":      methods.UrlUser,
	"pass":      methods.UrlPass,
	"h":         methods.UrlHost,
	"host":      methods.UrlHost,
	"p":         methods.UrlPort,
	"port":      methods.UrlPort,
	"ph":        methods.UrlPath,
	"path":      methods.UrlPath,
	"f":         methods.UrlFrag,
	"fragment":  methods.UrlFrag,
	"q":         methods.UrlRawQuery,
	"query":     methods.UrlRawQuery,
	"v":         methods.UrlValue,
	"value":     methods.UrlValue,
	"s":         methods.UrlScheme,
	"scheme":    methods.UrlScheme,
	"d":         methods.UrlDomain,
	"domain":    methods.UrlDomain,
	"alldir":    methods.UrlAllDir,
}

var availableEncoders = map[string]pencode.Encoder{
	"b64e":             pencode.Base64Encoder{},
	"b64encode":        pencode.Base64Encoder{},
	"b64d":             pencode.Base64Decoder{},
	"b64decode":        pencode.Base64Decoder{},
	"filename.tmpl":    pencode.Template{},
	"hexe":             pencode.HexEncoder{},
	"hexencode":        pencode.HexEncoder{},
	"hexd":             pencode.HexDecoder{},
	"hexdecode":        pencode.HexDecoder{},
	"jsone":            pencode.JSONEscaper{},
	"jsonescape":       pencode.JSONEscaper{},
	"jsonu":            pencode.JSONUnescaper{},
	"jsonunescape":     pencode.JSONUnescaper{},
	"md5":              pencode.MD5Hasher{},
	"sha1":             pencode.SHA1Hasher{},
	"sha224":           pencode.SHA224Hasher{},
	"sha256":           pencode.SHA256Hasher{},
	"sha384":           pencode.SHA384Hasher{},
	"sha512":           pencode.SHA512Hasher{},
	"unicoded":         pencode.UnicodeDecode{},
	"unicodedecode":    pencode.UnicodeDecode{},
	"unicodee":         pencode.UnicodeEncodeAll{},
	"unicodeencodeall": pencode.UnicodeEncodeAll{},
	"urle":             pencode.URLEncoder{},
	"urlencode":        pencode.URLEncoder{},
	"urld":             pencode.URLDecoder{},
	"urldecode":        pencode.URLDecoder{},
	"urlea":            pencode.URLEncoderAll{},
	"urlencodeall":     pencode.URLEncoderAll{},
	"utf16":            pencode.UTF16LEEncode{},
	"utf16be":          pencode.UTF16BEEncode{},
	"xmle":             pencode.XMLEscaper{},
	"xmlescape":        pencode.XMLEscaper{},
	"xmlu":             pencode.XMLUnescaper{},
	"xmlunescape":      pencode.XMLUnescaper{},
}

func applyMethods(vallll []string, meths []string, array *[]string) {

	tmp := []string{}
	analyseMethods := [][]string{}

	for _, g := range meths {
		if strings.Contains(g, "[") {
			name, value := parse.ReadSqBr(g)
			analyseMethods = append(analyseMethods, []string{strings.ToLower(name), value})
		} else {
			analyseMethods = append(analyseMethods, []string{strings.ToLower(g), ""})
		}
	}

	for _, v := range analyseMethods {
		f := v[0]
		value := v[1]
		if fn, exists := methodFunc[f]; exists {
			fn(vallll, value, &tmp)
		} else if fn, exists := urlFuncitons[f]; exists {
			methods.AnalyzeURLs(vallll, fn, &tmp)
		} else if e, exists := availableEncoders[f]; exists {
			for _, v := range vallll {
				output, err := e.Encode([]byte(v))
				if err != nil {
					log.Fatalln("Err")
				}
				tmp = append(tmp, string(output))
			}
		} else {
			fmt.Fprintf(os.Stderr, "\nFunc \"%s\" Doesn't exists\n", f)
			mistypedCheck(f)
		}
		vallll = tmp
		tmp = nil
	}

	*array = append(*array, vallll...)
}

func mistypedCheck(mistyped string) {
	fmt.Fprintln(os.Stderr)

	if len(mistyped) < 3 {
		return
	}
	fmt.Fprintln(os.Stderr, "Similar Methods")

	check := func(k string) {
		similarity := strutil.Similarity(mistyped, k, metrics.NewHamming())
		if similarity >= 0.3 {
			fmt.Println("-", k)
		}
	}

	for k := range methodFunc {
		check(k)
	}

	for k := range urlFuncitons {
		check(k)
	}

	for k := range availableEncoders {
		check(k)
	}

}

func checkMethods(p string, array *[]string) bool {
	if strings.Count(p, ".") > 0 {
		splitS := splitMethods(p)
		u := splitS[0]
		if _, exists := params[u]; exists {

			vallll := []string{}

			if !checkParam(u, &vallll) && !cook.CheckYaml(u, &vallll) {
				return false
			}

			applyMethods(vallll, splitS[1:], array)

			return true
		}

	}
	return false
}
