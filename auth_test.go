package auth

import (
	"github.com/aminzdev/db"
	"testing"
	"time"
)

func TestAuth(t *testing.T) {
	_db, err := db.NewDB(
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
		t.Fatal(err)
	}

	auth := NewAuth(_db)
	t.Log("signing up ...")
	if err = auth.SignUp("user", "pass"); err != nil {
		t.Fatal(err)
	}
	t.Log("signing in ...")
	if err = auth.SignIn("user", "pass", time.Hour*5); err != nil {
		t.Fatal(err)
	}
	t.Log("signing out ...")
	time.Sleep(time.Second)
	if err = auth.SignOut("user", "pass"); err != nil {
		t.Fatal(err)
	}
	t.Log("signing in ...")
	time.Sleep(time.Second)
	if err = auth.SignIn("user", "pass", time.Hour*5); err != nil {
		t.Fatal(err)
	}
	t.Log("signing out ...")
	time.Sleep(time.Second)
	if err = auth.SignOut("user", "pass"); err != nil {
		t.Fatal(err)
	}
}
