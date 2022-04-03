package update

import (
	"github.com/Jeffail/gabs/v2"
	"github.com/cavaliergopher/grab/v3"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"thaYt/Goverlay/src/printer"
	"thaYt/Goverlay/src/utils"
	"time"
)

func CheckForUpdates(doUpdate bool) {
	go func() {
		err := os.Remove("./Goverlay.exe.del")
		if err != nil {
		}
	}()
	var cVer, err = strconv.ParseFloat(utils.Version[:3], 64)
	if err != nil {
		printer.SetStatus("Error parsing version number: " + err.Error())
		time.Sleep(time.Millisecond * 250)
		return
	}
	w, err := http.Get("https://api.github.com/repos/thayt/goverlay/releases/latest")
	if err != nil {
		printer.SetStatus("Could not get update info. " + err.Error())
		time.Sleep(time.Millisecond * 250)
		return
	}

	ww, err := io.ReadAll(w.Body)
	if err != nil {
		printer.SetStatus("Could not get update info. " + err.Error())
		time.Sleep(time.Millisecond * 250)
		return
	}

	v, err := gabs.ParseJSON(ww)
	if err != nil {
		printer.SetStatus("Could not get update info. " + err.Error())
		time.Sleep(time.Millisecond * 250)
		return
	}
	if !strings.Contains(v.Path("message").String(), "null") {
		printer.SetStatus("API rate limit exceeded. Cannot check for updates.")
		time.Sleep(time.Millisecond * 250)
		return
	}
	version := v.Path("tag_name").Data().(string)
	wVer, err := strconv.ParseFloat(version, 64)
	if err != nil {
		printer.SetStatus("Error parsing version number: " + strconv.FormatFloat(wVer, 'f', -1, 64) + " " + err.Error())
		time.Sleep(time.Millisecond * 250)
		return
	}
	if wVer > cVer {
		printer.SetStatus("Update found! Please update to version " + version + ".")
		if !doUpdate {
			return
		}
		printer.SetStatus("Downloading Update")
		// from the Spicetify project at github.com/spicetify/spicetify-cli, slightly modified
		e, err := os.Executable()
		if err != nil {
			printer.SetStatus("Update failed." + err.Error())
			time.Sleep(time.Millisecond * 250)
			return
		}
		var ed = e + ".del"
		err = os.Rename(e, ed)
		if err != nil {
			printer.SetStatus("Update failed. " + err.Error())
			time.Sleep(time.Millisecond * 250)
			return
		}
		pathToDownload := v.Path("assets").Index(0).Path("browser_download_url").Data().(string)
		_, err = grab.Get("./Goverlay.exe", pathToDownload)
		if err != nil {
			printer.SetStatus("Update failed. " + err.Error())
			time.Sleep(time.Millisecond * 250)
			return
		}
		printer.SetStatus("Please restart the program to complete the update.")
		return
	}
	if doUpdate {
		printer.SetStatus("No updates found.")
	}
}
