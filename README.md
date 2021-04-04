```

      ░          ░ ░      ░ ░  ░  ░
    ░        ░ ░ ░ ▒  ░ ░ ░ ▒  ░ ░░ ░
      ░  ▒     ░ ▒ ▒░   ░ ▒ ▒░ ░ ░▒ ▒░
    ░ ░▒ ▒  ░░ ▒░▒░▒░ ░ ▒░▒░▒░ ▒ ▒▒ ▓▒
    ▄████▄   ▒█████   ▒█████   ██ ▄█▀
    ▒██▀ ▀█  ▒██▒  ██▒▒██▒  ██▒ ██▄█▒
    ▒▓█    ▄ ▒██░  ██▒▒██░  ██▒▓███▄░
    ▒▓▓▄ ▄██▒▒██   ██░▒██   ██░▓██ █▄
    ▒ ▓███▀ ░░ ████▓▒░░ ████▓▒░▒██▒ █▄ V1

```
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

# Usage
### Basic Usage

```
  cook -start admin,root  -sep _,-  -end secret,critical  start:sep:end
```
OUTPUT
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

### EXTENSION CATEGORY USAGE
```
cook -start admin,root  -sep _ -end secret  start:sep:archive
```
OUTPUT
```
admin_.7z
admin_.a
admin_.apk
admin_.xapk
admin_.ar
...
...
```

### REGEX INPUT FROM FILE
```
cook -start admin -exp raft-large-extensions.txt:\.asp.*  /:start:exp
```
OUTPUT
```
/admin.aspx
/admin.asp
/admin.aspx.cs
/admin.aspx.vb
/admin.asp.asp
...
...
```
