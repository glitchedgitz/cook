# Using COOK with other tools

### Direct fuzzing with [GoBuster](https://github.com/OJ/gobuster)
```
 cook admin,root:_:archive | gobuster dir -u https://example.com/ -w -
```

### Direct fuzzing with [ffuf](https://github.com/ffuf/ffuf)
```
 cook admin,root:_:archive | ffuf -u https://example.com/FUZZ -w -
```

# Tips and Tricks
[Logical Fuzzing CDNs](https://twitter.com/krizzsk/status/1377666014347980801) by [Joel Verghese @krizzsk](https://twitter.com/krizzsk) 

# All Wordlist Repo
https://github.com/HacktivistRO/Bug-Bounty-Wordlists

# Finding IIS Longnames using Shortnames
After using IIS Short name scanner.
- Put your all guessing words in one param `words`
- Using another param `ex` for extensions and using regex `\.asp.*` for all extentions those are starting with `.asp`
- Using separaters `,-,_` becuase filename styling can be anything.
    > Note I have added extra comma in starting so command will use nothing to join previous and next columns
- Directly pipe the input in the FFUF
```
cook -words access,convert,convertor,converting,employee,employees,encrypt,encryption,encrypted,engine,engineinstall,export,exporter,exportor,failure,install,installation,location,located,locating,market,marketting,markets,orderfor,orderforward,sysupdates,sysupdater,sysupdator,sysupdate,sysupd,tanslator,translation,translating,usagereport -ex raft_ext:\.asp.* words:,-,_:ex | ffuf -u https://vulnerable.com/FUZZ -w -
```