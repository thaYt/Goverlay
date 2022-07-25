package file

import (
	"Goverlay/api"
	"Goverlay/global"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/gookit/color"
)

var (
	Path      string
	config, _ = os.UserConfigDir()
	confFile  = config + "/goverlay/config"
	home, _   = os.UserHomeDir()
)

func FirstRun() bool {
	_, err := os.ReadFile(confFile)
	return err != nil
}

func InitConfig(key string, time int, dir string) {
	var e []string
	e = append(e, "key="+key)
	e = append(e, "refreshTime="+strconv.Itoa(time))
	e = append(e, "dir="+dir)
	e = append(e, "customdir="+"\r\n")
	os.Mkdir(config+"/goverlay", 0777)
	if os.WriteFile(confFile, []byte(strings.Join(e, "\r\n")), 0777) != nil {
		// idk man
	}
}

func ReadConfig() bool {
	e, _ := os.ReadFile(confFile)
	log := strings.Split(string(e), "\r\n")
	if len(log) < 4 {
		global.RefreshTime = 50
		// todo continue
		return false
	}
	key := log[0][4:]
	api.Key = key

	var err error

	global.RefreshTime, err = strconv.Atoi(log[1][12:])
	if err != nil {
		global.RefreshTime = 50
	}

	if len(log[3]) > 10 {
		Path = log[3][10:]
	} else {
		Path = log[2][4:]
	}
	return true
}

func SetKey(key string) {
	var a []string
	api.Key = key
	e, _ := os.ReadFile(confFile)
	log := strings.Split(string(e), "\r\n")
	for b, c := range log {
		if b == 0 {
			a = append(a, "key="+key)
			continue
		}
		if c == "\r\n" || c == "\n" || b > 4 {
			continue
		}
		a = append(a, c)
	}
	os.WriteFile(confFile, []byte(strings.Join(a, "\r\n")), 0777)
	api.CheckKey()
}

func SetDir(dir string) {
	var a []string
	Reset, Path = true, dir
	e, _ := os.ReadFile(confFile)
	log := strings.Split(string(e), "\r\n")
	for b, c := range log {
		if b == 2 {
			a = append(a, "dir="+dir)
			continue
		}
		if c == "\r\n" || c == "\n" || b > 4 {
			continue
		}
		a = append(a, c)
	}
	os.WriteFile(confFile, []byte(strings.Join(a, "\r\n")), 0777)
}

type dir struct {
	time     int64
	filename string
}

// FindDir is just used to find the most recent directory at first start.
func FindDir() string {
	switch runtime.GOOS {
	case "windows":
		var dirs []dir
		mcLogfile, err := os.Stat(home + "\\AppData\\Roaming\\.minecraft\\logs\\latest.log")
		if err == nil {
			dirs = append(dirs, dir{
				mcLogfile.ModTime().UnixMilli(),
				home + "\\AppData\\Roaming\\.minecraft\\logs\\" + mcLogfile.Name(),
			})
		}

		lcLogDir, err := os.ReadDir(home + "\\.lunarclient\\offline\\")
		if err == nil {
			for _, subdir := range lcLogDir {
				if subdir.IsDir() {
					lcLogfile, err := os.Stat(home + "\\.lunarclient\\offline\\" + subdir.Name() + "\\logs\\latest.log")
					if err == nil {
						dirs = append(dirs, dir{
							lcLogfile.ModTime().UnixMilli(),
							home + "\\.lunarclient\\offline\\" + subdir.Name() + "\\logs\\" + lcLogfile.Name(),
						})
					}
				}
			}
		}

		blcLogfile, err := os.Stat(home + "\\AppData\\Roaming\\.minecraft\\logs\\blclient\\minecraft\\latest.log")
		if err == nil {
			dirs = append(dirs, dir{
				blcLogfile.ModTime().UnixMilli(),
				home + "\\AppData\\Roaming\\.minecraft\\logs\\blclient\\minecraft\\" + blcLogfile.Name(),
			})
		}
		sort.Slice(dirs, func(i, j int) bool {
			return dirs[i].time > dirs[j].time
		})
		return dirs[0].filename
	case "darwin":
		// might do eventually, not in to/do
	case "linux":
		var dirs []dir
		mcLogfile, err := os.Stat(home + "/.minecraft/logs/latest.log")
		if err == nil {
			dirs = append(dirs, dir{
				mcLogfile.ModTime().UnixMilli(),
				home + "/.minecraft/logs/" + mcLogfile.Name(),
			})
		}

		lcLogDir, err := os.ReadDir(home + "/.lunarclient/offline/")
		if err == nil {
			for _, subdir := range lcLogDir {
				if subdir.IsDir() {
					lcLogfile, err := os.Stat(home + "/.lunarclient/offline/" + subdir.Name() + "/logs/latest.log")
					if err == nil {
						dirs = append(dirs, dir{
							lcLogfile.ModTime().UnixMilli(),
							home + "/.lunarclient/offline/" + subdir.Name() + "/logs/" + lcLogfile.Name(),
						})
					}
				}
			}
		}

		blcLogfile, err := os.Stat(home + "/.minecraft/logs/blclient/minecraft/latest.log")
		if err == nil {
			dirs = append(dirs, dir{
				blcLogfile.ModTime().UnixMilli(),
				home + "/.minecraft/logs/blclient/minecraft/" + blcLogfile.Name(),
			})
		}
		sort.Slice(dirs, func(i, j int) bool {
			return dirs[i].time > dirs[j].time
		})
		return dirs[0].filename
	}
	color.Red.Println("Unsupported OS.")
	os.Exit(1)
	return "awojevoaiwejvipoawievoapwimevpowaievpoaivepowe "
}
