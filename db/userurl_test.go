package db

import (
	"testing"

	"github.com/GoRustNet/xurl/defs"
)

func TestGetUserUrlLiteViewByUrl(t *testing.T) {
	testInit(t)
	u := &defs.Url{
		ID:     "2YgfLM",
		UserID: 1,
	}
	uuv, err := GetUserUrlLiteViewByUrl(u)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(uuv)
}
