package printer

import (
	"strconv"
	"strings"
	"thaYt/Goverlay/src/colors"
	"thaYt/Goverlay/src/utils"
	"time"

	"github.com/gookit/color"
	"github.com/inancgumus/screen"
)

var (
	clines    int
	Status    string
	pw        int
	ph        int
	toomany   bool
	headerLen int
	header    string
)

func InitDraw() {
	screen.Clear()
	screen.MoveTopLeft()
	time.Sleep(100 * time.Millisecond)
	color.White.Println("Welcome to Goverlay.")
	getHeaderData()
	time.Sleep(time.Second)
	utils.NeedRefresh = true
	Draw()
}

func Draw() {
	for {
		clines = 0
		w, h := screen.Size()
		if pw != w || ph != h {
			utils.NeedRefresh = true
		}
		pw, ph = w, h
		if utils.NeedRefresh {
			screen.Clear()
			screen.MoveTopLeft()
			if w-headerLen > 0 {
				color.White.Println("┏ Goverlay " + strings.Repeat("━", headerLen-13) + "┳" + strings.Repeat("━", w-headerLen))
			} else if w-headerLen == 0 {
				color.White.Println("┏ Goverlay " + strings.Repeat("━", headerLen-13) + "┓")
			} else {
				color.Red.Println("Please increase the width of the terminal.")
				time.Sleep(time.Duration(utils.UpdateSpeed) * time.Millisecond)
				Draw()
			}
			setHeader()
			clines++
			for _, r := range utils.NickedPlayers {
				if clines > ph-3 {
					toomany = true
				} else {
					toomany = false
				}
				if !toomany {
					setNicked(r)
					clines++
				}
			}
			for _, r := range utils.Players {
				if clines > ph-3 {
					toomany = true
				} else {
					toomany = false
				}
				if !toomany {
					setPlayer(r)
					clines++
				}
			}
			if !((h-clines)-4 < 0) {
				color.White.Print(strings.Repeat("┃\n", (h-clines)-5))
			}
			setFooter()
			utils.NeedRefresh = false
		}
		time.Sleep(time.Duration(utils.UpdateSpeed) * time.Millisecond)
	}
}

func getHeaderData() {
	h := "┣ "
	for _, flag := range utils.Layout {
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
	// i honestly hate life after debugging this for an hour
	headerLen = len(strings.Split(h, ""))
	header = h[:len(h)-5] + " ┛"
}

func setHeader() {
	color.White.Println(header)
}

func setFooter() {
	color.Println(color.White.Text("┗━"), color.Gray.Text(Status), strings.Repeat("━", pw-19-len(Status)), color.Gray.Text(color.OpUnderscore.Render("u"))+"pdate", "━", color.Gray.Text(color.OpUnderscore.Render("q"))+"uit")
}

func setPlayer(r utils.Player) {
	color.Printf("┣ ")
	for _, flag := range utils.Layout {
		switch flag {
		case "-lv":
			colors.GetLevelColor(r.Level).Printf("%-8s", strconv.Itoa(r.Level))
		case "-n":
			colors.GetRankColor(r.Rank).Printf("%-18s", r.Name)
		case "-fkdr":
			colors.GetFBColor(r.FKDR).Printf("%-7s", strconv.FormatFloat(r.FKDR, 'f', -1, 64))
		case "-f":
			colors.GetFKColor(r.Finals).Printf("%-9s", strconv.Itoa(r.Finals))
		case "-fd":
			colors.Gray.Printf("%-10s", strconv.Itoa(r.FDeaths))
		case "-wlr":
			colors.GetWLRColor(r.WLR).Printf("%-7s", strconv.FormatFloat(r.WLR, 'f', -1, 64))
		case "-w":
			colors.GetWColor(r.Wins).Printf("%-7s", strconv.Itoa(r.Wins))
		case "-l":
			colors.Gray.Printf("%-9s", strconv.Itoa(r.Losses))
		case "-blr":
			colors.GetWLRColor(r.BBLR).Printf("%-7s", strconv.FormatFloat(r.BBLR, 'f', -1, 64))
		case "-b":
			colors.GetFKColor(r.Beds).Printf("%-7s", strconv.Itoa(r.Beds))
		case "-bl":
			colors.Gray.Printf("%-8s", strconv.Itoa(r.BedsLost))
		case "-i":
			colors.GetIColor(r.Index).Printf("%-8s", strconv.Itoa(r.Index))
		}
	}
	color.Println()
}

func setNicked(r utils.Player) {
	color.Printf("┣ ")
	for _, flag := range utils.Layout {
		switch flag {
		case "-lv":
			colors.Red.Printf("%-8s", "-")
		case "-n":
			colors.Red.Printf("%-18s", r.Name)
		case "-fkdr":
			colors.Red.Printf("%-7s", "-")
		case "-f":
			colors.Red.Printf("%-9s", "-")
		case "-fd":
			colors.Red.Printf("%-10s", "-")
		case "-wlr":
			colors.Red.Printf("%-7s", "-")
		case "-w":
			colors.Red.Printf("%-7s", "-")
		case "-l":
			colors.Red.Printf("%-9s", "-")
		case "-blr":
			colors.Red.Printf("%-7s", "-")
		case "-b":
			colors.Red.Printf("%-7s", "-")
		case "-bl":
			colors.Red.Printf("%-8s", "-")
		case "-i":
			colors.Red.Printf("%-8s", "-")
		}
	}
	color.Println()
}

func SetStatus(str string) {
	Status = str
	utils.NeedRefresh = true
}
