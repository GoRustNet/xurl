package db

import (
	"testing"

	"github.com/GoRustNet/xurl/conf"
)

func testInit(t *testing.T) {
	if err := conf.InitFrom("../config.json"); err != nil {
		t.Fatal(err)
	}
	if err := Init(conf.Cfg.Pg); err != nil {
		t.Fatal(err)
	}
}
