package service

import (
	"testing"
	"time"

	"github.com/GoRustNet/xurl/defs"
)

func TestUrlAddCustomizeExists(t *testing.T) {
	testInit(t)
	userID := int64(2)
	targetUrl := "https://gorust.net"
	urlID := "gorust"
	u := &defs.Url{
		ID:          urlID,
		Url:         targetUrl,
		IsCustomize: true,
		IsDel:       false,
	}
	uu := &defs.UserUrl{
		UserID:      userID,
		UrlID:       urlID,
		IsProtected: false,
		Password:    "",
		IsDel:       false,
		Dateline:    time.Now(),
	}
	uuv, err := UrlAdd(u, uu)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(uuv)
}
func TestUrlAddCustomizeNotExists(t *testing.T) {
	testInit(t)
	userID := int64(2)
	targetUrl := "https://github.com"
	urlID := "github"
	u := &defs.Url{
		ID:          urlID,
		Url:         targetUrl,
		IsCustomize: true,
		IsDel:       false,
	}
	uu := &defs.UserUrl{
		UserID:      userID,
		UrlID:       urlID,
		IsProtected: false,
		Password:    "",
		IsDel:       false,
		Dateline:    time.Now(),
	}
	uuv, err := UrlAdd(u, uu)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", uuv)
}
func TestUrlAddCustomizeNotExistsProtected(t *testing.T) {
	testInit(t)
	userID := int64(2)
	targetUrl := "https://gitlab.com"
	urlID := "gitlab"
	u := &defs.Url{
		ID:          urlID,
		Url:         targetUrl,
		IsCustomize: true,
		IsDel:       false,
	}
	uu := &defs.UserUrl{
		UserID:      userID,
		UrlID:       urlID,
		IsProtected: true,
		Password:    "foobar",
		IsDel:       false,
		Dateline:    time.Now(),
	}
	uuv, err := UrlAdd(u, uu)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", uuv)
}
func TestUrlAddNotExists(t *testing.T) {
	testInit(t)
	userID := int64(2)
	targetUrl := "https://example.com"
	u := &defs.Url{
		Url:         targetUrl,
		IsCustomize: false,
		IsDel:       false,
	}
	uu := &defs.UserUrl{
		UserID:      userID,
		IsProtected: true,
		Password:    "foo.bar",
		IsDel:       false,
		Dateline:    time.Now(),
	}
	uuv, err := UrlAdd(u, uu)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(uuv)
}
