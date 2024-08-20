package models

import "time"

type UsersModels struct {
	Id         int       `json:"id"`
	Nama       string    `json:"nama"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"createdAt"`
	Updated_at time.Time `json:"updatedAt"`
}
