package main

import (
	"thaYt/Goverlay/src/config"
	"thaYt/Goverlay/src/event"
	"thaYt/Goverlay/src/lines"
	"thaYt/Goverlay/src/printer"
	"thaYt/Goverlay/src/update"
	"thaYt/Goverlay/src/utils"
)

func main() {
	utils.Version = "0.3 dev"
	update.CheckForUpdates(false)
	config.GetConfig()
	go event.KeyPressEvent()
	go lines.InitLogFile()
	_, err := utils.SetTitle("Goverlay " + utils.Version)
	if err.Error() != "The operation completed successfully." {
		printer.SetStatus("Error in setting terminal title: " + err.Error())
	}
	printer.InitDraw()
}
