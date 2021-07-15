package core

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func TerminalSize(c []string) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	fmt.Printf("out: %v", string(out))
	if err != nil {
		log.Fatal(err)
	}
}

var FlagsHelp = `
` + Blue + `GITHUB` + Reset + `
    https://github.com/giteshnxtlvl/cook

` + Blue + `USAGE` + Reset + `
    cook [params and values] [pattern]
    cook -param1 value -file: filename -param3 value param3,admin_set param1 file,[1-300]

` + Blue + `SEARCH` + Reset + `
    cook search [word] 

` + Blue + `HELP` + Reset + `
    cook help [word] (case, encode, file, functions, patterns, usage)

` + Blue + `UPDATE` + Reset + `
    cook update [filename]
    This will update the file's cache.
        - Use "all" to update everything you have previously fetched
        - Use "cook" to update cooks-wordlists-database

` + Blue + `FLAGS` + Reset + `
    -l           : a->4, b->8, e->3 ...
                      It has two modes [0 Calm][1 Aggressive - Print every combinations]
    -a           : Append to the previous lines, instead of permutations
    -col         : Print column numbers and there values
    -c           : Define Cases, (U)ppercase, (L)owercase, (T)Titlecase, (C)Camelcase
    -conf        : Config Information
    -e           : Encode the whole output in 
    -m           : Minimum no of columns to print			  
    -v           : Verbose
    -h           : Help
`

var CaseHelp = `
` + Blue + `FOR FINAL OUTPUT` + Reset + `
    -case A for ALL 
    -case U for Uppercase
    -case L for Lowercase
    -case T for Titlecase
    -case C for Camelcase

` + Blue + `FOR PARTICULAR COLUMN, (no camel)` + Reset + `
    -case 0:U,2:T
        Col 0 - Uppercase
        Col 2 - Titlecase

` + Blue + `MULTIPLE CASES` + Reset + `
    -case 0:UT,2:A (column wise)
    -case TC (final output)
`

var EncodeHelp = `
WRAP YOUR PAYLOADS USING THIS

` + Blue + `ENCODERS` + Reset + `
    b64e,      b64encode           - Base64 encoder
    hexe,      hexencode           - Hex string encoder
    jsone,     jsonescape          - JSON escape
    urle,      urlencode           - URL encode reserved characters
               utf16               - UTF-16 encoder (Little Endian)
               utf16be             - UTF-16 encoder (Big Endian)
    xmle,      xmlescape           - XML escape
    urleall,   urlencodeall        - URL encode all characters
    unicodee,  unicodeencodeall    - Unicode escape string encode (all characters)
 
` + Blue + `DECODERS` + Reset + `
    b64d,      b64decode           - Base64 decoder
    hexd,      hexdecode           - Hex string decoder
    jsonu,     jsonunescape        - JSON unescape
    unicoded,  unicodedecode       - Unicode escape string decode
    urld,      urldecode           - URL decode
    xmlu,      xmlunescape         - XML unescape
 
` + Blue + `HASHES` + Reset + `
    md5                            - MD5 sum
    sha1                           - SHA1 checksum
    sha224                         - SHA224 checksum
    sha256                         - SHA256 checksum
    sha384                         - SHA384 checksum
    sha512                         - SHA512 checksum
`

var FuncHelp = `
` + Blue + `METHODS` + Reset + `

    leet                    - a->4, b->8, e->3 ...
    case                    - Apply cases
                              case[U] (U, L, T) or case[ULT] multiple cases

    json                    - Read JSON 
                              json[key] or json[key:subkey:sub-subkey]

    encode                  - Encode Functions 
                              encode[encoding]
                                     encode[en1:en2:en3] Apply one after other method

    wordplay                - This will split the word from naming convensions
                              redirectUri, redirect_uri, redirect-uri  ->  [redirect, uri] (outcome)

    fb   filebase           - Extract filename from path or url
    s    scheme             - Extract http, https, gohper, ws, etc. from URL
    u    user               - Extract username from url
    p    pass               - Extract password from url
    u:p  user:pass          - Extract username:password from url with colon
    h    host               - Extract host from url
    pr   port               - Extract port from url
    h:p  host:port          - Extract host:port from url with colon
    ph   path               - Extract path from url
    f    fragment           - Extract fragment from url
    q    query              - Extract whole query from url
    k    keys               - Extract keys from url
    v    values             - Extract values from url
    k:v  keys:values        - Extract keys:values from url
    d    domain             - Extract domain from url
         tld                - Extract tld from url
         alldir             - Extract all dirrectories from url's path
    sub  subdomain          - Extract subdomain from url
    `

var UsageHelp = `
` + Blue + `BASIC USAGE` + Reset + `
    $ cook -start admin,root  -sep _,-  -end secret,critical  /:start:sep:end
    $ cook /:admin,root:_,-:secret,critical

` + Blue + `FILE WITH REGEX` + Reset + `
    $ cook -s company -ext raft-large-extensions:\.asp.*  /:s:exp

` + Blue + `FUNCTIONS` + Reset + `
    Use functions such as date for different variations of values
    $ cook -name elliot -birth date(17,Sep,1994) name:birth

` + Blue + `RANGES` + Reset + `
    Use ranges like [1-100], [A-Z], [a-z] or [A-z] in pattern of command
    $ cook -name elliot -birth date(17,Sep,1994) name:birth

` + Blue + `USING CASES` + Reset + `
    Uppercase, lowercase, titlecase, camelcase or ALL
    For use case check FLGAS above
    $ cook camel:[1-10]:case -case C

` + Blue + `RAW STRINGS` + Reset + `
    Print value without any parsing/modification.
    Use to take ",", ":", "` + "`" + `" or any pre-defined sets or functions as raw strings.
    $ cook -date ` + "`" + `date(17,Sep,1994)` + "`" + ` raw:date

` + Blue + `PIPE INPUT` + Reset + `
    Use - as param value for pipe input
    $ cook -d - d:/:test

` + Blue + `USING -min` + Reset + `
    Print value without any parsing/modification
    $ cook n:n:n -min 1
`
