package main

import (
	"log"
	"net/url"
	"strings"

	"github.com/ffuf/pencode/pkg/pencode"
	"github.com/giteshnxtlvl/cook/pkg/cook"
	"github.com/giteshnxtlvl/cook/pkg/methods"
	"github.com/giteshnxtlvl/cook/pkg/parse"
)

var methodFunc = map[string]func([]string, string, *[]string){
	"upper":      methods.Upper,
	"lower":      methods.Lower,
	"title":      methods.Title,
	"filebase":   methods.FileBase,
	"leet":       methods.Leet,
	"json":       methods.GetJsonField,
	"smart":      methods.SmartWords,
	"encode":     methods.Encode,
	"smartjoin":  methods.SmartWords,
	"regex":      methods.Regex,
	"split":      methods.Split,
	"sort":       methods.Sort,
	"sortu":      methods.SortUnique,
	"splitindex": methods.SplitIndex,
	"reverse":    methods.Reverse,
}

var urlFuncitons = map[string]func(*url.URL, *[]string){
	"keys":   methods.UrlKey,
	"subs":   methods.UrlSub,
	"allsub": methods.UrlAllSub,
	"tld":    methods.UrlTld,
	"user":   methods.UrlUser,
	"pass":   methods.UrlPass,
	"host":   methods.UrlHost,
	"port":   methods.UrlPort,
	"path":   methods.UrlPath,
	"frag":   methods.UrlFrag,
	"query":  methods.UrlRawQuery,
	"value":  methods.UrlValue,
	"scheme": methods.UrlScheme,
	"domain": methods.UrlDomain,
	"alldir": methods.UrlAllDir,
}

var availableEncoders = map[string]pencode.Encoder{
	"b64encode":        pencode.Base64Encoder{},
	"b64e":             pencode.Base64Encoder{},
	"b64decode":        pencode.Base64Decoder{},
	"b64d":             pencode.Base64Decoder{},
	"filename.tmpl":    pencode.Template{},
	"hexencode":        pencode.HexEncoder{},
	"hexe":             pencode.HexEncoder{},
	"hexdecode":        pencode.HexDecoder{},
	"hexd":             pencode.HexDecoder{},
	"jsonescape":       pencode.JSONEscaper{},
	"jsonunescape":     pencode.JSONUnescaper{},
	"lower":            pencode.StrToLower{},
	"md5":              pencode.MD5Hasher{},
	"sha1":             pencode.SHA1Hasher{},
	"sha224":           pencode.SHA224Hasher{},
	"sha256":           pencode.SHA256Hasher{},
	"sha384":           pencode.SHA384Hasher{},
	"sha512":           pencode.SHA512Hasher{},
	"unicodedecode":    pencode.UnicodeDecode{},
	"unicodeencodeall": pencode.UnicodeEncodeAll{},
	"upper":            pencode.StrToUpper{},
	"urlencode":        pencode.URLEncoder{},
	"urle":             pencode.URLEncoder{},
	"urld":             pencode.URLDecoder{},
	"urlencodeall":     pencode.URLEncoderAll{},
	"utf16":            pencode.UTF16LEEncode{},
	"utf16be":          pencode.UTF16BEEncode{},
	"xmlescape":        pencode.XMLEscaper{},
	"xmlunescape":      pencode.XMLUnescaper{},
}

func applyMethods(vallll []string, meths []string, array *[]string) {

	tmp := []string{}
	analyseMethods := [][]string{}

	for _, g := range meths {
		if strings.Contains(g, "[") {
			name, value := parse.ReadSqBr(g)
			analyseMethods = append(analyseMethods, []string{name, value})
		} else {
			analyseMethods = append(analyseMethods, []string{g, ""})
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
			log.Fatalf("Func \"%s\" Doesn't exists", f)
		}
		vallll = tmp
		tmp = nil
	}

	*array = append(*array, vallll...)
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
