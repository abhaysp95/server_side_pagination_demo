package main

type Car struct {
	Make string `json:"make"`
	Model string `json:"model"`
	Vin string `json:"vin"`
}

type UserPurchase struct {
	Id int `json:"id"`
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Gender string `json:"gender"`
	Car *Car `json:"car"`
}
