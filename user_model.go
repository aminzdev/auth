package auth

import "github.com/aminzdev/db"

type User struct {
	db.Model
	Name string
	Code string
}
