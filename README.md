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
## Basic Usage
```
Structure:
cook -ingredient1 <value1> -ingredient2 <value2> ... [pattern]
```
