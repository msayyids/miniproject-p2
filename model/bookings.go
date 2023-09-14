package model

type Bookings struct {
	Id          int    `json:"id"`
	User_id     int    `json:"user_id"`
	Room_id     int    `json:"room_id"`
	Total_day   int    `json:"total_day"`
	Total_Price int    `json:"total_amount"`
	Status      string `json:"status"`
}
