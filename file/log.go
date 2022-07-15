package file

import (
	"Goverlay/api"
	"Goverlay/global"
	"os"
	"regexp"
	"strings"
	"time"
)

func InitLogReading() {
	oldStat, _ := os.Stat(Path)
	oldSize := oldStat.Size()
	b, _ := os.ReadFile(Path)
	oldLen := len(strings.Split(string(b), "\r\n"))
	for {
		newStat, _ := os.Stat(Path)
		newSize := newStat.Size()
		if oldSize != newSize {
			c, _ := os.ReadFile(Path)
			begDiff := strings.Split(string(c), "\r\n")
			newLen := len(begDiff)
			for i := oldLen; i < newLen; i++ {
				line := begDiff[i-1]
				if len(line) <= 40 {
					continue
				}
				if line[33:40] != "[CHAT] " {
					continue
				}
				go parseLines(filterLine(line[40:]))
			}
			oldSize = newSize
			oldLen = newLen
		}
		time.Sleep(time.Duration(global.RefreshTime) * time.Millisecond)
	}
}

func filterLine(line string) string {
	e, _ := regexp.Compile("�[a-f0-9rl]")
	return e.ReplaceAllString(line, "")
}

func parseLines(line string) {
	if strings.HasPrefix(line, "ONLINE: ") {
		api.Nuke()
		for _, name := range strings.Split(line[8:], ", ") {
			go api.GetStats(name)
		}
	} else if strings.Contains(line, " has joined (") && strings.HasSuffix(line, ")!") || (strings.HasSuffix(line, "reconnected.") && !(strings.Contains(line, ":"))) {
		go api.GetStats(strings.Split(line, " ")[0])
	} else if strings.HasSuffix(line, "has quit!") {
		go api.RemovePlayer(strings.Split(line, " ")[0])
		go api.RemoveNicked(strings.Split(line, " ")[0])
	} else if strings.Contains(line, "joined the lobby!") || strings.Contains(line, "Sending you to ") {
		go api.Nuke()
	} else if strings.HasSuffix(line, " has quit!") || strings.HasSuffix(line, "FINAL KILL!") || strings.HasSuffix(line, "disconnected.") {
		go api.GetStats(strings.Split(line, " ")[0])
	} else if strings.HasPrefix(line, "Your new API key") {
		go SetKey(strings.Split(line, " ")[5])
	}
}
