<img src="./images/test1.png"> 

<h3 align="center">
<a href="https://twitter.com/giteshnxtlvl"><img src="./images/twitter.png"></a>
<a href="https://github.com/giteshnxtlvl/cook/discussions/new"><img src="./images/New Feature Ideas.png"></a>
<a href="https://github.com/giteshnxtlvl/cook/issues/new"><img src="./images/New ListPattern.png"></a>
<a href="https://www.buymeacoffee.com/giteshnxtlvl"><img src="./images/BMEC.png"></a>
</h3>

<h1 align="center">COOK</h1>
<h4 align="center">Next level wordlist and password generator.</h4>

# WHY?
- Because creating/modifing wordlists are painful and time consuming process.
- Every target is fking uniq and we need to modify our wordlist according to it...
- And we have literaly too many wordlist out there, that managment and updating them is another sort of burdern.

# Features
:heavy_check_mark: Wordlist URL Support
:heavy_check_mark: 1337 Mode
:heavy_check_mark: 

# Fast Travel
- [WHY?](#why)
- [Features](#features)
- [Fast Travel](#fast-travel)
- [Installation](#installation)
      - [Using Go](#using-go)
      - [Download latest builds](#download-latest-builds)
- [Customizing tool](#customizing-tool)
- [Basic Permutation](#basic-permutation)
- [Advance Permutation](#advance-permutation)
- [Predefined Sets](#predefined-sets)
    - [Create your own unique sets](#create-your-own-unique-sets)
    - [Use it like CRUNCH](#use-it-like-crunch)
- [Patterns/Functions](#patternsfunctions)
- [Ranges](#ranges)
- [Files](#files)
    - [Regex Input from File](#regex-input-from-file)
    - [Save Wordlists by Unique Names](#save-wordlists-by-unique-names)
    - [File not found](#file-not-found)
- [Cases](#cases)
    - [Minimum](#minimum)
    - [Pipe input](#pipe-input)
    - [Raw String](#raw-string)
- [Contribute](#contribute)
- [Satisfied?](#satisfied)
- [THE MAIN FILE](#the-main-file)
- [Final Words](#final-words)

# Installation
#### Using Go
Install/Update using these commands  
`go get -u github.com/giteshnxtlvl/cook`
OR
`GO111MODULE=on go get -u github.com/giteshnxtlvl/cook`

#### Download latest builds  
  https://github.com/giteshnxtlvl/cook/releases/

# Customizing tool
Tool is using a file `cook.yaml`, this file is database for cook.  
**Method 1**  
Default location in linux `$HOME/.config/cook/cook.yaml`.  
For windows it will be `%USERPROFILE%/.config/cook/cook.yaml`

**Method 2**
1. Download [cook.yaml](https://gist.githubusercontent.com/giteshnxtlvl/55048a76a060da849ca8fefde2258da3/raw/eda15049d56d37afb1bb1f8ee07daba2db1b9628/cook.yaml)
1. Create an environment variable `COOK` =`Path of file`  
3. Done, Run `cook -config` to confirm.

**Method 3**   
Use `-config-path` flag to specify location of the config file. This is useful if you want to try different config files.


# Basic Permutation

  <img src="./images/02.png">  
  
  **Recipe**
  ```
    cook -start admin,root  -sep _,-  -end secret,critical  start sep end
  ```
  ```
    cook admin,root _,- secret,critical
  ```
# Advance Permutation
Understanding concept is important!
<img src="./images/09.png">
  

# Predefined Sets
  <img src="./images/03.png">    
  
  **Recipe**
  ```
   cook -start admin,root  -sep _ -end secret  start:sep:archive
  ```
  ```
   cook admin,root:_:archive
  ```
### Create your own unique sets  
  <img src="./images/06.png">

### Use it like CRUNCH  
  <img src="./images/08.png">


# Patterns/Functions

<img src="./images/11.png"> 

**Recipe**
```
  cook -name elliot -birth date(17,Sep,1994) name:birth
```

# Ranges
<img width="640" src="./images/13.png"> 

# Files
  ### Regex Input from File  
  Use this feature to fuzz [IIS Shortnames](https://www.youtube.com/watch?v=HrJW6Y9kHC4)
  <img src="./images/07.png">    
  
  **Recipe**
  ```
   cook -exp raft-large-extensions.txt:\.asp.*  /:admin:exp
  ```
  
  ### Save Wordlists by Unique Names  
  Now you don't need to make aliases or type those huge filenames. Just one single name.
<img src="./images/05.png">

  ### File not found  
  If file mentioned in param not found, then there will be no errors, Instead it will do this.
  ```
   cook -file file_not_exists.txt admin,root:_:file
  ```
  ```
    admin_file_not_exists.txt
    root_file_not_exists.txt
  ```

# Cases
<img src="./images/12.png">


### Minimum
Use `-min <num>` to print minimum no of columns to print.  
Example this command will print 1,2,3 digit numbers
```
cook n:n:n -min 1
```

### Pipe input
Use `-` as param value for pipe input
```
cook -d - d:anything
```

### Raw String
Don't parse the value
```
cook -word `date(10,12,1999)`
```

# Community Powered
This tool is already powered by some awesome community members. Drive this power to next level by comtributing using following.

Contribute 
- [cook.yaml](https://gist.github.com/giteshnxtlvl/55048a76a060da849ca8fefde2258da3#file-cook-yaml) is the backbone of the tool.  
- Share useful lists and patterns.
  Modify here [cook.yaml](https://gist.github.com/giteshnxtlvl/55048a76a060da849ca8fefde2258da3#file-cook-yaml)
- Share your awesome recipes.
- Share Ideas or new Feature Request.
- Check out [discussions](https://github.com/giteshnxtlvl/cook/discussions).

# Satisfied? 

<a href="https://www.buymeacoffee.com/giteshnxtlvl"><img width="300" src="./images/BMEC2X.png"></a>

# TODO
Instead of permutations, appending to line by line mode

# Final Words 
> *COOKING IS AN ART, ART NEEDS CREATIVITY*    

This tool will not help you find bugs but your creativity does.

