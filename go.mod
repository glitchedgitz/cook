module github.com/giteshnxtlvl/cook

go 1.17

require (
	github.com/adrg/strutil v0.2.3
	github.com/buger/jsonparser v1.1.1
	github.com/ffuf/pencode v0.0.0-20210513164852-47e578431524
	github.com/manifoldco/promptui v0.9.0
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

require (
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	golang.org/x/sys v0.0.0-20211216021012-1d35b9e2eb4e // indirect
)

retract (
	v2.0.0+incompatible
	v1.6.0
)
