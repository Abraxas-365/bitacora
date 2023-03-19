package models

import (
	"time"

	"github.com/google/uuid"
)

type Report struct {
	Id           uuid.UUID `bson:"_id" json:"id"`
	Title        string    `bson:"title" json:"title"`
	Tags         []string  `bson:"tags" json:"tags"`
	Images       []string  `bson:"images" json:"images"`
	Data         string    `bson:"description" json:"description"`
	Error        string    `bson:"error" json:"error"`
	Solution     string    `bson:"solution" json:"solution"`
	UserNickname string    `bson:"nickname" json:"nickname"`
	Status       bool      `bson:"status" json:"status"`
	Delete       bool      `bson:"delete" json:"delete"`
	CreationDate time.Time `bson:"date_time" json:"date_time"`
}
type Reports []Report

func (r *Report) Constructor(nickname string) Report {
	r.Id = uuid.New()
	r.CreationDate = time.Now()
	r.Delete = false
	r.UserNickname = nickname
	return *r
}
