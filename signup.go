package auth

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"os"
	"regexp"
)

func (a *Auth) SignUp(name, code string) error {
	if !regexp.MustCompile(`^\w+$`).MatchString(name) {
		return ErrInvalidUserName
	}

	res := a.db.First(&User{}, "name = ?", name)
	if res.Error == nil {
		return ErrUserAlreadyExists
	}

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		if res = a.db.Create(&User{Name: name, Code: code}); res.Error != nil {
			fmt.Fprintf(os.Stderr, "could not create user: %s", res.Error)
			return ErrInternal
		}
	} else {
		fmt.Fprintf(os.Stderr, "could not verify user existance: %s", res.Error)
		return ErrInternal
	}

	return nil
}
