package db

import (
	"testing"
	"time"

	"github.com/GoRustNet/xurl/defs"
	"github.com/GoRustNet/xurl/shorturl"
)

func TestUrlAdd(t *testing.T) {
	testInit(t)
	userID := int64(1)
	targetUrl := "https://gorust.net"
	urlID, err := shorturl.ShortUrlWithSeed(targetUrl, uint32(userID))
	if err != nil {
		t.Fatal(err)
	}
	u := &defs.Url{
		ID:          urlID,
		Url:         targetUrl,
		IsCustomize: false,
		UserID:      userID,
		IsProtected: false,
		Password:    "",
		Dateline:    time.Now(),
		IsDel:       false,
	}

	if err := UrlAdd(u); err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", u)
}
func TestUrlAddCustomize(t *testing.T) {
	testInit(t)
	targetUrl := "https://gorust.net"
	urlID := "gorust"
	u := &defs.Url{
		Url:         targetUrl,
		ID:          urlID,
		IsCustomize: true,
	}

	if err := UrlAdd(u); err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", u)
}
