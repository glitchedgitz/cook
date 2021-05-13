# Using COOK with other tools

### Direct fuzzing with [GoBuster](https://github.com/OJ/gobuster)
```
 cook admin,root:_:archive | gobuster dir -u https://example.com/ -w -
```

### Direct fuzzing with [ffuf](https://github.com/ffuf/ffuf)
```
 cook admin,root:_:archive | ffuf -u https://example.com/FUZZ -w -
```

# Worslists and Tips
| List | Description |
| --- | --- |
| [raft-large-extensions.txt](https://github.com/danielmiessler/SecLists/blob/master/Discovery/Web-Content/raft-large-extensions.txt) | List of all extensions |
| [all_tlds.txt](https://raw.githubusercontent.com/publicsuffix/list/master/public_suffix_list.dat) | List of all tlds |
| [Tip](https://twitter.com/krizzsk/status/1377666014347980801) by [Joel Verghese](https://twitter.com/krizzsk) | FUZZ CDNs - Logical Fuzzing |