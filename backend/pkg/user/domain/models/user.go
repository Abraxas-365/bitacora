package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id           interface{} `bson:"_id" json:"id"`
	Email        string      `bson:"email" json:"email"`
	Nickname     string      `bson:"nickname" json:"nickname"`
	Password     string      `bson:"password" json:"password"`
	CreationDate time.Time   `bson:"creation_date" json:"creation_date"`
}
type Users []*User

func (u *User) Constructor() User {
	u.Id = uuid.New()
	u.CreationDate = time.Now()
	return *u
}
