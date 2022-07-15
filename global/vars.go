package global

var (
	RefreshTime int
	NeedRefresh bool
)

func Refresh() {
	NeedRefresh = true
}
