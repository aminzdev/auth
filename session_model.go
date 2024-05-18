package auth

import (
	"github.com/aminzdev/db"
	"time"
)

type Session struct {
	db.Model
	Name      string `gorm:"unique"`
	ExpiredAt time.Time
}
