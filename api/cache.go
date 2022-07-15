package api

var CacheList []Player

func inCache(player Player) (bool, Player) {
	if len(CacheList) == 0 {
		return false, player
	}

	for _, p := range CacheList {
		if p.Name == player.Name {
			return true, p
		}
	}
	return false, player
}

func ClearCache() {
	CacheList = []Player{}
}

func addCache(player Player) {
	contains, _ := inCache(player)
	if !contains {
		CacheList = append(CacheList, player)
	}
}
