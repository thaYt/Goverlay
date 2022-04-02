package lines

import (
	"os"
	"strings"
	"thaYt/Goverlay/src/api"
	"thaYt/Goverlay/src/config"
	"thaYt/Goverlay/src/player"
	"thaYt/Goverlay/src/printer"
	"thaYt/Goverlay/src/utils"
	"time"
)

func InitLogFile() {
	homeDir, _ := os.UserHomeDir()
	file, _ := os.ReadFile(homeDir + "/.lunarclient/offline/1.8/logs/latest.log")
	linesRead := len(strings.Split(string(file), "\r\n"))
	for {
		nFile, _ := os.ReadFile(homeDir + "/.lunarclient/offline/1.8/logs/latest.log")
		nLinesRead := len(strings.Split(string(nFile), "\r\n"))
		if linesRead < nLinesRead {
			for _, p := range utils.Difference(strings.Split(string(nFile), "\r\n"), strings.Split(string(file), "\r\n")) {
				p = utils.StripLine(p)
				if strings.HasPrefix(p, "ONLINE:") {
					player.Nuke()
					for _, b := range strings.Split(p[7:], ", ") {
						go api.GetHypixelData(strings.TrimPrefix(b, " "), 0)
					}
				} else if (strings.Contains(p, "has joined ") && (strings.HasSuffix(p, "/8)!") || strings.HasSuffix(p, "/12)!") || strings.HasSuffix(p, "/16)!"))) || (strings.HasSuffix(p, "reconnected.") && !(strings.Contains(p, ":"))) {
					go api.GetHypixelData(strings.Split(p, " ")[0], 0)
				} else if strings.HasPrefix(p, "Your new API key") {
					go printer.SetStatus("Got new API key!")
					go config.SetKey(strings.Split(p, " ")[5])
					utils.NeedRefresh = true
				} else if strings.HasSuffix(p, " has quit!") || strings.HasSuffix(p, "FINAL KILL!") || strings.HasSuffix(p, "disconnected.") {
					go api.GetHypixelData(strings.Split(p, " ")[0], 1)
				} else if strings.Contains(p, "joined the lobby!") || strings.Contains(p, "Sending you to ") {
					go player.Nuke()
				}
			}
			linesRead = nLinesRead
			file = nFile
		}
		time.Sleep(time.Duration(utils.UpdateSpeed) * time.Millisecond)
	}
}
