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
	kb := termbox.New()

	kb.Bind(func() {
		go func() {
			filename, err := dialog.File().Filter("Log File", "log").Load()
			if err != nil {
				if err.Error() == "Cancelled" {
					return
				}
				panic(err)
			}
			file.SetDir(filename)
			global.Refresh()
		}()
	}, "l")

	kb.Bind(func() {
		screen.Clear()
		screen.MoveTopLeft()
		color.Println()
		os.Exit(0)
	}, "q")

	kb.Bind(func() {
		api.ClearCache()
		global.Refresh()
	}, "c")

	for {
		kb.Poll(term.PollEvent())
	}
}
