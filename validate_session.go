package auth

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"os"
	"time"
)

func (a *Auth) ValidateSession(name, code string) error {
	var user *User
	if res := a.db.First(&user, "name = ? and code = ?", name, code); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrWrongCredentials
		}
		fmt.Fprintf(os.Stderr, "could not find user: %s", res.Error)
		return ErrInternal
	}

	var session *Session
	if res := a.db.First(&session, "name = ?", name); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrSessionNotFound
		}
		fmt.Fprintf(os.Stderr, "could not find user: %s", res.Error)
		return ErrInternal
	}

	if time.Now().After(session.ExpiredAt) {
		return ErrExpiredSession
	}

	return nil
}
