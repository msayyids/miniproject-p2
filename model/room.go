package model

type Rooms struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Price        int    `json:"price"`
	Availibility bool   `json:"availibility"`
}
