package config

import (
	"os"
	"strconv"
	"strings"
	"thaYt/Goverlay/src/printer"
	"thaYt/Goverlay/src/utils"
)

func GetConfig() {
	configDir, _ := os.UserConfigDir()
	data := []byte(`key: 

# -lv = level, -n = name, -fkdr = fkdr, -f = finals, -fd = final deaths, -wlr = wlr, -w = wins, -l = losses, -blr = bblr, -b = beds, -bl = beds lost, -i = index
layout: -lv -n -fkdr -wlr -blr -f -i

# The update flag marks how fast the program updates in milliseconds.
update: 250`)
	if utils.Exists(configDir + "/goverlay") {
		if utils.Exists(configDir + "/goverlay/settings") {
			e, err := os.ReadFile(configDir + "/goverlay/settings")
			if err != nil {
				utils.Key = ""
				printer.SetStatus("There was an error reading the config file. Please report this bug!")
			}
			d := strings.Split(string(e), "\n")
			for _, f := range d {
				if strings.HasPrefix(f, "key:") {
					key := strings.Split(f, " ")
					if len(key) == 1 {
						utils.Key = ""
						printer.SetStatus("No key entered.")
					} else {
						utils.Key = key[1]
					}
				} else if strings.HasPrefix(f, "layout:") {
					mods := strings.Split(f, " ")
					println(len(mods))
					if len(mods) == 1 {
						utils.Layout = []string{
							"-lv",
							"-n",
						}
						printer.SetStatus("No layout entered.")
					} else {
						utils.Layout = mods[1:]
					}
				} else if strings.HasPrefix(f, "update:") {
					b := strings.Split(f, " ")
					if len(b) == 1 {
						utils.UpdateSpeed = 250
						printer.SetStatus("No update time entered.")
					} else {
						speed, err := strconv.Atoi(b[1])
						if err == nil {
							utils.UpdateSpeed = speed
						} else {
							printer.SetStatus("Update times formatted incorrectly.")
						}
					}
				}
			}
		} else {
			os.Create(configDir + "/goverlay/settings")
			os.WriteFile(configDir+"/goverlay/settings", data, os.ModeExclusive)
			GetConfig()
		}
	} else {
		os.Mkdir(configDir+"/goverlay", os.ModeExclusive)
		os.Create(configDir + "/goverlay/settings")
		os.WriteFile(configDir+"/goverlay/settings", data, os.ModeExclusive)
		GetConfig()
	}
}

func SetKey(key string) {
	configDir, _ := os.UserConfigDir()
	if utils.Exists(configDir + "/goverlay") {
		if !utils.Exists(configDir + "/goverlay/settings") {
			utils.Key = key
			return
		}
		e, err := os.ReadFile(configDir + "/goverlay/settings")
		if err != nil {
			utils.Key = key
			return
		}
		data := "key: " + key + "\n" + strings.Join(strings.Split(string(e), "\n")[1:], "\n")
		os.WriteFile(configDir+"/goverlay/settings", []byte(data), os.ModeExclusive)
	}
	utils.Key = key
}
