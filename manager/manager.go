package manager

type Filter struct {
	Skip     int
	Limit    int
	Query    interface{}
	SortCols []string
}
