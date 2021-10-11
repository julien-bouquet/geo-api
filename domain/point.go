package domain

type Points []Point

type Point struct {
	Name string `json:"name"`
	Lat  string `json:"lattitude"`
	Long string `json:"longitude"`
}
