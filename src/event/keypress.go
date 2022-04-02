package event

import (
	"github.com/inancgumus/screen"
	"github.com/julienroland/keyboard-termbox"
	term "github.com/nsf/termbox-go"
	"os"
	"thaYt/Goverlay/src/update"
)

func KeyPressEvent() { // i think this code is from the first import, github.com/julienroland/keyboard-termbox
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()
	var e bool
	kb := termbox.New()

	kb.Bind(func() {
		e = true
		screen.Clear()
		screen.MoveTopLeft()
		os.Exit(0)
	}, "q")

	kb.Bind(func() {
		update.CheckForUpdates(true)
	}, "u")

	for !e {
		kb.Poll(term.PollEvent())
	}
}
