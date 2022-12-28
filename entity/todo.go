package entity

import "time"

type Todo struct {
	ID        int32     `json:"id" gorm:"primaryKey;autoIncrement"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
