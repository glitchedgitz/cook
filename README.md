<img src="./images/test2.png"> 

# What is COOK?
Next level wordlist and password generator.

# Why?
- Because creating/modifing wordlists are painful and time consuming process.
- Every target is fking uniq and we need to modify our wordlist according to it...
- And we have literaly too many wordlist out there, that managment and updating them is another problem.
- We all have custom wordlists.

# Features
<div style="display:grid;grid-template-columns: auto auto auto auto;" >
  <div>✔ <a href="#installation">Pre-defined Sets </a></div>
  <div>✔ <a href="#installation">Wordlist URL</a></div>
  <div>✔ <a href="#installation">Charsets like crunch </a></div>
  <div>✔ <a href="#installation">Ranges [69-1337] [F-k]</a></div>
  <div>✔ <a href="#installation">1337 Mode  </a></div>
  <div>✔ <a href="#installation">Assetnotes Wordlists </a></div>
  <div>✔ <a href="#installation">Seclists Wordlists </a></div>
  <div>✔ <a href="#installation">Files Regex</a></div>
  <div>✔ <a href="#installation">Update wordlists   </a></div>
  <div>✔ <a href="#installation">Clean Wordlists </a></div>
  <div>✔ <a href="#installation">Url analyser</a></div>
  <div>✔ <a href="#installation">Customizable</a></div>
</div>

# Fast Travel
- [Installation](#installation)
- [Customizing tool](#customizing-tool)
- [Basic Permutation](#basic-permutation)
- [Advance Permutation](#advance-permutation)
- [Predefined Sets](#predefined-sets)
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

# Flags

| Flags  | Usage |
| ------------- | ------------- |
|  -case  | Define Cases |
| -min  | Minimum no of columns to print  |
| -config  | Config Information *cook.yaml*  |
| -config-path  | Specify path for custom yaml file.  |
| -update-all  | Update all file's cache  |
| -h  | Help  |
| -v  | Verbose  |

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


# Community Powered
This tool is already powered by some awesome community members. Drive this power to next level by comtributing using following.

### Contributors
- [@Flangvik](https://twitter.com/Flangvik)
- [@noraj_rawsec](https://twitter.com/noraj_rawsec)
### Contribute 
- [cook.yaml](https://gist.github.com/giteshnxtlvl/55048a76a060da849ca8fefde2258da3#file-cook-yaml) is the backbone of the tool.  
- Share useful lists and patterns.
  Modify here [cook.yaml](https://gist.github.com/giteshnxtlvl/55048a76a060da849ca8fefde2258da3#file-cook-yaml)
- Share your awesome recipes.
- Share Ideas or new Feature Request.
- Check out [discussions](https://github.com/giteshnxtlvl/cook/discussions).

# Thanks to...
- [Assetnote](https://assetnote.io/) and [Seclist](https://github.com/danielmiessler/SecLists) for awesome wordlist
- All the [contributors](#contributors)

# Satisfied? 
<a href="https://www.buymeacoffee.com/giteshnxtlvl"><img width="300" src="./images/BMEC2X.png"></a>

# Todo
-  [X] Search wordlist `cook search [wordlist]` , that means wordlists needs tag or something
-  [ ] Flag `cook update [wordlist]`, so it will update those wordlists
-  [ ] Flag `cook add [name] [wordlist]`, to add new wordlist, directly from cmd. If `name` already exists, then ask to overwrite or not
-  [ ] Flag `cook delete [name]`, ask to confirm the delete
-  [ ] Get all directories from URL to create a list 
-  [ ] Updating cook.yaml from this repo and keeping user modifications
-  [ ] Add new assetnote's wordlists in their sets
-  [ ] Print url + local saved wordlist path, so user can use them
-  [ ] Specify start and stop of wordlist

# Bugs
-  [ ] Pipe input can't be used multiple times

# TODO examples

- Using cook as a encoder, decoder, and hash

cypcat
  -  [ ] Copying list in burp
  -  [ ] github dork
  -  [ ] shodan dork

runcat
  -  [ ] multithread any tool
  -  [ ] automate


# Final Words 
> *COOKING IS AN ART, ART NEEDS CREATIVITY*    

This tool will not help you find bugs but your creativity does.

