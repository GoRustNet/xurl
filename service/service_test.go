package service

import (
	"testing"

	"github.com/GoRustNet/xurl/conf"
	"github.com/GoRustNet/xurl/db"
)

func testInit(t *testing.T) {

	if err := conf.InitFrom("../config.json"); err != nil {
		t.Fatal(err)
	}
	if err := db.Init(conf.Cfg.Pg); err != nil {
		t.Fatal(err)
	}
}
