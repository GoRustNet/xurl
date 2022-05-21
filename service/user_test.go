package service

import (
	"testing"
	"time"

	"github.com/GoRustNet/xurl/defs"
)

func TestUserRegister(t *testing.T) {
	testInit(t)

	u := &defs.User{
		Email:      "foo@bar.com",
		Password:   "net.gorust",
		Permission: defs.UserPermissionSysGenerateAndCustomizeUrl,
		Status:     defs.UserStatusNormal,
		Dateline:   time.Now(),
	}
	if err := UserRegister(u); err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestGetUserByEmail(t *testing.T) {
	testInit(t)
	u, err := GetUserByEmail("team@gorust.net")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestGetUserById(t *testing.T) {
	testInit(t)
	u, err := GetUserById(2)
	if err != nil {
		t.Fatal(err.Debug())
	}
	t.Log(u)
}
