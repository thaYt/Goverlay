package api

import (
	"sort"
)

func SortPlayers() {
	sort.Slice(PlayerList, func(i, j int) bool {
		return PlayerList[i].Level > PlayerList[j].Level
	})
}

func SortNicked() {
	sort.Slice(NickedList, func(i, j int) bool {
		return NickedList[i].Name > NickedList[j].Name
	})
}
