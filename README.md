<img src="./images/header.png">

# COOK
A highly customizable custom-wordlist generator.
- #### Highly Customizable using cook.yaml

- #### Parameter & Value approach
  - Name your own parameter `-[anything] admin`
  - Input multiple values like `-p1 admin,root,su`
  - Use extension from **pre-defined dictionary**

- #### Pre-defined Extentions Categories
  - Use `archive` for `.rar, .7z, .zip, .tar,  .tgz, ...`  
  - Use `web` for `.html, .php, .aspx, .js, .jsx, .jsp, ...`
  - Many More...
  - Create your own category in **cook.yaml**

- #### Smart file detection
  - Set `file.txt` as param’s value
  - Regex input from `file.txt`:**^apps.***
  - File not found means use filename as value

# Installation
```
go get github.com/giteshnxtlvl/cook
```

## cook.yaml
This file contains character sets, words's set and extensions set specified.
```yaml
charSet:  # Each character is a separate value
    n     : [0123456789]
    A     : [ABCDEFGHIJKLMNOPQRSTUVWXYZ]
    a     : [abcdefghijklmnopqrstuvwxyz]
    aAn   : [abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789]
    An    : [ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789]
    an    : [abcdefghijklmnopqrstuvwxyz0123456789]
    aA    : [abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ]
    s     : ["!#$%&'()*+,-./:;<=>?@[\\]^_`{|}~&\""]
    all   : ["!#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~\""]

words:    # Define your own words and values
    admin: [admin, root, su]
    files: [masters, files, password]

extensions: # Extensions
    archive: [.7z, .a, .apk, .xapk, .ar, .bz2, .cab, .cpio, .deb, .dmg, .egg, .gz, .iso, .jar, .lha, .mar, .pea, .rar, .rpm, .s7z, .shar, .tar, .tbz2, .tgz, .tlz, .war, .whl, .xpi, .zip, .zipx, .xz, .pak]
    config : [.conf, .config]
    sheet  : [.ods, .xls, .xlsx, .csv, .ics .vcf]
    exec   : [.exe, .msi, .bin, .command, .sh, .bat, .crx]
    code   : [.c, .cc, .class, .clj, .cpp, .cs, .cxx, .el, .go, .h, .java, .lua, .m, .m4, .php, .php3, .php5, .php7, .pl, .po, .py, .rb, .rs, .sh, .swift, .vb, .vcxproj, .xcodeproj, .xml, .diff, .patch, .js, .jsx]
    web    : [.html, .html5, .htm, .css, .js, .jsx, .less, .scss, .wasm, .php, .php3, .php5, .php7]
    backup : [.bak, .backup, .backup1, .backup2]
    slide  : [.ppt, .odp]
    font   : [.eot, .otf, .ttf, .woff, .woff2]
    text   : [.doc, .docx, .ebook, .log, .md, .msg, .odt, .org, .pages, .pdf, .rtf, .rst, .tex, .txt, .wpd, .wps]
    audio  : [.aac, .aiff, .ape, .au, .flac, .gsm, .it, .m3u, .m4a, .mid, .mod, .mp3, .mpa, .pls, .ra, .s3m, .sid, .wav, .wma, .xm]
    book   : [.mobi, .epub, .azw1, .azw3, .azw4, .azw6, .azw, .cbr, .cbz]
    video  : [.3g2, .3gp, .aaf, .asf, .avchd, .avi, .drc, .flv, .m2v, .m4p, .m4v, .mkv, .mng, .mov, .mp2, .mp4, .mpe, .mpeg, .mpg, .mpv, .mxf, .nsv, .ogg, .ogv, .ogm, .qt, .rm, .rmvb, .roq, .srt, .svi, .vob, .webm, .wmv, .yuv]
    image  : [.3dm, .3ds, .max, .bmp, .dds, .gif, .jpg, .jpeg, .png, .psd, .xcf, .tga, .thm, .tif, .tiff, .yuv, .ai, .eps, .ps, .svg, .dwg, .dxf, .gpx, .kml, .kmz, .webp]
```
## Modifying cook.yaml
> Note: You can use above pre-defined sets without modifying anything

Steps to modify cook.yaml 
1. Create an environment varirable names `COOK` 
2. Sets it's value to file's path, doesn't matter file exists or not
   Example: COOK: E:\tools\config\cook.yaml
3. Done, run the tool and it will create `cook.yaml`.

# Usage

### Param Rules
- Always start with dash
- Can overwrite pre-defined wordlist's name.


### Basic Usage
```
  cook -start admin,root  -sep _,-  -end secret,critical  start:sep:end
```
```
  admin_secret
  admin_critical
  admin-secret
  admin-critical
  root_secret
  root_critical
  root-secret
  root-critical
```

### Extention Category Usage
```
cook -start admin,root  -sep _ -end secret  start:sep:archive
```
```
admin_.7z
admin_.a
admin_.apk
admin_.xapk
admin_.ar
...
...
```

### Regex Input from File
```
cook -start admin -exp raft-large-extensions.txt:\.asp.*  /:start:exp
```
```
/admin.aspx
/admin.asp
/admin.aspx.cs
/admin.aspx.vb
/admin.asp.asp
...
...
```
