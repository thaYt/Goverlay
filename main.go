package main

import (
	"Goverlay/api"
	"Goverlay/file"
	"Goverlay/ui"
	"time"
)

func main() {
	go ui.KeyPressEvent()
	time.Sleep(100 * time.Millisecond)
	if file.FirstRun() {
		file.InitConfig("", 50, file.FindDir())
	}
	v := file.ReadConfig()
	b := api.CheckKey()
	if v && b {
		api.ValidKey = true
	}
	go ui.Draw()
	go api.Moderate()
	file.InitLogReading()
}
