package auth

import (
	"github.com/aminzdev/db"
	"testing"
)

func TestAuth_SignUp(t *testing.T) {
	database, err := db.NewDB(
		&db.Postgres{
			Host:     "localhost:5432",
			User:     "user",
			Pass:     "pass",
			DBName:   "db",
			SSL:      "disable",
			TimeZone: "Asia/Tehran",
		},
		&User{},
		&Session{},
	)
	if err != nil {
		t.Fatalf("expected no error but got: %s", err)
	}

	auth := NewAuth(database)

	validUserNames := []string{"Username", "username", "user_name", "user1", "1user"}
	invalidUserNames := []string{"space in username", "user!", "*user*"}

	for _, name := range validUserNames {
		auth.db.Unscoped().Where("name = ?", name).Delete(&User{})
	}
	for _, name := range invalidUserNames {
		auth.db.Unscoped().Where("name = ?", name).Delete(&User{})
	}

	for _, name := range validUserNames {
		if err = auth.SignUp(name, "pass"); err != nil {
			t.Fatalf("expected no error but got: %s", err)
		}
	}
	for _, name := range invalidUserNames {
		if err = auth.SignUp(name, "pass"); err == nil {
			t.Fatalf("expected error %s but got: %s", ErrInvalidUserName, err)
		}
	}

	for _, name := range validUserNames {
		auth.db.Unscoped().Where("name = ?", name).Delete(&User{})
	}
	for _, name := range invalidUserNames {
		auth.db.Unscoped().Where("name = ?", name).Delete(&User{})
	}
}
