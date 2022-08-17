package api

import (
	"Goverlay/global"
	json "github.com/buger/jsonparser"
	"github.com/gookit/color"
	"io"
	"math"
	"net/http"
	"strings"
)

type Player struct {
	Name    string
	Rank    string
	Level   int
	Finals  int
	FDeaths int
	Wins    int
	Losses  int
	Beds    int
	Bl      int
	FKDR    float64
	BBLR    float64
	WLR     float64
}

var (
	Key      string
	ValidKey bool
	KDR      string
)

func GetStats(name string) {
	player := Player{
		Name: name,
	}
	addLoading(player)
	println(player.Name)

	cached, player := inCache(player)
	global.Refresh()

	if cached {
		addPlayer(player)
		removeLoading(player)
		global.Refresh()
		return
	}

	mojangHTTP, err := http.Get("https://api.mojang.com/users/profiles/minecraft/" + name)
	if err != nil {
		handle(name, err, player)
		return
	}
	mojangJSON, err := io.ReadAll(mojangHTTP.Body)
	if err != nil {
		handle(name, err, player)
		return
	}
	id, err := json.GetString(mojangJSON, "id")
	if err != nil {
		handle(name, err, player)
		return
	}

	hypixelHTTP, err := http.Get("https://api.hypixel.net/player?uuid=" + id + "&key=" + Key)
	if err != nil {
		handle(name, err, player)
		return
	}
	hypixelJSON, err := io.ReadAll(hypixelHTTP.Body)
	if err != nil {
		handle(name, err, player)
		return
	}
	if strings.Contains(string(hypixelJSON), `"player":null}`) {
		handle(name, nil, player)
	}
	success, _ := json.GetBoolean(hypixelJSON)
	if !success {
		cause, _ := json.GetString(hypixelJSON, "cause")
		if cause == "Key throttle" {
			KDR = "THROTTLE"
			ValidKey = false
		}
	}

	// if any of these are null, then the player doesn't exist, is nicked, doesn't have achievements (impossible), or hasn't played bedwars before
	_, _, _, err = json.Get(hypixelJSON, "player")
	if err != nil && err.Error() != "Key path not found" {
		handle(name, err, player)
		return
	}
	_, _, _, err = json.Get(hypixelJSON, "player", "stats", "bedwars")
	if err != nil && err.Error() != "Key path not found" {
		handle(name, err, player)
		return
	}
	_, _, _, err = json.Get(hypixelJSON, "player", "achievements")
	if err != nil && err.Error() != "Key path not found" {
		handle(name, err, player)
		return
	}

	// stat time
	level, err := json.GetInt(hypixelJSON, "player", "achievements", "bedwars_level") // level or star
	if err != nil {
		level = 0
	}
	finals, err := json.GetInt(hypixelJSON, "player", "stats", "Bedwars", "final_kills_bedwars") // finals
	if err != nil {
		finals = 0
	}
	fdeaths, err := json.GetInt(hypixelJSON, "player", "stats", "Bedwars", "final_deaths_bedwars") // final deaths
	if err != nil {
		fdeaths = 0
	}
	wins, err := json.GetInt(hypixelJSON, "player", "stats", "Bedwars", "wins_bedwars") // wins
	if err != nil {
		wins = 0
	}
	losses, err := json.GetInt(hypixelJSON, "player", "stats", "Bedwars", "losses_bedwars") // losses
	if err != nil {
		losses = 0
	}
	beds, err := json.GetInt(hypixelJSON, "player", "stats", "Bedwars", "beds_broken_bedwars") // beds
	if err != nil {
		beds = 0
	}
	blost, err := json.GetInt(hypixelJSON, "player", "stats", "Bedwars", "beds_lost_bedwars") // beds lost
	if err != nil {
		blost = 0
	}

	fkdr := math.Round(float64(finals)/float64(fdeaths)*100) / 100
	wlr := math.Round(float64(wins)/float64(losses)*100) / 100
	bblr := math.Round(float64(beds)/float64(blost)*100) / 100

	// todo check if infinite - or + and use 0 or numerator respectively

	player.Level = int(level)
	player.Rank = getRank(hypixelJSON)
	player.Finals = int(finals)
	player.FDeaths = int(fdeaths)
	player.Wins = int(wins)
	player.Losses = int(losses)
	player.Beds = int(beds)
	player.Bl = int(blost)
	player.FKDR = fkdr
	player.WLR = wlr
	player.BBLR = bblr

	removeLoading(player)
	addPlayer(player)
	addCache(player)
	global.Refresh()
}

func handle(name string, err error, player Player) {
	if global.Debug {
		color.Println("Failed " + name)
		color.Red.Println(err.Error())
	}
	addNicked(player)
	removeLoading(player)
}

func CheckKey() bool {
	KeyHTML, _ := http.Get("https://api.hypixel.net/key?key=" + Key)
	KeyJSON, _ := io.ReadAll(KeyHTML.Body)
	success, _ := json.GetBoolean(KeyJSON, "success")
	if !success {
		reason, _ := json.GetString(KeyJSON, "cause")
		if reason == "Invalid API key" {
			KDR = "INVALID"
		} else if reason == "Key throttle" {
			KDR = "THROTTLE"
		} else {
			KDR = "UNKNOWN REASON"
		}
		return false
	}
	return true
}
