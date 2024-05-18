package auth

import (
	"errors"
	"github.com/aminzdev/db"
	"testing"
	"time"
)

func TestAuth_SignOut(t *testing.T) {
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

	auth.db.Unscoped().Where("name = ?", "user").Delete(&User{})
	auth.db.Unscoped().Where("name = ?", "user").Delete(&Session{})

	if err = auth.SignOut("user", "pass"); !errors.Is(err, ErrWrongCredentials) {
		t.Fatalf("expected error %s but got %v", ErrWrongCredentials, err)
	}

	if err = auth.SignUp("user", "pass"); err != nil {
		t.Fatalf("expected no error but got: %s", err)
	}

	if err = auth.SignOut("not_user", "pass"); !errors.Is(err, ErrWrongCredentials) {
		t.Fatalf("expected error %s but got %v", ErrWrongCredentials, err)
	}

	if err = auth.SignOut("user", "wrong pass"); !errors.Is(err, ErrWrongCredentials) {
		t.Fatalf("expected error %s but got %v", ErrWrongCredentials, err)
	}

	if err = auth.SignOut("user", "pass"); err != nil {
		t.Fatalf("expected no error but got %v", err)
	}

	if err = auth.ValidateSession("user", "pass"); !errors.Is(err, ErrSessionNotFound) {
		t.Fatalf("expected error %s but got %v", ErrSessionNotFound, err)
	}

	if err = auth.SignIn("user", "pass", time.Minute); err != nil {
		t.Fatalf("expected no error but got: %s", err)
	}

	if err = auth.SignOut("user", "pass"); err != nil {
		t.Fatalf("expected no error but got: %s", err)
	}

	if err = auth.ValidateSession("user", "pass"); !errors.Is(err, ErrExpiredSession) {
		t.Fatalf("expected error %s but got %v", ErrExpiredSession, err)
	}

	auth.db.Unscoped().Where("name = ?", "user").Delete(&User{})
	auth.db.Unscoped().Where("name = ?", "user").Delete(&Session{})
}
