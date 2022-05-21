package db

import (
	"testing"

	"github.com/GoRustNet/xurl/defs"
	"github.com/GoRustNet/xurl/shorturl"
)

func TestUrlAdd(t *testing.T) {
	testInit(t)
	targetUrl := "https://gorust.net"
	urlID, err := shorturl.ShortUrl(targetUrl)
	if err != nil {
		t.Fatal(err)
	}
	u := &defs.Url{
		Url:         targetUrl,
		ID:          urlID,
		IsCustomize: false,
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
