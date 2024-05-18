package auth

import "github.com/aminzdev/db"

type Auth struct {
	db *db.DB
}

func NewAuth(db *db.DB) *Auth {
	return &Auth{db: db}
}
