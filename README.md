<img src="assets/head.png">

# COOK
An overpower wordlist generator, splitter, merger, finder, saver, create words permutation and combinations, apply different encoding/decoding and everything you need.  

Frustration killer! & Customizable!

# Customizable
Cook is highly customizable and it depends on
[cook-ingredients](https://github.com/giteshnxtlvl/cook-ingredients). Cook Ingredients consists YAML Collection of word-sets, extensions, funcitons to generate pattern and wordlists.

### Installation
Use Go or download [latest builds](https://github.com/giteshnxtlvl/cook/releases/)  
```
go install github.com/giteshnxtlvl/cook/cmd/cook@latest
```

> After installation, run `cook` for one time, it will download [cook-ingredients](https://github.com/giteshnxtlvl/cook-ingredients) automatically at `%USERPROFILE%/cook-ingredients` for windows and `$home/cook-ingredients` for linux.

# Basic
Without basics, everything is useless.
<img src="assets/basic.png">

## Parametric Approach
You can define your own params and use them to generate the pattern. This will be useful once you understand [methods](#methods)

# Save wordlists and word sets
<img src="assets/savewordlist.png">

### Search Wordlist
```
cook search keyword
```

## Reading File using Cook
If you want to use a file from current working directory.  
Use `:` after param name. 
```
cook -f: live.txt f
```

# Methods
Methods will let you apply diffenent sets of operation on final output or particular column as you want. You can encode, decode, reverse, split, sort, extract different part of urls and much more...

- `-m/-method` to apply methods on the final output
- `-mc/-methodcol` to apply column-wise.
- `param.methodname` apply to any parameter-wise, will example this param thing later.
- `param.md5.b64e` apply multiple methods, this will first md5 hash the value and then base64 encode the hashed value.


<details><summary>All methods</summary>

```
METHODS
    Apply different sets of operations to your wordlists

STRING/LIST/JSON
    sort                           - Sort them
    sortu                          - Sort them with unique values only
    reverse                        - Reverse string
    split                          - split[char]
    splitindex                     - splitindex[char:index]
    replace                        - Replace All replace[this:tothis]
    leet                           - a->4, b->8, e->3 ...
                                     leet[0] or leet[1]
    json                           - Extract JSON field
                                     json[key] or json[key:subkey:sub-subkey]
    smart                          - Separate words with naming convensions
                                     redirectUri, redirect_uri, redirect-uri  ->  [redirect, uri]
    smartjoin                      - This will split the words from naming convensions &
                                     param.smartjoin[c,_] (case, join)
                                     redirect-uri, redirectUri, redirect_uri ->  redirect_Uri

    u          upper               - Uppercase
    l          lower               - Lowercase
    t          title               - Titlecase

URLS
    fb         filebase            - Extract filename from path or url
    s          scheme              - Extract http, https, gohper, ws, etc. from URL
               user                - Extract username from url
               pass                - Extract password from url
    h          host                - Extract host from url
    p          port                - Extract port from url
    ph         path                - Extract path from url
    f          fragment            - Extract fragment from url
    q          query               - Extract whole query from url
    k          keys                - Extract keys from url
    v          values              - Extract values from url
    d          domain              - Extract domain from url
               tld                 - Extract tld from url
               alldir              - Extract all dirrectories from url's path
    sub        subdomain           - Extract subdomain from url
               allsubs             - Extract subdomain from url

ENCODERS
    b64e       b64encode           - Base64 encoder
    hexe       hexencode           - Hex string encoder
               charcode            - Give charcode encoding
                                     charcode[0] without semicolon
                                     charcode[1] with semicolon
    jsone      jsonescape          - JSON escape
    urle       urlencode           - URL encode reserved characters
               utf16               - UTF-16 encoder (Little Endian)
               utf16be             - UTF-16 encoder (Big Endian)
    xmle       xmlescape           - XML escape
    urleall    urlencodeall        - URL encode all characters
    unicodee   unicodeencodeall    - Unicode escape string encode (all characters)

DECODERS
    b64d       b64decode           - Base64 decoder
    hexd       hexdecode           - Hex string decoder
    jsonu      jsonunescape        - JSON unescape
    unicoded   unicodedecode       - Unicode escape string decode
    urld       urldecode           - URL decode
    xmlu       xmlunescape         - XML unescape

HASHES
    md5                            - MD5 sum
    sha1                           - SHA1 checksum
    sha224                         - SHA224 checksum
    sha256                         - SHA256 checksum
    sha384                         - SHA384 checksum
    sha512                         - SHA512 checksum
  
```
</details>

## Multiple Methods
You can apply multiple set of operations on partiocular column or final output in one command. So you don't have to re-run the tool again and again.

# Direct fuzzing with FUFF
You can use generated output from cook directly with [ffuf](https://github.com/ffuf/ffuf) using pipe

```
cook usernames_list : passwords_list -m b64e | ffuf -u https://target.com -w - -H "Authorization: Basic FUZZ"
```

Similarly you can fuzz directories/headers/params/numeric ids... And can apply required algorithms on your payloads.

# Ranges


# Functions
```
cook -dob date[17,Sep,1994] elliot _,-, dob
```
> Customize:    
 Create your own functions in `cook-ingredients/my.yaml` under functions:
