package model

type Users struct {
	Id             int    `json:"id"`
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Password       string `json:"password" binding:"required,min=6"`
	PhoneNumber    string `json:"phone_number" binding:"required"`
	Deposit_amount int    `json:"deposit_amount"`
}
