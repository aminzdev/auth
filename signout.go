package auth

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"os"
	"time"
)

func (a *Auth) SignOut(name, code string) error {
	var user *User
	if res := a.db.First(&user, "name = ? and code = ?", name, code); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrWrongCredentials
		}
		fmt.Fprintf(os.Stderr, "could not find user: %s", res.Error)
		return ErrInternal
	}

	res := a.db.Model(&Session{}).Where("name = ?", name).Update("expired_at", time.Now())
	if res.Error != nil {
		fmt.Fprintf(os.Stderr, "could not find session: %s", res.Error)
		return ErrInternal
	}
	return nil
}
