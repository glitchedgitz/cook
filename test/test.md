## Test cases after every update

```
  cook -start admin,root  -sep _,-  -end secret,critical  start:sep:end
  cook admin,root:_,-:secret,critical
```

```
  cook -start admin,root  -sep _ -end secret  start:sep:archive
  cook admin,root:_:archive
```

```
  cook -start admin -exp raft-large-extensions.txt:\.asp.*  /:start:exp
  cook -exp raft-large-extensions.txt:\.asp.*  /:admin:exp
```

```
  cook -start admin,root -file file_not_exists.txt start:_:file
  cook -file file_not_exists.txt admin,root:_:file
```
