package models

import "time"

type History struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	Amount    int       `json:"amount"`
	ShopsId   int       `json:"shops_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Items struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	ShopsId   int       `json:"shops_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Shop struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `jsono:"updated_at"`
}

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
