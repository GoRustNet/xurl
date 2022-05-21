package shorturl

import "testing"

func TestShortUrl(t *testing.T) {
	const targetUrl = "https://github.com"
	const shortedUrl = "Sf5oy"
	surl, err := ShortUrl(targetUrl)
	if err != nil {
		t.Fatal(err)
	}
	if surl != shortedUrl {
		t.Fatalf("Want %q but got %q", shortedUrl, surl)
	}
}
func TestShortUrlWithSeed(t *testing.T) {
	const targetUrl = "https://github.com"
	const shortedUrl = "1YW552"
	surl, err := ShortUrlWithSeed(targetUrl, 100)
	if err != nil {
		t.Fatal(err)
	}
	if surl != shortedUrl {
		t.Fatalf("Want %q but got %q", shortedUrl, surl)
	}
}
