package user

import "time"

type User struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (u *User) TableName() string {
	return "user"
}
