package auth

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"os"
	"time"
)

func (a *Auth) SignIn(name, code string, duration time.Duration) error {
	var user *User
	if res := a.db.First(&user, "name = ? and code = ?", name, code); res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return ErrWrongCredentials
		}
		fmt.Fprintf(os.Stderr, "could not find user: %s", res.Error)
		return ErrInternal
	}

	res := a.db.
		Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "name"}}, UpdateAll: true}).
		Create(&Session{Name: name, ExpiredAt: time.Now().Add(duration)})
	if res.Error != nil {
		fmt.Fprintf(os.Stderr, "could not create session: %s", res.Error)
		return ErrInternal
	}

	return nil
}
