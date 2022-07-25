package ui

import (
	"Goverlay/api"
	"Goverlay/file"
	"Goverlay/global"
	"github.com/gookit/color"
	"github.com/inancgumus/screen"
	termbox "github.com/julienroland/keyboard-termbox"
	term "github.com/nsf/termbox-go"
	"github.com/sqweek/dialog"
	"os"
)

func KeyPressEvent() { // github.com/julienroland/keyboard-termbox
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()
	var e bool
	kb := termbox.New()

	kb.Bind(func() {
		go func() {
			filename, err := dialog.File().Filter("Log File", "log").Load()
			if err != nil {
				panic(err)
			}
			file.SetDir(filename)
			global.Refresh()
		}()
	}, "l")

	kb.Bind(func() {
		e = true
		screen.Clear()
		screen.MoveTopLeft()
		color.Println()
		os.Exit(0)
	}, "q")

	kb.Bind(func() {
		global.Refresh()
	}, "u")

	kb.Bind(func() {
		global.Refresh()
	}, "x")

	kb.Bind(func() {
		api.ClearCache()
		global.Refresh()
	}, "c")

	for !e {
		kb.Poll(term.PollEvent())
	}
}
