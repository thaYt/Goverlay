package file

import (
	"Goverlay/api"
	"Goverlay/global"
	"os"
	"strconv"
	"strings"
)

var (
	Path      string
	config, _ = os.UserConfigDir()
	confFile  = config + "/goverlay/config"
)

func FirstRun() bool {
	_, err := os.ReadFile(confFile)
	if err != nil {
		return true
	}
	return false
}

func InitConfig(key string, time int, dir string) {
	var e []string
	e = append(e, "key="+key)
	e = append(e, "refreshTime="+strconv.Itoa(time))
	e = append(e, "dir="+dir+"\r\n")
	if os.WriteFile(confFile, []byte(strings.Join(e, "\r\n")), 0777) != nil {
		// todo failed idfk
	}
}

func ReadConfig() bool {
	e, _ := os.ReadFile(confFile)
	fuck := strings.Split(string(e), "\r\n")
	if len(fuck) < 4 {
		global.RefreshTime = 50
		// todo continue
		return false
	}
	key := fuck[0][4:]
	api.Key = key

	var err error

	global.RefreshTime, err = strconv.Atoi(fuck[1][12:])
	if err != nil {
		global.RefreshTime = 50
	}
	return true
}

func SetKey(key string) {
	api.Key = key
	// todo write key
}
