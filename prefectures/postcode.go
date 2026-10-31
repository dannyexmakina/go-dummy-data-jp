package prefectures

type Prefecture struct {
	Name  string
	Rate  float64
	Areas []Area
}
type Area struct {
	Code    string
	Address string
	Tel     string // 市外局番
}
