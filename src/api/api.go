package api

import (
	"io"
	"math"
	"net/http"
	"thaYt/Goverlay/src/player"
	"thaYt/Goverlay/src/utils"

	"github.com/Jeffail/gabs/v2"
)

func GetHypixelData(name string, y int) {
	if y == 0 {
		w, err := http.Get("https://api.mojang.com/users/profiles/minecraft/" + name)
		if err != nil {
			e := &utils.Player{
				Nicked: true,
				Name:   name,
			}
			player.AddPlayer(*e)
			return
		}
		defer w.Body.Close()

		rr, err := io.ReadAll(w.Body)
		if err != nil {
			e := &utils.Player{
				Nicked: true,
				Name:   name,
			}
			player.AddPlayer(*e)
			return
		}

		v, err := gabs.ParseJSON(rr)
		if err != nil {
			e := &utils.Player{
				Nicked: true,
				Name:   name,
			}
			player.AddPlayer(*e)
			return
		}
		p := v.Search("id").String()
		uuid := p[1 : len(p)-1]

		resp, err := http.Get("https://api.hypixel.net/player?uuid=" + uuid + "&key=" + utils.Key)
		if err != nil {
			e := &utils.Player{
				Nicked: true,
				Name:   name,
			}
			player.AddPlayer(*e)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			e := &utils.Player{
				Nicked: true,
				Name:   name,
			}
			player.AddPlayer(*e)
			return
		}

		jsonParsed, err := gabs.ParseJSON(body)
		if err != nil {
			e := &utils.Player{
				Nicked: true,
				Name:   name,
			}
			player.AddPlayer(*e)
			return
		}
		if jsonParsed.Search("player").String() == "null" {
			e := &utils.Player{
				Nicked: true,
				Name:   name,
			}
			player.AddPlayer(*e)
			return
		}
		var (
			rank     string
			star     float64
			finals   float64
			fdeaths  float64
			wins     float64
			losses   float64
			beds     float64
			bedsLost float64
			index    float64
			fkdr     float64
			bblr     float64
			wlr      float64
		)
		if jsonParsed.Search("player", "stats", "Bedwars").String() == "null" || jsonParsed.Search("player", "achievements").String() == "null" {
			e := &utils.Player{
				Nicked: true,
				Name:   name,
			}
			player.AddPlayer(*e)
			return
		}

		// this fuckery is due to it dying if an unhandled null exception occuring
		// if *any one* of these don't exist, the entire program breaks
		// if anyone has a better way to do this, please tell me because i hate this
		if jsonParsed.Search("player", "achievements", "bedwars_level").String() == "null" {
			star = 0
		} else {
			star = jsonParsed.Search("player", "achievements", "bedwars_level").Data().(float64)
		}

		if jsonParsed.Search("player", "stats", "Bedwars", "final_kills_bedwars").String() == "null" {
			finals = 0
		} else {
			finals = jsonParsed.Search("player", "stats", "Bedwars", "final_kills_bedwars").Data().(float64)
		}
		if jsonParsed.Search("player", "stats", "Bedwars", "final_deaths_bedwars").String() == "null" {
			fdeaths = 0
		} else {
			fdeaths = jsonParsed.Search("player", "stats", "Bedwars", "final_deaths_bedwars").Data().(float64)
		}

		if jsonParsed.Search("player", "stats", "Bedwars", "wins_bedwars").String() == "null" {
			wins = 0
		} else {
			wins = jsonParsed.Search("player", "stats", "Bedwars", "wins_bedwars").Data().(float64)
		}
		if jsonParsed.Search("player", "stats", "Bedwars", "losses_bedwars").String() == "null" {
			losses = 0
		} else {
			losses = jsonParsed.Search("player", "stats", "Bedwars", "losses_bedwars").Data().(float64)
		}

		if jsonParsed.Search("player", "stats", "Bedwars", "beds_broken_bedwars").String() == "null" {
			beds = 0
		} else {
			beds = jsonParsed.Search("player", "stats", "Bedwars", "beds_broken_bedwars").Data().(float64)
		}
		if jsonParsed.Search("player", "stats", "Bedwars", "beds_lost_bedwars").String() == "null" {
			bedsLost = 0
		} else {
			bedsLost = jsonParsed.Search("player", "stats", "Bedwars", "beds_lost_bedwars").Data().(float64)
		}
		fkdr = math.Round((finals/fdeaths)*100) / 100
		wlr = math.Round((wins/losses)*100) / 100
		bblr = math.Round((beds/bedsLost)*100) / 100
		index = math.Round(wlr*50 + fkdr*25 + star*10)
		rank = utils.GetRank(jsonParsed)
		e := &utils.Player{
			Nicked:   false,
			Rank:     rank,
			Level:    int(star),
			Name:     name,
			FKDR:     fkdr,
			WLR:      wlr,
			BBLR:     bblr,
			Finals:   int(finals),
			FDeaths:  int(fdeaths),
			Wins:     int(wins),
			Losses:   int(losses),
			Beds:     int(beds),
			BedsLost: int(bedsLost),
			Index:    int(index),
		}
		player.AddPlayer(*e)
	} else {
		e := &utils.Player{
			Name: name,
		}
		player.RemovePlayer(*e)
	}
}
