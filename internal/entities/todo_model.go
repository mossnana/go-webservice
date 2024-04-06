package entities

import "time"

type Todo struct {
	ID          string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Title       string
	Description string
	CreateDate  time.Time
	UpdateDate  time.Time
	DeleteDate  time.Time
}

type DraftTodo struct {
	Title       string
	Description string
}
