package ui

import (
	"Goverlay/api"
	"Goverlay/global"
	"github.com/gookit/color"
	"github.com/inancgumus/screen"
	termbox "github.com/julienroland/keyboard-termbox"
	term "github.com/nsf/termbox-go"
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
		SetDir()
	}, "d")

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
		api.ClearCache()
		global.Refresh()
	}, "c")

	// this is all setup stuff
	kb.Bind(func() {
		if selecting {
			if Selected != 0 {
				Selected--
			} else {
				Selected = 2
			}
			global.Refresh()
		}
	}, "up")

	kb.Bind(func() {
		if selecting {
			Selected++
			global.Refresh()
		}
	}, "down")

	kb.Bind(func() {
		if selecting || scanInt {
			Enter = true
			global.Refresh()
		}
	}, "enter")

	// I know this is horrible, and I couldn't care less
	// if someone has a better way, I'd like to see it
	kb.Bind(func() {
		if scanInt {
			str += "0"
			global.Refresh()
		}
	}, "0")
	kb.Bind(func() {
		if scanInt {
			str += "1"
			global.Refresh()
		}
	}, "1")
	kb.Bind(func() {
		if scanInt {
			str += "2"
			global.Refresh()
		}
	}, "2")
	kb.Bind(func() {
		if scanInt {
			str += "3"
			global.Refresh()
		}
	}, "3")
	kb.Bind(func() {
		if scanInt {
			str += "4"
			global.Refresh()
		}
	}, "4")
	kb.Bind(func() {
		if scanInt {
			str += "5"
			global.Refresh()
		}
	}, "5")
	kb.Bind(func() {
		if scanInt {
			str += "6"
			global.Refresh()
		}
	}, "6")
	kb.Bind(func() {
		if scanInt {
			str += "7"
			global.Refresh()
		}
	}, "7")
	kb.Bind(func() {
		if scanInt {
			str += "8"
			global.Refresh()
		}
	}, "8")
	kb.Bind(func() {
		if scanInt {
			str += "9"
			global.Refresh()
		}
	}, "9")
	kb.Bind(func() {
		if scanInt && len(str) > 0 {
			str = str[:len(str)-1]
			global.Refresh()
		}
	}, "backspace")

	for !e {
		kb.Poll(term.PollEvent())
	}
}
