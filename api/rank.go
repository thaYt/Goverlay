package api

import json "github.com/buger/jsonparser"

func getRank(data []byte) string {
	_, _, _, err := json.Get(data, "player", "rank")
	if err == nil {
		e, _ := json.GetString(data, "player", "rank")
		if e == "ADMIN" {
			return "ADMIN"
		} else if e == "GAME_MASTER" {
			return "GM"
		} else if e == "YOUTUBER" {
			return "YOUTUBE"
		}
	}
	_, _, _, err = json.Get(data, "player", "monthlyPackageRank")
	if err == nil {
		e, _ := json.GetString(data, "player", "monthlyPackageRank")
		if e == "NONE" {
			return "MVP+"
		} else if e == "SUPERSTAR" {
			return "MVP++"
		}
	}
	_, _, _, err = json.Get(data, "player", "newPackageRank")
	if err == nil {
		e, _ := json.GetString(data, "player", "newPackageRank")
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
	}
	return "NON"
}
