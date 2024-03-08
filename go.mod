module github.com/tpfeiffer67/console

go 1.21

require (
	github.com/eiannone/keyboard v0.0.0-20220611211555-0d226195f203
	github.com/gdamore/tcell/v2 v2.7.0
	github.com/lucasb-eyer/go-colorful v1.2.0
	github.com/nathan-fiscaletti/consolesize-go v0.0.0-20220204101620-317176b6684d
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
)

replace (
	github.com/tpfeiffer67/console/render => ./render
	github.com/tpfeiffer67/console/tcellterm => ./tcellterm
	github.com/tpfeiffer67/console/terminal => ./terminal
	github.com/tpfeiffer67/console/ui/engine => ./ui/engine
	github.com/tpfeiffer67/console/ui/message => ./ui/message
	github.com/tpfeiffer67/console/ui/ntt => ./ui/ntt
)

require (
	github.com/gdamore/encoding v1.0.0 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	golang.org/x/sys v0.16.0 // indirect
	golang.org/x/term v0.16.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
