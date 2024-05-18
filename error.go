package auth

import "errors"

var (
	ErrInternal          = errors.New("internal error")
	ErrWrongCredentials  = errors.New("wrong credentials")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidUserName   = errors.New("invalid username")
	ErrSessionNotFound   = errors.New("not signed in")
	ErrExpiredSession    = errors.New("session is expired")
)
