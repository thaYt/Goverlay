package api

import (
	"Goverlay/global"
	"time"
)

var (
	LoadingList []Player
	PlayerList  []Player
	NickedList  []Player
)

func addLoading(player Player) {
	if inLoading(player) {
		return
	}
	LoadingList = append(LoadingList, player)
}

func removeLoading(player Player) {
	var list []Player
	for _, p := range LoadingList {
		if p.Name != player.Name {
			list = append(list, p)
		}
	}
	LoadingList = list
}

func inLoading(player Player) bool {
	for _, i := range LoadingList {
		if i.Name == player.Name {
			return true
		}
	}
	return false
}

func addPlayer(player Player) {
	if inList(player) {
		return
	}
	PlayerList = append(PlayerList, player)
}

func RemovePlayer(player string) {
	var list []Player
	for _, p := range PlayerList {
		if p.Name != player {
			list = append(list, p)
		}
	}
	PlayerList = list
	global.Refresh()
}

func inList(player Player) bool {
	for _, i := range PlayerList {
		if i.Name == player.Name {
			return true
		}
	}
	return false
}

func addNicked(player Player) {
	if inNicked(player) {
		return
	}
	NickedList = append(NickedList, player)
}

func RemoveNicked(player string) {
	var list []Player
	for _, p := range NickedList {
		if p.Name != player {
			list = append(list, p)
		}
	}
	NickedList = list
	global.Refresh()
}

func inNicked(player Player) bool {
	for _, i := range NickedList {
		if i.Name == player.Name {
			return true
		}
	}
	return false
}

func Nuke() {
	NickedList, PlayerList, LoadingList = []Player{}, []Player{}, []Player{}
	global.Refresh()
}

func Moderate() {
	for {
		for _, player := range LoadingList {
			for _, p := range PlayerList {
				if p.Name == player.Name {
					GetStats(p.Name)
				}
			}
		}
		time.Sleep(10 * time.Second)
	}
}
