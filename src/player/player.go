package player

import (
	"sort"
	"thaYt/Goverlay/src/printer"
	"thaYt/Goverlay/src/utils"
)

var cleared bool

func AddPlayer(player utils.Player) {
	if !inSlice(player) {
		if player.Nicked {
			utils.NickedPlayers = append(utils.NickedPlayers, player)
		} else if !player.Nicked {
			utils.Players = append(utils.Players, player)
		}
	}
	sortNicked()
	sortPlayers()
	cleared, utils.NeedRefresh = false, true
}

func inSlice(player utils.Player) bool {
	for _, e := range utils.Players {
		if player.Name == e.Name {
			return true
		}
	}

	return false
}

func Nuke() {
	if !cleared {
		utils.Players, utils.NickedPlayers, utils.NeedRefresh, cleared = []utils.Player{}, []utils.Player{}, true, true
	}
}

func RemovePlayer(player utils.Player) {
	if !inSlice(player) {
		return
	}
	var r []utils.Player
	for _, e := range utils.Players {
		if e.Name != player.Name {
			r = append(r, e)
		}
	}
	var v []utils.Player
	for _, e := range utils.NickedPlayers {
		if e.Name != player.Name {
			printer.SetStatus(printer.Status + " status, " + player.Name + "removed")
			v = append(v, e)
		}
	}
	utils.Players, utils.NickedPlayers = r, v
	sortPlayers()
	sortNicked()
	utils.NeedRefresh = true
}

func sortPlayers() {
	sort.Slice(utils.Players, func(i, j int) bool {
		return utils.Players[i].Level > utils.Players[j].Level
	})
}

func sortNicked() {
	sort.Slice(utils.NickedPlayers, func(i, j int) bool {
		return utils.NickedPlayers[i].Name > utils.NickedPlayers[j].Name
	})
}
