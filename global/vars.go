package global

var (
	RefreshTime int
	NeedRefresh bool
	Debug       bool
)

const (
	Version = 2.0
)

func Refresh() {
	NeedRefresh = true
}
