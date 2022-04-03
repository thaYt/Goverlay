package utils

import (
	"errors"
	"os"
	"regexp"
	"strings"
	"syscall"
	"unsafe"

	"github.com/Jeffail/gabs/v2"
)

var (
	NeedRefresh   bool
	Key           string
	UpdateSpeed   int
	Layout        []string
	Players       []Player
	NickedPlayers []Player
	Version       string
)

type Player struct {
	Nicked   bool
	Rank     string
	Level    int
	Name     string
	FKDR     float64
	WLR      float64
	BBLR     float64
	Finals   int
	FDeaths  int
	Beds     int
	Wins     int
	Losses   int
	BedsLost int
	Index    int
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func StripLine(e string) string {
	if !(strings.HasPrefix(e, "[") && strings.Contains(e, " [Client thread/INFO]: [CHAT] ")) {
		return ""
	} // making sure it doesnt crash
	if !strings.Contains((e)[:40], " [Client thread/INFO]: [CHAT] ") {
		return ""
	}
	return regexp.MustCompile(` [ [](x\d+])+`).ReplaceAllString(e, "")[40:]
}

func Difference(a, b []string) []string { // Credit: https://stackoverflow.com/questions/19374219/how-to-find-the-difference-between-two-slices-of-strings
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func GetRank(json *gabs.Container) string {
	if json.Search("player").String() != "null" {
		e := json.Search("player", "rank").String()
		e = strings.ReplaceAll(e, "\"", "")
		if e != "null" {
			if e == "ADMIN" {
				return "ADMIN"
			} else if e == "GAME_MASTER" {
				return "GM"
			} else if e == "YOUTUBER" {
				return "YOUTUBE"
			}
		} else {
			e = json.Search("player", "monthlyPackageRank").String()
			e = strings.ReplaceAll(e, "\"", "")
			if e != "null" {
				if e == "NONE" {
					return "MVP+"
				} else if e == "SUPERSTAR" {
					return "MVP++"
				}
			} else {
				e = json.Search("player", "newPackageRank").String()
				e = strings.ReplaceAll(e, "\"", "")
				if e != "null" {
					if e == "null" {
						return "NON"
					} else if e == "VIP" {
						return "VIP"
					} else if e == "VIP_PLUS" {
						return "VIP+"
					} else if e == "MVP" {
						return "MVP"
					} else if e == "MVP_PLUS" {
						return "MVP+"
					}
				} else {
					return "NON"
				}
			}

		}
	}
	return "bruh"
}

func SetTitle(title string) (int, error) { // From https://github.com/lxi1400/GoTitle
	handle, err := syscall.LoadLibrary("Kernel32.dll")
	if err != nil {
		return 0, err
	}
	defer syscall.FreeLibrary(handle)
	proc, err := syscall.GetProcAddress(handle, "SetConsoleTitleW")
	if err != nil {
		return 0, err
	}
	e, _ := syscall.UTF16PtrFromString(title)
	r, _, err := syscall.Syscall(proc, 1, uintptr(unsafe.Pointer(e)), 0, 0)
	return int(r), err
}
