package auth

import (
	"fmt"
	"os"
	"time"
)

func (a *Auth) SignOut(name string) error {
	res := a.db.Model(&Session{}).Where("name = ?", name).Update("expired_at", time.Now())
	if res.Error != nil {
		fmt.Fprintf(os.Stderr, "could not find session: %s", res.Error)
		return ErrInternal
	}
	return nil
}
