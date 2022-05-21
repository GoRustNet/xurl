package db

import (
	"fmt"
	"testing"
	"time"

	"github.com/GoRustNet/xurl/defs"
	"github.com/GoRustNet/xurl/pwd"
)

func TestUserList(t *testing.T) {
	testInit(t)
	p, err := UserList(0)
	if err != nil {
		t.Fatal(err)
	}
	for _, u := range p.Data {
		t.Logf("%#v", u)
	}
}

func TestUserAdd(t *testing.T) {
	testInit(t)
	p, err := pwd.Hash("net.gorust")
	if err != nil {
		t.Fatal(err)
	}
	//rand.Seed(int64(time.Now().Nanosecond()))
	u := &defs.User{
		Email:      fmt.Sprintf("team@gorust.net"),
		Password:   p,
		Permission: defs.UserPermissionSysGenerateAndCustomizeUrl,
		Status:     defs.UserStatusNormal,
		Dateline:   time.Now(),
	}
	id, err := UserAdd(u)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}
