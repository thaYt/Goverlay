package global

var (
	RefreshTime int
	NeedRefresh bool
	Debug       bool
)

const (
	Version = 1.0
)

func Refresh() {
	NeedRefresh = true
}
