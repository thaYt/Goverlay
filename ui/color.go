package ui

import "github.com/gookit/color"

var ( // colors
	DarkAqua  = color.HEXStyle("249d9f")
	DarkGreen = color.HEXStyle("0a0")
	DarkRed   = color.HEXStyle("a00")
	DarkPink  = color.HEXStyle("8b008b")
	Aqua      = color.HEXStyle("3ff")
	Blue      = color.HEXStyle("043cc8")
	Green     = color.HEXStyle("5f5")
	Yellow    = color.HEXStyle("ff5")
	Pink      = color.HEXStyle("f699cd")
	Gold      = color.HEXStyle("fa0")
	Red       = color.HEXStyle("f55")
	White     = color.HEXStyle("fff")
	Gray      = color.HEXStyle("7d7d7d")
	Vblack    = color.HEXStyle("333")
)

func GetLevelColor(level int) *color.RGBStyle {
	if level >= 1000 {
		return Vblack
	} else if level >= 900 {
		return DarkPink
	} else if level >= 800 {
		return Blue
	} else if level >= 700 {
		return Pink
	} else if level >= 600 {
		return DarkRed
	} else if level >= 500 {
		return DarkAqua
	} else if level >= 400 {
		return DarkGreen
	} else if level >= 300 {
		return Aqua
	} else if level >= 200 {
		return Gold
	} else if level >= 100 {
		return White
	} else {
		return Gray
	}
}

func GetRankColor(rank string) *color.RGBStyle {
	switch rank {
	case "ADMIN":
		return Red
	case "YOUTUBE":
		return Red
	case "GM":
		return DarkGreen
	case "MVP++":
		return Gold
	case "MVP+":
		return Aqua
	case "MVP":
		return Aqua
	case "VIP+":
		return Green
	case "VIP":
		return Green
	default:
		return Gray
	}
}

func GetFBColor(fkdr float64) *color.RGBStyle {
	if fkdr >= 10 {
		return DarkRed
	} else if fkdr >= 5 {
		return Red
	} else if fkdr >= 3 {
		return Gold
	} else if fkdr >= 2 {
		return Yellow
	} else if fkdr >= 1 {
		return Green
	} else {
		return DarkGreen
	}
}

func GetWLRColor(wlr float64) *color.RGBStyle {
	if wlr >= 5 {
		return DarkRed
	} else if wlr >= 2.5 {
		return Red
	} else if wlr >= 1.5 {
		return Gold
	} else if wlr >= 1 {
		return Yellow
	} else if wlr >= .5 {
		return Green
	} else {
		return DarkGreen
	}
}

func GetFKColor(fk int) color.Color {
	if fk >= 1000 {
		return color.White
	} else {
		return color.Gray
	}
}

func GetWColor(w int) color.Color {
	if w >= 250 {
		return color.White
	} else {
		return color.Gray
	}
}
