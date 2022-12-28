package entity

type User struct {
	ID       int32  `json:"id" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username"`
}
