package ui

import (
	"Goverlay/api"
	"Goverlay/file"
	"Goverlay/global"
	"github.com/gookit/color"
	"github.com/inancgumus/screen"
	"strconv"
	"strings"
	"time"
)

var (
	clines    int
	pw        int
	ph        int
	headerLen int
	header    string
	Selected  int
	Layout    = []string{
		"-lv",
		"-n",
		"-fkdr",
		"-blr",
		"-wlr",
		"-f",
		"-w",
		"-b",
	}
)

func Draw() {
	getHeaderData()
	for {
		clines = 0
		w, h := screen.Size()
		if pw != w || ph != h {
			global.NeedRefresh = true
		}
		pw, ph = w, h
		if global.NeedRefresh {
			screen.Clear()
			screen.MoveTopLeft()
			title := "┏ " + color.Bold.Render("Goverlay v2") + " ━ " + color.Green.Render(strconv.Itoa(len(api.CacheList))+" cached ")
			if api.ValidKey {
				title += "━ Key status: " + color.Bold.Render(color.Green.Render("VALID "))
			} else {
				title += "━ Key status: " + color.Bold.Render(color.Red.Render(api.KDR)) + " "
			}
			if w-headerLen > 0 {
				color.White.Println(title + strings.Repeat("━", headerLen-len(strings.Split(title, ""))+32) + "┳" + strings.Repeat("━", w-headerLen))
			} else if w-headerLen == 0 {
				color.White.Println(title + strings.Repeat("━", headerLen-len(strings.Split(title, ""))+32) + "┓")
			} else {
				color.Red.Println("Please increase the width of the terminal.")
				time.Sleep(time.Duration(global.RefreshTime) * time.Millisecond)
				continue
			}
			color.White.Println(header)
			clines++
			api.SortNicked()
			api.SortPlayers()
			for _, r := range api.LoadingList {
				if clines <= ph-3 {
					setLoading(r)
					clines++
				}
			}
			for _, r := range api.NickedList {
				if clines <= ph-3 {
					setNicked(r)
					clines++
				}
			}
			for _, r := range api.PlayerList {
				if clines <= ph-3 {
					setPlayer(r)
					clines++
				}
			}
			if h-clines-3 >= 0 {
				color.White.Print(strings.Repeat("┃\n", h-clines-3))
			} // filler lines
			setFooter()
			global.NeedRefresh = false
		}
		time.Sleep(time.Millisecond)
	}
}

func getHeaderData() {
	h := "┣ "
	for _, flag := range Layout {
		switch flag {
		case "-lv":
			h += "LEVEL ━ " // 8
		case "-n":
			h += "NAME ━━━━━━━━━━━━ " // 18
		case "-fkdr":
			h += "FKDR ━ " // 7
		case "-f":
			h += "FINALS ━ " // 9
		case "-fd":
			h += "FDEATHS ━ " // 10
		case "-wlr":
			h += "WLR ━━ " // 7
		case "-w":
			h += "WINS ━ " // 7
		case "-l":
			h += "LOSSES ━ " // 9
		case "-blr":
			h += "BBLR ━ " // 7
		case "-b":
			h += "BEDS ━ " // 7
		case "-bl":
			h += "BLOST ━ " // 8
		case "-i":
			h += "INDEX ━ " // 8
		}
	}

	headerLen = len(strings.Split(h, ""))
	header = h[:len(h)-5] + " ┛"
}

func setFooter() {
	footer := color.White.Text("┗") + getType() + strings.Repeat("━", pw-40-len(getType())) + " "
	footer += color.Gray.Text(color.OpUnderscore.Render("l")) + "ogfile ━ "
	footer += color.Gray.Text(color.OpUnderscore.Render("c")) + "lear cache ━ "
	footer += color.Gray.Text(color.OpUnderscore.Render("u")) + "pdate ━ "
	footer += color.Gray.Text(color.OpUnderscore.Render("q")) + "uit"
	color.Println(footer)
}

func getType() string {
	if strings.Contains(file.Path, "lunarclient") {
		return " using Lunar "
	} else if strings.Contains(file.Path, ".minecraft") {
		if strings.Contains(file.Path, "blclient") {
			return " using Badlion "
		}
		return " using MCLauncher "
	} else {
		return " using ? "
	}
}

func setPlayer(r api.Player) {
	color.Printf("┣ ")
	for _, flag := range Layout {
		switch flag {
		case "-lv":
			GetLevelColor(r.Level).Printf("%-8s", strconv.Itoa(r.Level))
		case "-n":
			GetRankColor(r.Rank).Printf("%-18s", r.Name)
		case "-fkdr":
			GetFBColor(r.FKDR).Printf("%-7s", strconv.FormatFloat(r.FKDR, 'f', -1, 64))
		case "-f":
			GetFKColor(r.Finals).Printf("%-9s", strconv.Itoa(r.Finals))
		case "-fd":
			Gray.Printf("%-10s", strconv.Itoa(r.FDeaths))
		case "-wlr":
			GetWLRColor(r.WLR).Printf("%-7s", strconv.FormatFloat(r.WLR, 'f', -1, 64))
		case "-w":
			GetWColor(r.Wins).Printf("%-7s", strconv.Itoa(r.Wins))
		case "-l":
			Gray.Printf("%-9s", strconv.Itoa(r.Losses))
		case "-blr":
			GetWLRColor(r.BBLR).Printf("%-7s", strconv.FormatFloat(r.BBLR, 'f', -1, 64))
		case "-b":
			GetFKColor(r.Beds).Printf("%-7s", strconv.Itoa(r.Beds))
		case "-bl":
			Gray.Printf("%-8s", strconv.Itoa(r.Bl))
		}
	}
	color.Println()
}

func setNicked(r api.Player) {
	color.Printf("┣ ")
	for _, flag := range Layout {
		switch flag {
		case "-lv":
			Red.Printf("%-8s", "-")
		case "-n":
			Red.Printf("%-18s", r.Name)
		case "-fkdr":
			Red.Printf("%-7s", "-")
		case "-f":
			Red.Printf("%-9s", "-")
		case "-fd":
			Red.Printf("%-10s", "-")
		case "-wlr":
			Red.Printf("%-7s", "-")
		case "-w":
			Red.Printf("%-7s", "-")
		case "-l":
			Red.Printf("%-9s", "-")
		case "-blr":
			Red.Printf("%-7s", "-")
		case "-b":
			Red.Printf("%-7s", "-")
		case "-bl":
			Red.Printf("%-8s", "-")
		}
	}
	color.Println()
}

func setLoading(r api.Player) {
	color.Printf("┣ ")
	for _, flag := range Layout {
		switch flag {
		case "-lv":
			Blue.Printf("%-8s", "-")
		case "-n":
			Blue.Printf("%-18s", r.Name)
		case "-fkdr":
			Blue.Printf("%-7s", "-")
		case "-f":
			Blue.Printf("%-9s", "-")
		case "-fd":
			Blue.Printf("%-10s", "-")
		case "-wlr":
			Blue.Printf("%-7s", "-")
		case "-w":
			Blue.Printf("%-7s", "-")
		case "-l":
			Blue.Printf("%-9s", "-")
		case "-blr":
			Blue.Printf("%-7s", "-")
		case "-b":
			Blue.Printf("%-7s", "-")
		case "-bl":
			Blue.Printf("%-8s", "-")
		}
	}
	color.Println()
}
