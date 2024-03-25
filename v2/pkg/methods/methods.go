package methods

import (
	"net/url"

	"github.com/ffuf/pencode/pkg/pencode"
)

func (m *Methods) SetupMethodFunc() {
	m.MethodFuncs = map[string]func([]string, string, *[]string){
		"u":          m.Upper,
		"upper":      m.Upper,
		"l":          m.Lower,
		"lower":      m.Lower,
		"t":          m.Title,
		"title":      m.Title,
		"fb":         m.FileBase,
		"filebase":   m.FileBase,
		"leet":       m.Leet,
		"json":       m.GetJsonField,
		"smart":      m.SmartWords,
		"smartjoin":  m.SmartWordsJoin,
		"regex":      m.Regex,
		"sort":       m.Sort,
		"sortu":      m.SortUnique,
		"split":      m.Split,
		"replace":    m.Replace,
		"splitindex": m.SplitIndex,
		"reverse":    m.Reverse,
		"c":          m.Charcode,
		"charcode":   m.Charcode,
	}
}

func (m *Methods) SetupUrlFunc() {
	m.UrlFuncs = map[string]func(*url.URL, *[]string){
		"k":         m.UrlKey,
		"keys":      m.UrlKey,
		"sub":       m.UrlSub,
		"subdomain": m.UrlSub,
		"allsub":    m.UrlAllSub,
		"tld":       m.UrlTld,
		"user":      m.UrlUser,
		"pass":      m.UrlPass,
		"h":         m.UrlHost,
		"host":      m.UrlHost,
		"p":         m.UrlPort,
		"port":      m.UrlPort,
		"ph":        m.UrlPath,
		"path":      m.UrlPath,
		"f":         m.UrlFrag,
		"fragment":  m.UrlFrag,
		"q":         m.UrlRawQuery,
		"query":     m.UrlRawQuery,
		"v":         m.UrlValue,
		"value":     m.UrlValue,
		"s":         m.UrlScheme,
		"scheme":    m.UrlScheme,
		"d":         m.UrlDomain,
		"domain":    m.UrlDomain,
		"alldir":    m.UrlAllDir,
	}
}

func (m *Methods) SetupEncodersFunc() {
	m.EncodersFuncs = map[string]pencode.Encoder{
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
}
