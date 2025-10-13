package models

import "time"

type UserModel struct {
	ID        int    `gorm:"primary_key"`
	Name      string `gorm:"size:16;not null;unique"`
	Age       int    `gorm:"default:18;check:age > 0"`
	CreatedAt time.Time
}
