package main

import (
	"Goverlay/api"
	"Goverlay/file"
	"Goverlay/ui"
	"time"
)

func init() {
	go ui.KeyPressEvent()
	time.Sleep(100 * time.Millisecond)
	if !file.FirstRun() {
		ui.WinSetup()
	}
	v := file.ReadConfig()
	b := api.CheckKey()
	if v && b {
		api.ValidKey = true
	}
}

func main() {
	go ui.Draw()
	file.InitLogReading()
}
