package entity

type Header struct {
	Id     int    `json:"-"`
	Name   string `json:"name"`
	Value  string `json:"value"`
	MockId int    `json:"mock_id"`
}
